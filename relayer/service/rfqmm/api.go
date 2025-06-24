package rfqmm

import (
	"context"
	"math/big"
	"time"

	"google.golang.org/grpc"

	"github.com/celer-network/intent-rfq-relayer/relayer/service/rfqmm/proto"
)

func (c *Client) Price(ctx context.Context, in *proto.PriceRequest, opts ...grpc.CallOption) (*proto.PriceResponse, error) {
	if ok, reason := validatePriceRequest(in); !ok {
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, reason).ToCommonErr()}, nil
	}
	return c.ApiClient.Price(ctx, in, opts...)
}

func (c *Client) Quote(ctx context.Context, in *proto.QuoteRequest, opts ...grpc.CallOption) (*proto.QuoteResponse, error) {
	if ok, reason := validateQuoteRequest(in); !ok {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, reason).ToCommonErr()}, nil
	}
	return c.ApiClient.Quote(ctx, in, opts...)
}

func validatePriceRequest(request *proto.PriceRequest) (bool, string) {
	if request.SrcToken == nil || request.DstToken == nil {
		return false, "SrcToken or DstToken is nil"
	}
	if request.SrcAmount == "" && request.DstAmount == "" {
		return false, "SrcAmount and DstAmount are both empty"
	}
	if request.SrcAmount == "" {
		if _, ok := new(big.Int).SetString(request.DstAmount, 10); !ok {
			return false, "invalid SrcAmount"
		}
	} else {
		if _, ok := new(big.Int).SetString(request.SrcAmount, 10); !ok {
			return false, "invalid DstAmount"
		}
	}
	return true, ""
}

func validateQuoteRequest(request *proto.QuoteRequest) (bool, string) {
	price := request.Price
	quote := request.Quote
	if request.Price == nil || request.Quote == nil {
		return false, "price or quote is nil"
	}
	if price.SrcToken == nil || price.DstToken == nil {
		return false, "price.SrcToken or price.DstToken is nil"
	}
	if price.SrcAmount == "" || price.DstAmount == "" || price.SrcReleaseAmount == "" {
		return false, "price.SrcAmount, price.DstAmount or price.SrcReleaseAmount is empty"
	}
	if time.Now().Unix() > price.ValidThru {
		return false, "past price valid time"
	}
	if price.GetMMAddr() != quote.GetMMAddr() {
		return false, "mm addr mismatch"
	}
	if !quote.SrcToken.EqualBasically(price.SrcToken) || !quote.DstToken.EqualBasically(price.DstToken) {
		return false, "token in price and quote mismatch"
	}
	if quote.SrcAmount != price.SrcAmount || quote.DstAmount != price.DstAmount || quote.SrcReleaseAmount != price.SrcReleaseAmount {
		return false, "amount in price and quote mismatch"
	}
	if quote.Sender == "" || quote.Receiver == "" || quote.MmAddr == "" {
		return false, "quote.Sender, quote.Receiver or quote.MmAddr is empty"
	}
	if time.Now().Unix() > quote.SrcDeadline {
		return false, "past src deadline"
	}
	if quote.DstDeadline < quote.SrcDeadline {
		return false, "dst deadline is earlier than src deadline"
	}
	if !quote.ValidateQuoteHash() {
		return false, "quote hahs mismatch"
	}
	return true, ""
}
