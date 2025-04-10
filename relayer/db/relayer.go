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
	BaseFee        *big.Int
	BaseFeeUsd     float64
	CreateTime     time.Time
}

const (
	allRelayerColumns      = "quote_hash, src_chain_id, src_token_symbol, dst_chain_id, dst_token_symbol, base_fee, base_fee_usd, create_time"
	allRelayerColumnParams = "$1, $2, $3, $4, $5, $6, $7, $8"
)

func (d *DAL) UpsertIntoRelayer(relayer *Relayer) error {
	q := fmt.Sprintf("INSERT INTO relayer (%s) VALUES (%s) ON CONFLICT DO NOTHING", allRelayerColumns, allRelayerColumnParams)
	_, err := d.Exec(q,
		relayer.QuoteHash,
		relayer.SrcChainId,
		relayer.SrcTokenSymbol,
		relayer.DstChainId,
		relayer.DstTokenSymbol,
		relayer.BaseFee.String(),
		relayer.BaseFeeUsd,
		relayer.CreateTime)
	return err
}
