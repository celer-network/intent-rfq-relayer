package rfqmm

import (
	"bytes"
	"encoding/json"
	"google.golang.org/grpc"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/celer-network/peti-rfq-relayer/relayer/common"
	"github.com/celer-network/peti-rfq-relayer/relayer/eth"
	rfqproto "github.com/celer-network/peti-rfq-relayer/relayer/service/rfq/proto"
	"github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm/proto"
	"github.com/gogo/protobuf/jsonpb"
)

type RfqMmConfig struct {
	MmId         string
	ApiKey       string
	Endpoint     string
	HttpEndpoint string
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

type RfqMmHttpClient struct {
	*http.Client
	Url string
}

func NewHttpClient(endpoint string) *RfqMmHttpClient {
	return &RfqMmHttpClient{
		Url:    endpoint,
		Client: &http.Client{},
	}
}

func (c *RfqMmHttpClient) Price(request *proto.PriceRequest) (*proto.PriceResponse, error) {
	unmarshaler := jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	resMsg, err := c.requestServer("POST", request)
	if nil != err {
		return nil, err
	}

	response := &proto.PriceResponse{}
	err = unmarshaler.Unmarshal(strings.NewReader(resMsg), response)
	return response, err
}

func (c *RfqMmHttpClient) Quote(request *proto.QuoteRequest) (*proto.QuoteResponse, error) {
	unmarshaler := jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	resMsg, err := c.requestServer("POST", request)
	if nil != err {
		return nil, err
	}

	response := &proto.QuoteResponse{}
	err = unmarshaler.Unmarshal(strings.NewReader(resMsg), response)
	return response, err
}

func (c *RfqMmHttpClient) SignQuoteHash(request *proto.SignQuoteHashRequest) (*proto.SignQuoteHashResponse, error) {
	unmarshaler := jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	resMsg, err := c.requestServer("POST", request)
	if nil != err {
		return nil, err
	}

	response := &proto.SignQuoteHashResponse{}
	err = unmarshaler.Unmarshal(strings.NewReader(resMsg), response)
	return response, err
}

func (c *RfqMmHttpClient) Tokens(request *proto.TokensRequest) (*proto.TokensResponse, error) {
	unmarshaler := jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	resMsg, err := c.requestServer("POST", request)
	if nil != err {
		return nil, err
	}

	response := &proto.TokensResponse{}
	err = unmarshaler.Unmarshal(strings.NewReader(resMsg), response)
	return response, err
}

func (c *RfqMmHttpClient) requestServer(method string, request interface{}) (string, error) {
	requestMsg, err := json.Marshal(request)
	if nil != err {
		return "", err
	}

	body := bytes.NewBuffer(requestMsg)
	req, err := http.NewRequest(method, c.Url, body)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	if err != nil {
		return "", err
	}

	for retry := 0; retry < 3; retry++ {
		resp, err := c.Client.Do(req)
		if err != nil {
			time.Sleep(300 * time.Millisecond)
			continue
		}

		defer resp.Body.Close()
		res, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		return string(res), nil
	}

	return "", err
}
