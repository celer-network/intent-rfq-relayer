package rfqmm

import (
	"math"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/peti-rfq-relayer/relayer/common"
)

var _ AmountCalculator = &DefaultAmtCalculator{}

type DefaultAmtCalculator struct {
	// fixed cost related fields
	DstGasCost uint64
	SrcGasCost uint64
	GasPrice   map[uint64]uint64

	// helper
	Querier       ChainQuerier
	PriceProvider PriceProvider
}

type FeeConfig struct {
	DstGasCost     uint64
	SrcGasCost     uint64
	PercGlobal     uint32
	ChainOverrides []*ChainOverride
	TokenOverrides []*TokenOverride
	GasPrices      []*GasPrice
}

type ChainOverride struct {
	SrcChainId, DstChainId uint64
	Perc                   uint32
}

type TokenOverride struct {
	SrcChainId, DstChainId uint64
	SrcToken, DstToken     string
	Perc                   uint32
}

type GasPrice struct {
	ChainId uint64
	Price   uint64
}

type PriceProvider interface {
	GetPrice(token *common.Token) (float64, error)
}

func NewDefaultAmtCalculator(feeConfig *FeeConfig, querier ChainQuerier, priceProvider PriceProvider) *DefaultAmtCalculator {
	ac := &DefaultAmtCalculator{
		DstGasCost:    feeConfig.DstGasCost,
		SrcGasCost:    feeConfig.SrcGasCost,
		GasPrice:      make(map[uint64]uint64),
		Querier:       querier,
		PriceProvider: priceProvider,
	}
	ac.SetGasPrice(feeConfig.GasPrices)
	return ac
}

func (ac *DefaultAmtCalculator) SetDstGasCost(gasCost uint64) {
	ac.DstGasCost = gasCost
}

func (ac *DefaultAmtCalculator) SetSrcGasCost(gasCost uint64) {
	ac.SrcGasCost = gasCost
}

func (ac *DefaultAmtCalculator) SetGasPrice(prices []*GasPrice) {
	for _, gasPrice := range prices {
		ac.GasPrice[gasPrice.ChainId] = gasPrice.Price
	}
}

// fixed cost is composed of 1. src gas cost(only cross chain swap) 2. dst gas cost 3. dst msg fee(only cross chain swap)
func (ac *DefaultAmtCalculator) CalFixedCost(tokenIn, tokenOut *common.Token) (fixedCost *big.Int, fixedCostUsd float64, err error) {
	chainIn := tokenIn.ChainId
	chainOut := tokenOut.ChainId
	tokenInPrice, err := ac.PriceProvider.GetPrice(tokenIn)
	if err != nil {
		return
	}

	// 1 src gas cost
	nativeIn, err := ac.Querier.GetNativeWrap(chainIn)
	if err != nil {
		return
	}
	nativeInPrice, err := ac.PriceProvider.GetPrice(nativeIn)
	if err != nil {
		return
	}
	srcGasPrice, err := ac.Querier.GetGasPrice(chainIn)
	if err != nil || srcGasPrice.Sign() == 0 {
		log.Warnf("Fail to get gas price on chain %d, err:%v", chainIn, err)
		srcGasPrice = big.NewInt(int64(ac.GasPrice[chainIn]))
	}
	srcGasCost := big.NewInt(int64(ac.SrcGasCost))
	// represent src gas cost by src token amount
	srcGasCostInIn := convertAmount(new(big.Int).Mul(srcGasCost, srcGasPrice), nativeInPrice, tokenInPrice, tokenIn.Decimals-nativeIn.Decimals)

	// 2 dst gas cost
	nativeOut, err := ac.Querier.GetNativeWrap(chainOut)
	if err != nil {
		return
	}
	nativeOutPrice, err := ac.PriceProvider.GetPrice(nativeOut)
	if err != nil {
		return
	}
	dstGasPrice, err := ac.Querier.GetGasPrice(chainOut)
	if err != nil || dstGasPrice.Sign() == 0 {
		log.Warnf("Fail to get gas price on chain %d, err:%v", chainOut, err)
		dstGasPrice = big.NewInt(int64(ac.GasPrice[chainOut]))
	}
	dstGasCost := big.NewInt(int64(ac.DstGasCost))
	// represent dst gas cost by src token amount
	dstGasCostInIn := convertAmount(new(big.Int).Mul(dstGasCost, dstGasPrice), nativeOutPrice, tokenInPrice, tokenIn.Decimals-nativeOut.Decimals)
	log.Infof("dstGasCost: %s, dstGasPrice: %s, nativeOutPrice: %f, tokenInPrice: %f, decimalDiff: %d, dstGasCostInIn: %s",
		dstGasCost.String(), dstGasPrice.String(), nativeOutPrice, tokenInPrice, tokenIn.Decimals-nativeOut.Decimals, dstGasCostInIn.String())

	// 3 dst msg fee
	msgFeeAmt, _ := ac.Querier.GetMsgFee(chainOut)
	// represent dst msg fee by src token amount
	msgFeeInIn := convertAmount(msgFeeAmt, nativeOutPrice, tokenInPrice, tokenIn.Decimals-nativeOut.Decimals)

	// sum all cost
	if chainIn != chainOut {
		// all of 3
		fixedCost = new(big.Int).Add(srcGasCostInIn, dstGasCostInIn)
		fixedCost.Add(fixedCost, msgFeeInIn)
	} else {
		fixedCost = new(big.Int).Set(dstGasCostInIn)
	}
	fixedCostUsd, _ = new(big.Float).Quo(new(big.Float).Mul(new(big.Float).SetInt(fixedCost), new(big.Float).SetFloat64(tokenInPrice)),
		big.NewFloat(math.Pow(10, float64(tokenIn.Decimals)))).Float64()

	return
}

// decimalDiff = tokenOut.decimal - tokenIn.decimal
func convertAmount(amountIn *big.Int, priceIn, priceOut float64, decimalDiff int32) *big.Int {
	value := new(big.Float).Mul(big.NewFloat(priceIn), new(big.Float).SetInt(amountIn))
	amountOut := new(big.Float).Quo(value, big.NewFloat(priceOut))
	amountOut.Mul(amountOut, big.NewFloat(math.Pow(10, float64(decimalDiff))))
	res, _ := amountOut.Int(nil)
	return res
}
