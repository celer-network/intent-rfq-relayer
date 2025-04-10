package relayer

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/celer-network/peti-rfq-relayer/relayer/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"sync"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/peti-rfq-relayer/relayer/bindings/rfq"
	"github.com/celer-network/peti-rfq-relayer/relayer/eth"
	rfqserver "github.com/celer-network/peti-rfq-relayer/relayer/service/rfq"
	rfqproto "github.com/celer-network/peti-rfq-relayer/relayer/service/rfq/proto"
	"github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm"
	rfqmmproto "github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm/proto"
	"github.com/spf13/viper"
)

const (
	LPConfig         = "lp"
	ChainConfig      = "multichain"
	FeeConfig        = "fee"
	PriceProviderUrl = "priceprovider.url"
	RfqServerUrl     = "rfqserver.url"
	RelayerConfig    = "relayer"
	MMConfig         = "mm"
	DbDriver         = "postgres"
	DbFmt            = "postgresql://root@%s/rfq_relayer?sslmode=disable"
	DbPoolSize       = 20

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
	Db                 *db.DAL
}

type ClientPair struct {
	MmId            string
	ApiKey          string
	RfqServerClient rfqproto.MMApiClient
	RfqMmClient     rfqmmproto.ApiClient
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
	Port  int64
	DbUrl string
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

	log.Infoln("Initializing DB...")
	dbConn, err := db.NewDAL(DbDriver, fmt.Sprintf(DbFmt, serverConfig.DbUrl), DbPoolSize, 1)
	if err != nil {
		log.Fatalf("failed to new db, dbUrl: %s, err: %v", serverConfig.DbUrl, err)
		return nil
	}
	log.Infoln("Successfully initialize DB")

	return &RfqRelayerServer{
		Ctl:                make(chan bool),
		ClientPairMap:      clientPairMap,
		Config:             serverConfig,
		ChainCaller:        cm,
		DefaultLiqProvider: lp,
		AmountCalculator:   ac,
		Db:                 dbConn,
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
		clientPairMap[mmConfig.ApiKey] = &ClientPair{
			MmId:            mmConfig.MmId,
			ApiKey:          mmConfig.ApiKey,
			RfqServerClient: rfqServerClient,
			RfqMmClient:     rfqMmClient,
		}
	}

	return clientPairMap
}

func (s *RfqRelayerServer) ReportConfigs() {
	log.Infof("Start reporting token config to rfq server every %d seconds", s.Config.ReportRetryPeriod)

	for _, clientPair := range s.ClientPairMap {
		go s.report(clientPair)
	}

	go func() {
		t := time.NewTicker(time.Duration(s.Config.ReportRetryPeriod) * time.Second)
		for range t.C {
			// report tokens to rfq server
			for _, clientPair := range s.ClientPairMap {
				go s.report(clientPair)
			}
		}
	}()
}

func (s *RfqRelayerServer) report(clientPair *ClientPair) {
	var err error
	for i := 0; i < 3; i++ {
		response, err := clientPair.RfqMmClient.Tokens(context.Background(), &rfqmmproto.TokensRequest{})
		if nil != err {
			log.Warnf("report token config, fail to get tokens, mmId: %s, apiKey: %s, err: %v",
				clientPair.MmId, clientPair.ApiKey, err)
			time.Sleep(3 * time.Second)
			continue
		}
		if len(response.Tokens) == 0 {
			log.Errorf("report token config, no token config, mmId: %s, apiKey: %s", clientPair.MmId, clientPair.ApiKey)
			return
		}

		request := &rfqproto.UpdateConfigsRequest{Config: &rfqmmproto.Config{Tokens: response.Tokens}}
		_, err = clientPair.RfqServerClient.UpdateConfigs(context.Background(), request)
		if err != nil {
			log.Warnf("report token config, fail to UpdateConfigs, mmId: %s, apiKey: %s, err: %v", clientPair.MmId, clientPair.ApiKey, err)
			time.Sleep(3 * time.Second)
			continue
		}

		log.Infof("report token config succeeded, mmId: %s, apiKey: %s", clientPair.MmId, clientPair.ApiKey)
		break
	}

	if err != nil {
		log.Errorf("report token config fail, mmId: %s, apiKey: %s, err: %v", clientPair.MmId, clientPair.ApiKey, err)
	}
}

func (s *RfqRelayerServer) StopProcessing(reason string) {
	log.Infof("Stopping server from processing pending orders, because %s", reason)
	s.Ctl <- true
}

func (s *RfqRelayerServer) DefaultProcessOrder() {
	log.Infof("Start processing order every %d seconds", s.Config.ProcessPeriod)
	go func() {
		if s.Ctl == nil {
			log.Panicln("nil control channel")
		}
		ticker := time.NewTicker(time.Duration(s.Config.ProcessPeriod) * time.Second)
		for {
			select {
			case <-ticker.C:
				// check component's functionality
				if s.DefaultLiqProvider.IsPaused() {
					s.StopProcessing("liquidity provider is paused in some reason")
					continue
				}

				for _, clientPair := range s.ClientPairMap {
					go s.processOrders(clientPair)
				}
			case <-s.Ctl:
				return
			}
		}
	}()
}

func (s *RfqRelayerServer) processOrders(clientPair *ClientPair) {
	resp, err := clientPair.RfqServerClient.PendingOrders(context.Background(), &rfqproto.PendingOrdersRequest{})
	if err != nil {
		log.Errorf("PendingOrders err:%s", err)
		return
	}

	var wg sync.WaitGroup
	for _, pendingOrder := range resp.Orders {
		wg.Add(1)
		go func(order *rfqproto.PendingOrder) {
			defer wg.Done()
			s.processOrder(order, clientPair)
		}(pendingOrder)
	}
	wg.Wait()
}

func (s *RfqRelayerServer) processOrder(pendingOrder *rfqproto.PendingOrder, clientPair *ClientPair) {
	quote := pendingOrder.Quote
	quoteHash := quote.GetQuoteHash()
	if !clientPair.ValidateQuote(quote, eth.Hex2Bytes(pendingOrder.QuoteSig)) {
		log.Errorf("Invalid quote, quoteHash %x", quoteHash)
		return
	}

	switch pendingOrder.Status {
	case rfqproto.OrderStatus_STATUS_SRC_DEPOSITED:
		// 1. check dst deadline
		timestamp := time.Now().Unix()
		if quote.DstDeadline < timestamp {
			log.Infof("SrcDeposited order with hash %x has past dst deadline %s, now is %s.", quoteHash,
				time.Unix(quote.DstDeadline, 0).Format("2006-01-02 15:04:06"),
				time.Unix(timestamp, 0).Format("2006-01-02 15:04:06"))
			// s.unfreeze(quote)
			// same chain swap, update status to refund initiated
			if quote.GetSrcChainId() == quote.GetDstChainId() {
				s.updateOrder(clientPair, quoteHash, rfqproto.OrderStatus_STATUS_REFUND_INITIATED, "")
			}
			return
		}

		// 2. verify tx on src chain
		ok, err := s.ChainCaller.VerifyRfqEvent(quote.GetSrcChainId(), eth.Hex2Hash(pendingOrder.SrcDepositTxHash), rfq.EventNameSrcDeposited)
		if err != nil {
			log.Warnf("VerifyRfqEvent err:%s, quoteHash %x", err, quoteHash)
			return
		}
		if !ok {
			log.Errorf("[Serious] Quote(hash %x) with status SRC_DEPOSITED does not pass event verification", quoteHash)
			// s.unfreeze(quote)
			s.StopProcessing(fmt.Sprintf("the order with hash %x does not pass event verification", quoteHash))
			return
		}

		// 3. check quoteHash on src chain
		statusOnChain, err := s.ChainCaller.GetQuoteStatus(quote.GetSrcChainId(), quoteHash)
		if err != nil {
			log.Errorf("GetQuoteStatus err:%s, quoteHash %x", err, quoteHash)
			return
		}
		if statusOnChain != rfq.QuoteStatusSrcDeposited {
			log.Errorf("[Serious] Quote(hash %x) status on src chain is %s, expected %s", quoteHash, rfq.GetQuoteStatusName(statusOnChain), rfq.GetQuoteStatusName(rfq.QuoteStatusSrcDeposited))
			// s.unfreeze(quote)
			s.StopProcessing(fmt.Sprintf("the order with hash %x is not truly deposited on src chain while rfq server thought it is", quoteHash))
			return
		}

		// 4. get sig from mm
		response, err := clientPair.RfqMmClient.Sign(context.Background(), &rfqmmproto.SignRequest{Data: quote.GetQuoteHash().Bytes()})
		if nil != err {
			log.Errorf("Get Sign, fail to request rfq mm, mmId: %s, apiKey: %s, err: %v", clientPair.MmId, clientPair.ApiKey, err)
			return
		}
		if nil != response.Err {
			log.Errorf("Get Sign with fail result, mmId: %s, apiKey: %s, err: %v", clientPair.MmId, clientPair.ApiKey, response.Err)
			return
		}

		// 5. send dst transfer
		txHash, err := s.DefaultLiqProvider.DstTransfer(pendingOrder.DstNative, quote.ToQuoteOnChain(), response.Sig)
		if err != nil {
			log.Errorf("DstTransfer err:%s, quoteHash %x", err, quoteHash)
			return
		}
		log.Infof("DstTransfer sent with txHash %x, quoteHash %x", txHash, quoteHash)
		// 5. update order's status
		s.updateOrder(clientPair, quoteHash, rfqproto.OrderStatus_STATUS_MM_DST_EXECUTED, eth.Bytes2Hex(txHash.Bytes()))
	case rfqproto.OrderStatus_STATUS_DST_TRANSFERRED:
		// 1. send src release
		txHash, err := s.DefaultLiqProvider.SrcRelease(pendingOrder.DstNative, quote.ToQuoteOnChain(), pendingOrder.ExecMsgCallData)
		if err != nil {
			log.Errorf("SrcRelease err:%s, quoteHash %x", err, quoteHash)
			return
		}
		log.Infof("SrcRelease sent with txHash %x, quoteHash %x", txHash, quoteHash)
		// 2. update order's status
		s.updateOrder(clientPair, quoteHash, rfqproto.OrderStatus_STATUS_MM_SRC_EXECUTED, eth.Bytes2Hex(txHash.Bytes()))
	}
}

func (s *ClientPair) ValidateQuote(quote *rfqmmproto.Quote, sig []byte) bool {
	// 1 check sig
	response, err := s.RfqMmClient.Verify(context.Background(), &rfqmmproto.VerifyRequest{Data: quote.GetQuoteHash().Bytes(), Sig: sig})
	if nil != err || !response.GetVerified() {
		// serious error
		log.Errorf("[Serious] Invalid sig, quoteHash %x, err: %v", eth.Hex2Hash(quote.Hash), err)
		return false
	}
	// 2 check quote hash
	if quote.GetQuoteHash() != quote.EncodeQuoteHash() {
		// serious error
		log.Errorf("[Serious] Quote hash mismatch, got %x, calculated %x", eth.Hex2Hash(quote.Hash), quote.EncodeQuoteHash())
		return false
	}
	return true
}

func (s *RfqRelayerServer) updateOrder(clientPair *ClientPair, quoteHash eth.Hash, toStatus rfqproto.OrderStatus, txHash string) {
	_, err := clientPair.RfqServerClient.UpdateOrders(context.Background(),
		&rfqproto.UpdateOrdersRequest{
			Updates: []*rfqproto.OrderUpdates{{QuoteHash: quoteHash.String(), OrderStatus: toStatus, ExecTxHash: txHash}},
		})
	if err != nil {
		log.Errorf("UpdateOrders err:%s, quoteHash %x, toStatus %s, txHash %s", err, quoteHash, toStatus, txHash)
	} else {
		log.Infof("Order updated, quoteHash %x, toStatus %s, txHash %s", quoteHash, toStatus, txHash)
	}
}

func (s *RfqRelayerServer) Serve(ops ...grpc.ServerOption) {
	log.Infof("Start relayer server, listen on port %d", s.Config.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.Config.Port))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(ops...)
	rfqmmproto.RegisterApiServer(grpcServer, s)
	grpcServer.Serve(lis)
}
