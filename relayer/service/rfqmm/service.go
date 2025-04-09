package rfqmm

import (
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/peti-rfq-relayer/relayer/bindings/rfq"
	"github.com/celer-network/peti-rfq-relayer/relayer/common"
	"github.com/celer-network/peti-rfq-relayer/relayer/eth"
	rfqserver "github.com/celer-network/peti-rfq-relayer/relayer/service/rfq"
	rfqproto "github.com/celer-network/peti-rfq-relayer/relayer/service/rfq/proto"
	"github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm/proto"
	"google.golang.org/grpc"
	"math/big"
)

const (
	DefaultReportRetryPeriod int64 = 5
	DefaultProcessPeriod     int64 = 5
	DefaultPriceValidPeriod  int64 = 300
	DefaultDstTransferPeriod int64 = 3000
	DefaultPortListenOn      int64 = 5555
)

type RfqMmConfig struct {
	MmId     string
	ApiKey   string
	Endpoint string
}

type Client struct {
	proto.ApiClient
	server string
	conn   *grpc.ClientConn
}

type Server struct {
	Ctl               chan bool
	RfqClient         *rfqserver.Client
	Config            *ServerConfig
	ChainCaller       ChainQuerier
	LiquidityProvider LiquidityProvider
	AmountCalculator  AmountCalculator
	RequestSigner     RequestSigner
}

type Update struct {
	Hash   eth.Hash
	Status rfqproto.OrderStatus
}

type ServerConfig struct {
	// the period for retrying report supported tokens to rfq server
	ReportRetryPeriod int64
	// the period for processing pending orders
	ProcessPeriod int64
	// indicates the period for a price to be valid
	PriceValidPeriod int64
	// minimum dst transfer period, in order to give mm enough time for dst transfer
	DstTransferPeriod int64
	// token pair policy list
	TPPolicyList []string
	// port num that mm would listen on
	PortListenOn int64
}

func (config *ServerConfig) clean() {
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
	if len(config.TPPolicyList) == 0 {
		log.Debugf("No token pair policy was given.")
	}
	if config.PortListenOn == 0 {
		config.PortListenOn = DefaultPortListenOn
		log.Debugf("Got 0 PortListenOn, use default value(%d) instead.", DefaultPortListenOn)
	}
}

type ChainQuerier interface {
	GetRfqFee(srcChainId, dstChainId uint64, amount *big.Int) (*big.Int, error)
	GetMsgFee(chainId uint64) (*big.Int, error)
	GetGasPrice(chainId uint64) (*big.Int, error)
	GetNativeWrap(chainId uint64) (*common.Token, error)
	GetERC20Balance(chainId uint64, token, account eth.Addr) (*big.Int, error)
	GetNativeBalance(chainId uint64, accoun eth.Addr) (*big.Int, error)
	GetQuoteStatus(chainId uint64, quoteHash eth.Hash) (uint8, error)
	VerifyRfqEvent(chainId uint64, tx eth.Hash, evName string) (bool, error)
}

type LiquidityProvider interface {
	// IsPaused returns whether the LiquidityProvider is paused or not
	IsPaused() bool
	// GetTokens returns a list of all supported tokens
	GetTokens() []*common.Token
	// SetupTokenPairs sets up supported token pairs based on a given policy list.
	SetupTokenPairs(policies []string)
	// HasTokenPair check if a given token pair is supported
	HasTokenPair(srcToken, dstToken *common.Token) bool
	// GetLiquidityProviderAddr returns the address of liquidity provider on specified chain
	GetLiquidityProviderAddr(chainId uint64) (eth.Addr, error)
	// AskForFreezing checks if there is sufficient liquidity for specified token on specified chain and returns freeze time
	AskForFreezing(chainId uint64, token eth.Addr, amount *big.Int, isNative bool) (int64, error)
	// FreezeLiquidity will freeze a certain liquidity for specified amount until specified timestamp with an index of hash.
	FreezeLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash, isNative bool) error
	// UnfreezeLiquidity will try to unfreeze a certain liquidity with specified hash.
	UnfreezeLiquidity(chainId uint64, hash eth.Hash) error
	// DstTransfer should send tx on dstChain to transfer dstToken to user
	DstTransfer(transferNative bool, _quote rfq.RFQQuote, opts ...ethutils.TxOption) (eth.Hash, error)
	// SrcRelease should send tx on srcChain to release srcToken to mm
	SrcRelease(_quote rfq.RFQQuote, _execMsgCallData []byte, opts ...ethutils.TxOption) (eth.Hash, error)
}

type AmountCalculator interface {
	CalRecvAmt(tokenIn, tokenOut *common.Token, amountIn *big.Int) (recvAmt, releaseAmt, fee *big.Int, err error)
	CalSendAmt(tokenIn, tokenOut *common.Token, amountOut *big.Int) (sendAmt, releaseAmt, fee *big.Int, err error)
}

type RequestSigner interface {
	Sign(data []byte) ([]byte, error)
	Verify(data, sig []byte) bool
}

func NewClient(server string, ops ...grpc.DialOption) *Client {
	conn, err := grpc.Dial(server, ops...)
	if err != nil {
		panic(err)
	}
	client := proto.NewApiClient(conn)
	return &Client{server: server, conn: conn, ApiClient: client}
}

func (c *Client) Close() {
	c.conn.Close()
}
