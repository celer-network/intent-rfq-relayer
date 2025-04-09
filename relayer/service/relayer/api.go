package relayer

import (
	"context"
	"github.com/celer-network/goutils/log"
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

	return clientPair.RfqMmClient.Price(ctx, request)
}

func (s *RfqRelayerServer) Quote(ctx context.Context, request *proto.QuoteRequest) (response *proto.QuoteResponse, err error) {
	apiKey := ctx.Value(apiKey).(string)
	clientPair := s.ClientPairMap[apiKey]
	if nil == clientPair {
		log.Warnf("Price, unknown apiKey: %s", apiKey)
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_UNDEFINED, "unknown apiKey").ToCommonErr()}, nil
	}

	return clientPair.RfqMmClient.Quote(ctx, request)
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
