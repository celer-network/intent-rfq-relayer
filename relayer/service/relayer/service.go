package relayer

import (
	"crypto/tls"
	"github.com/celer-network/goutils/log"
	rfqserver "github.com/celer-network/peti-rfq-relayer/relayer/service/rfq"
	rfqproto "github.com/celer-network/peti-rfq-relayer/relayer/service/rfq/proto"
	"github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm"
	rfqmmproto "github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	LPConfig         = "lp"
	ChainConfig      = "multichain"
	FeeConfig        = "fee"
	PriceProviderUrl = "priceprovider.url"
	RfqServerUrl     = "rfqserver.url"
	RequestSigner    = "requestsigner.chainid"
	RelayerConfig    = "relayer"
	MMConfig         = "mm"

	DefaultReportRetryPeriod int64 = 5
	DefaultProcessPeriod     int64 = 5
	DefaultPriceValidPeriod  int64 = 300
	DefaultDstTransferPeriod int64 = 3000
	DefaultPortListenOn      int64 = 5565
)

type RfqRelayerServer struct {
	Ctl                chan bool
	ClientPairMap      map[string]*ClientPair
	Config             *RfqRelayerConfig
	ChainCaller        rfqmm.ChainQuerier
	AmountCalculator   rfqmm.AmountCalculator
	DefaultLiqProvider *rfqmm.DefaultLiquidityProvider
	Client             rfqproto.RfqServerApiClient
}

type ClientPair struct {
	MmId            string
	ApiKey          string
	RfqServerClient rfqproto.RfqServerApiClient
	RfqMmClient     rfqmmproto.RfqMmApiClient
}

type RfqRelayerConfig struct {
	// the period for retrying report supported tokens to rfq server
	ReportRetryPeriod int64
	// the period for processing pending orders
	ProcessPeriod int64
	// indicates the period for a price to be valid
	PriceValidPeriod int64
	// minimum dst transfer period, in order to give mm enough time for dst transfer
	DstTransferPeriod int64
	// port num that mm would listen on
	Port int64
}

func (config *RfqRelayerConfig) clean() {
	if config.ReportRetryPeriod == 0 {
		config.ReportRetryPeriod = DefaultReportRetryPeriod
		log.Debugf("Got 0 ReportRetryPeriod, use default value(%d) instead.", DefaultReportRetryPeriod)
	}
	if config.ProcessPeriod == 0 {
		config.ProcessPeriod = DefaultProcessPeriod
		log.Debugf("Got 0 ProcessPeriod, use default value(%d) instead.", DefaultProcessPeriod)
	}
	if config.PriceValidPeriod == 0 {
		config.PriceValidPeriod = DefaultPriceValidPeriod
		log.Debugf("Got 0 PriceValidPeriod, use default value(%d) instead.", DefaultPriceValidPeriod)
	}
	if config.DstTransferPeriod == 0 {
		config.DstTransferPeriod = DefaultDstTransferPeriod
		log.Debugf("Got 0 DstTransferPeriod, use default value(%d) instead.", DefaultDstTransferPeriod)
	}

	if config.Port == 0 {
		config.Port = DefaultPortListenOn
		log.Debugf("Got 0 PortListenOn, use default value(%d) instead.", DefaultPortListenOn)
	}
}

func NewRfqRelayerServer() *RfqRelayerServer {
	// new Client of Rfq server
	rfqServerUrl := viper.GetString(RfqServerUrl)
	// new rfqmm clients
	var mmConfigs []*rfqmm.RfqMmConfig
	err := viper.UnmarshalKey(MMConfig, &mmConfigs)
	if err != nil {
		log.Fatalf("failed to load mm configs:%v", err)
		return nil
	}
	clientPairMap := newClients(rfqServerUrl, mmConfigs)

	// new Liquidity Manager
	var lpConfig rfqmm.LPConfig
	err = viper.UnmarshalKey(LPConfig, &lpConfig)
	if err != nil {
		log.Fatalf("failed to load liquidity-provider configs:%v", err)
		return nil
	}
	// lm := rfqmm.NewLiqManager(lpConfig)

	// // get a signer from Liquidity Manager by chain id
	// rsChainId := viper.GetUint64(RequestSigner)
	// requestSigner, _ := lm.GetLP(rsChainId)

	// new Chain Manager
	var chainConfig []*rfqmm.RfqMmChainConfig
	err = viper.UnmarshalKey(ChainConfig, &chainConfig)
	if err != nil {
		log.Fatalf("failed to load multichain configs:%v", err)
		return nil
	}
	cm := rfqmm.NewChainManager(chainConfig)

	// new default liquidity provider
	lp := rfqmm.NewDefaultLiquidityProvider(cm, &lpConfig)

	// new Amount Calculator
	feeConfig := new(rfqmm.FeeConfig)
	err = viper.UnmarshalKey(FeeConfig, feeConfig)
	if err != nil {
		log.Fatalf("failed to load mmfee configs:%v", err)
		return nil
	}
	// prepare a Price Provider
	priceUrl := viper.GetString(PriceProviderUrl)
	priceProvider := NewPriceProvider(priceUrl)
	priceProvider.UpdatePrice()
	ac := rfqmm.NewDefaultAmtCalculator(feeConfig, cm, priceProvider)

	// new Server
	serverConfig := new(RfqRelayerConfig)
	err = viper.UnmarshalKey(RelayerConfig, serverConfig)
	if err != nil {
		log.Fatalf("failed to load mm server configs:%v", err)
		return nil
	}

	serverConfig.clean()
	return &RfqRelayerServer{
		Ctl:                make(chan bool),
		ClientPairMap:      clientPairMap,
		Config:             serverConfig,
		ChainCaller:        cm,
		DefaultLiqProvider: lp,
		AmountCalculator:   ac,
	}
}

func newClients(rfqServerUrl string, mmConfigs []*rfqmm.RfqMmConfig) map[string]*ClientPair {
	if len(mmConfigs) == 0 {
		log.Fatal("mm config is empty")
	}

	clientPairMap := make(map[string]*ClientPair)
	for _, mmConfig := range mmConfigs {
		rfqServerClient := rfqserver.NewClient(rfqServerUrl, mmConfig.ApiKey, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
		opts := grpc.WithTransportCredentials(insecure.NewCredentials())
		rfqMmClient := rfqmm.NewClient(mmConfig.Endpoint, opts)
		clientPairMap[mmConfig.MmId] = &ClientPair{
			MmId:            mmConfig.MmId,
			ApiKey:          mmConfig.ApiKey,
			RfqServerClient: rfqServerClient,
			RfqMmClient:     rfqMmClient,
		}
	}

	return clientPairMap
}
