package rfqmm

import (
	"google.golang.org/grpc"
	"math/big"

	"github.com/celer-network/peti-rfq-relayer/relayer/common"
	"github.com/celer-network/peti-rfq-relayer/relayer/eth"
	rfqproto "github.com/celer-network/peti-rfq-relayer/relayer/service/rfq/proto"
	"github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm/proto"
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

type Update struct {
	Hash   eth.Hash
	Status rfqproto.OrderStatus
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

type AmountCalculator interface {
	CalFixedCost(tokenIn, tokenOut *common.Token) (fixedCost *big.Int, fixedCostUsd float64, err error)
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
