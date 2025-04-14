package relayer

import (
	"context"
	"google.golang.org/grpc/metadata"
	"math"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/peti-rfq-relayer/relayer/db"
	"github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm/proto"
)

func (s *RfqRelayerServer) Price(ctx context.Context, request *proto.PriceRequest) (response *proto.PriceResponse, err error) {
	apiKey, found := getApiKey(ctx)
	if !found {
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "unknown apiKey").ToCommonErr()}, nil
	}

	clientPair := s.ClientPairMap[apiKey]
	if nil == clientPair {
		log.Warnf("Price, unknown apiKey: %s", apiKey)
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "unknown apiKey").ToCommonErr()}, nil
	}
	log.Infof("Price, apiKey: %s, request: %v", apiKey, request)

	// calculate base fee and sub from src release amount
	baseFee, _, err := s.AmountCalculator.CalFixedCost(request.SrcToken, request.DstToken)
	if nil != err {
		log.Errorf("Price, fail to CalFixedCost, srcToken: %v, dstToken: %v, err: %v",
			request.SrcToken, request.DstToken, err)
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, err.Error()).ToCommonErr()}, nil
	}
	request.BaseAmount = baseFee.String()

	response, err = clientPair.RfqMmClient.Price(ctx, request)
	if nil != err || nil != response.Err {
		return response, err
	}
	log.Infof("Price, response: %v", response)

	// validate response src release amount
	srcAmount := new(big.Int)
	srcAmount.SetString(response.Price.SrcAmount, 10)
	rfqFee, err := s.ChainCaller.GetRfqFee(request.SrcToken.ChainId, request.DstToken.ChainId, srcAmount)
	if nil != err {
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "invalid rfq fee").ToCommonErr()}, nil
	}
	srcReleaseAmount := new(big.Int).Sub(new(big.Int).Sub(srcAmount, baseFee), rfqFee)
	resSrcReleaseAmount := new(big.Int)
	resSrcReleaseAmount.SetString(response.Price.SrcReleaseAmount, 10)
	if srcReleaseAmount.Cmp(resSrcReleaseAmount) != 0 {
		log.Warnf("srcAmount: %s, srcReleaseAmount: %s, baseFee: %s, rfqFee: %s, resSrcReleaseAmount: %s",
			srcAmount.String(), srcReleaseAmount.String(), baseFee.String(), rfqFee.String(), resSrcReleaseAmount.String())
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "invalid src release amount").ToCommonErr()}, nil
	}

	return response, nil
}

func (s *RfqRelayerServer) Quote(ctx context.Context, request *proto.QuoteRequest) (response *proto.QuoteResponse, err error) {
	apiKey, found := getApiKey(ctx)
	if !found {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "unknown apiKey").ToCommonErr()}, nil
	}

	clientPair := s.ClientPairMap[apiKey]
	if nil == clientPair {
		log.Warnf("Quote, unknown apiKey: %s", apiKey)
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "unknown apiKey").ToCommonErr()}, nil
	}
	log.Infof("Quote, apiKey: %s, request: %v", apiKey, request)

	response, err = clientPair.RfqMmClient.Quote(ctx, request)
	if nil != err || response.Err != nil {
		return response, err
	}

	srcAmount := new(big.Int)
	srcAmount.SetString(request.Quote.SrcAmount, 10)
	srcReleaseAmount := new(big.Int)
	srcReleaseAmount.SetString(request.Quote.SrcReleaseAmount, 10)
	decimal := big.NewFloat(math.Pow(10, float64(request.Quote.SrcToken.Decimals)))
	tokenInPrice, err := s.AmountCalculator.PriceProvider.GetPrice(request.Quote.SrcToken)
	if nil != err {
		log.Errorf("Quote, fail to GetPrice, srcToken: %v, err: %v", request.Quote.SrcToken, err)
		return response, nil
	}

	// calculate rfq fee and record to db
	rfqFee, err := s.ChainCaller.GetRfqFee(request.Quote.SrcToken.ChainId, request.Quote.DstToken.ChainId, srcAmount)
	if nil != err {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "invalid rfq fee").ToCommonErr()}, nil
	}
	rfqFeeUsd, _ := new(big.Float).Quo(new(big.Float).SetInt(rfqFee), decimal).Float64()
	rfqFeeUsd *= tokenInPrice

	// calculate base fee and record to db
	baseFee := new(big.Int).Sub(new(big.Int).Sub(srcAmount, srcReleaseAmount), rfqFee)
	baseFeeUsd, _ := new(big.Float).Quo(new(big.Float).SetInt(baseFee), decimal).Float64()
	baseFeeUsd *= tokenInPrice

	err = s.Db.UpsertIntoRelayer(&db.Relayer{
		QuoteHash:      request.Quote.Hash,
		SrcChainId:     request.Quote.SrcToken.ChainId,
		SrcTokenSymbol: request.Quote.SrcToken.Symbol,
		DstChainId:     request.Quote.DstToken.ChainId,
		DstTokenSymbol: request.Quote.DstToken.Symbol,
		RfqFee:         rfqFee,
		RfqFeeUsd:      rfqFeeUsd,
		BaseFee:        baseFee,
		BaseFeeUsd:     baseFeeUsd,
		CreateTime:     time.Now()})
	if nil != err {
		log.Errorf("Quote, fail to UpsertIntoRelayer, err: %v", err)
	}

	log.Infof("Quote, response: %v", response)

	return response, nil
}

func (s *RfqRelayerServer) SignQuoteHash(ctx context.Context, request *proto.SignQuoteHashRequest) (*proto.SignQuoteHashResponse, error) {
	return &proto.SignQuoteHashResponse{}, nil
}

func (s *RfqRelayerServer) Sign(ctx context.Context, request *proto.SignRequest) (response *proto.SignResponse, err error) {
	return &proto.SignResponse{}, nil
}

func (s *RfqRelayerServer) Verify(ctx context.Context, request *proto.VerifyRequest) (response *proto.VerifyResponse, err error) {
	return &proto.VerifyResponse{}, nil
}

func (s *RfqRelayerServer) Tokens(ctx context.Context, request *proto.TokensRequest) (response *proto.TokensResponse, err error) {
	return &proto.TokensResponse{}, nil
}

func getApiKey(ctx context.Context) (string, bool) {
	mk, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	var apiKey string
	if val, ok := mk["apikey"]; ok {
		apiKey = val[0]
	} else {
		return "", false
	}

	return apiKey, true
}
