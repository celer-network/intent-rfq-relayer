package relayer

import (
	"context"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/peti-rfq-relayer/relayer/db"
	"github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm/proto"
)

const apiKey string = "apiKey"

func (s *RfqRelayerServer) Price(ctx context.Context, request *proto.PriceRequest) (response *proto.PriceResponse, err error) {
	apiKey := ctx.Value(apiKey).(string)
	clientPair := s.ClientPairMap[apiKey]
	if nil == clientPair {
		log.Warnf("Price, unknown apiKey: %s", apiKey)
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "unknown apiKey").ToCommonErr()}, nil
	}
	log.Infof("Price, apiKey: %s, request: %v", apiKey, request)

	response, err = clientPair.RfqMmClient.Price(ctx, request)
	if nil != err || response.Err != nil {
		return response, err
	}

	// calculate base fee and sub from src release amount
	baseFee, _, err := s.AmountCalculator.CalFixedCost(request.SrcToken, request.DstToken)
	if nil != err {
		log.Errorf("Price, fail to CalFixedCost, srcToken: %v, dstToken: %v, err: %v",
			request.SrcToken, request.DstToken, err)
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, err.Error()).ToCommonErr()}, nil
	}
	srcReleaseAmount := new(big.Int)
	srcReleaseAmount.SetString(response.Price.SrcReleaseAmount, 10)
	srcReleaseAmount = new(big.Int).Sub(srcReleaseAmount, baseFee)
	response.Price.SrcReleaseAmount = srcReleaseAmount.String()

	return response, nil
}

func (s *RfqRelayerServer) Quote(ctx context.Context, request *proto.QuoteRequest) (response *proto.QuoteResponse, err error) {
	apiKey := ctx.Value(apiKey).(string)
	clientPair := s.ClientPairMap[apiKey]
	if nil == clientPair {
		log.Warnf("Price, unknown apiKey: %s", apiKey)
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "unknown apiKey").ToCommonErr()}, nil
	}
	log.Infof("Quote, apiKey: %s, request: %v", apiKey, request)

	response, err = clientPair.RfqMmClient.Quote(ctx, request)
	if nil != err || response.Err != nil {
		return response, err
	}

	// calculate base fee and record to db
	baseFee, baseFeeUsd, err := s.AmountCalculator.CalFixedCost(request.Quote.SrcToken, request.Quote.DstToken)
	if nil != err {
		log.Errorf("Price, fail to CalFixedCost, srcToken: %v, dstToken: %v, err: %v",
			request.Quote.SrcToken, request.Quote.DstToken, err)
		return response, nil
	}
	err = s.Db.UpsertIntoRelayer(&db.Relayer{
		QuoteHash:      request.Quote.Hash,
		SrcChainId:     request.Quote.SrcToken.ChainId,
		SrcTokenSymbol: request.Quote.SrcToken.Symbol,
		DstChainId:     request.Quote.DstToken.ChainId,
		DstTokenSymbol: request.Quote.DstToken.Symbol,
		BaseFee:        baseFee,
		BaseFeeUsd:     baseFeeUsd,
		CreateTime:     time.Now()})
	if nil != err {
		log.Errorf("Quote, fail to UpsertIntoRelayer, err: %v", err)
	}

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
