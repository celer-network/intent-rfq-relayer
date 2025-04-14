package db

import (
	"fmt"
	"math/big"
	"time"
)

type Relayer struct {
	QuoteHash      string
	SrcChainId     uint64
	SrcTokenSymbol string
	DstChainId     uint64
	DstTokenSymbol string
	RfqFee         *big.Int
	RfqFeeUsd      float64
	BaseFee        *big.Int
	BaseFeeUsd     float64
	CreateTime     time.Time
}

const (
	allRelayerColumns = "quote_hash, src_chain_id, src_token_symbol, dst_chain_id, dst_token_symbol, rfq_fee, " +
		"rfq_fee_usd, base_fee, base_fee_usd, create_time"
	allRelayerColumnParams = "$1, $2, $3, $4, $5, $6, $7, $8, $9, $10"
)

func (d *DAL) UpsertIntoRelayer(relayer *Relayer) error {
	q := fmt.Sprintf("INSERT INTO relayer (%s) VALUES (%s) ON CONFLICT DO NOTHING", allRelayerColumns, allRelayerColumnParams)
	_, err := d.Exec(q,
		relayer.QuoteHash,
		relayer.SrcChainId,
		relayer.SrcTokenSymbol,
		relayer.DstChainId,
		relayer.DstTokenSymbol,
		relayer.RfqFee.String(),
		relayer.RfqFeeUsd,
		relayer.BaseFee.String(),
		relayer.BaseFeeUsd,
		relayer.CreateTime)
	return err
}
