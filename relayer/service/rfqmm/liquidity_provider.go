package rfqmm

import (
	"fmt"
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/peti-rfq-relayer/relayer/bindings/rfq"
	"github.com/celer-network/peti-rfq-relayer/relayer/eth"
	"github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm/proto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	dstTransfer          = "DstTransfer"
	srcRelease           = "SrcRelease"
	sameChainTransfer    = "SameChainTransfer"
	NativeTokenReference = "ffffffffffffffffffffffffffffffffffffffff"
)

type DefaultLiquidityProvider struct {
	paused       bool
	txrs         map[uint64]*ethutils.Transactor
	chainManager *ChainManager
	liqManager   *LiqManager
	liqConfig    *LPConfig
	tokenPair    map[string]bool
}

func NewDefaultLiquidityProvider(cm *ChainManager, config *LPConfig) *DefaultLiquidityProvider {
	lp := &DefaultLiquidityProvider{
		paused:       false,
		txrs:         make(map[uint64]*ethutils.Transactor),
		chainManager: cm,
		liqConfig:    config,
		tokenPair:    make(map[string]bool),
	}
	// construct transactor for each chain
	for chainId := range cm.chains {
		signer, addr, err := createSigner(config.Keystore, config.Passphrase, big.NewInt(int64(chainId)))
		if err != nil {
			panic(err)
		}
		chain, err := cm.GetChain(chainId)
		if err != nil {
			log.Errorf("GetChain err:%s", err)
			continue
		}
		lp.txrs[chainId] = ethutils.NewTransactorByExternalSigner(addr, signer, chain.Client, big.NewInt(int64(chain.ChainId)), chain.TxOptions...)
	}

	return lp
}

func (d *DefaultLiquidityProvider) IsPaused() bool {
	return d.paused
}

func (d *DefaultLiquidityProvider) GetContractAddress(chainId uint64) (string, bool) {
	chain, err := d.chainManager.GetChain(chainId)
	if nil != err {
		return "", false
	}
	return chain.RfqAddress.String(), true
}

func (d *DefaultLiquidityProvider) DstTransfer(transferNative bool, _quote rfq.RFQQuote, sig []byte, opts ...ethutils.TxOption) (eth.Hash, error) {
	if d.paused {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, "liquidity provider is paused due to some serious error")
	}

	log.Infof("DstTransfer, sig: %x, %v", sig, sig)

	// check if it's a same chain swap
	if _quote.DstChainId == _quote.SrcChainId {
		return d.sameChainTransfer(transferNative, _quote, sig, opts...)
	}
	chain, err := d.chainManager.GetChain(_quote.DstChainId)
	if err != nil {
		return eth.ZeroHash, err
	}
	txr, ok := d.txrs[_quote.DstChainId]
	if !ok {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("no transactor for chain %d", _quote.DstChainId))
	}
	var method ethutils.TxMethod
	if transferNative {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.DstTransferNative(opts, _quote)
		}
		opts = append(opts, ethutils.WithEthValue(new(big.Int).Add(_quote.DstAmount, chain.MsgFee)))
	} else {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.DstTransferWithSig(opts, _quote, sig)
		}
		opts = append(opts, ethutils.WithEthValue(chain.MsgFee))
	}
	tx, err := txr.Transact(d.genTxHandler(dstTransfer, _quote), method, opts...)
	if err != nil {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("Transact err:%s", err))
	}
	return tx.Hash(), nil
}

func (d *DefaultLiquidityProvider) SrcRelease(transferNative bool, _quote rfq.RFQQuote, _execMsgCallData []byte, opts ...ethutils.TxOption) (eth.Hash, error) {
	if d.paused {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, "liquidity provider is paused due to some serious error")
	}
	chain, err := d.chainManager.GetChain(_quote.SrcChainId)
	if err != nil {
		return eth.ZeroHash, err
	}
	txr, ok := d.txrs[_quote.SrcChainId]
	if !ok {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("no transactor for chain %d", _quote.SrcChainId))
	}
	// determine release native or not
	releaseNative := transferNative
	// if chain.NativeWrap.GetAddr() == _quote.SrcToken {
	// 	releaseNative, err = d.liqManager.ReleaseNative(_quote.SrcChainId)
	// 	if err != nil {
	// 		return eth.ZeroHash, err
	// 	}
	// }
	var method ethutils.TxMethod
	if releaseNative {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.SrcReleaseNative(opts, _quote, _execMsgCallData)
		}
	} else {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.SrcRelease(opts, _quote, _execMsgCallData)
		}
	}
	tx, err := txr.Transact(d.genTxHandler(srcRelease, _quote), method, opts...)
	if err != nil {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("Transact err:%s", err))
	}
	return tx.Hash(), nil
}

func (d *DefaultLiquidityProvider) sameChainTransfer(transferNative bool, _quote rfq.RFQQuote, sig []byte, opts ...ethutils.TxOption) (eth.Hash, error) {
	chain, err := d.chainManager.GetChain(_quote.DstChainId)
	if err != nil {
		return eth.ZeroHash, err
	}
	// confirm liquidity before same chain transfer
	txr, ok := d.txrs[_quote.DstChainId]
	if !ok {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("no transactor for chain %d", _quote.DstChainId))
	}
	// determine release native or not
	releaseNative := transferNative
	// if chain.NativeWrap.GetAddr() == _quote.SrcToken {
	// 	releaseNative, err = d.liqManager.ReleaseNative(_quote.SrcChainId)
	// 	if err != nil {
	// 		return eth.ZeroHash, err
	// 	}
	// }
	var method ethutils.TxMethod
	if transferNative {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.SameChainTransferNative(opts, _quote, releaseNative)
		}
		opts = append(opts, ethutils.WithEthValue(_quote.DstAmount))
	} else {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.SameChainTransferWithSig(opts, _quote, releaseNative, sig)
		}
	}
	tx, err := txr.Transact(
		d.genTxHandler(sameChainTransfer, _quote), method, opts...)
	if err != nil {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("Transact err:%s", err))
	}
	return tx.Hash(), nil
}

func (d *DefaultLiquidityProvider) genTxHandler(methodName string, _quote rfq.RFQQuote) *ethutils.TransactionStateHandler {
	quoteHash := _quote.Hash()
	switch methodName {
	case dstTransfer:
		return &ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("%s succeeded, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
				} else {
					log.Errorf("%s failed, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
					d.pause()
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("%s err: %s. quote hash %s", methodName, err, quoteHash)
			},
		}
	case srcRelease:
		return &ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("%s succeeded, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
				} else {
					log.Errorf("%s failed, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
					d.pause()
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("%s err: %s. quote hash %s", methodName, err, quoteHash)
			},
		}
	case sameChainTransfer:
		return &ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("%s succeeded, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
				} else {
					log.Errorf("%s failed, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
					d.pause()
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("%s err: %s. quote hash %s", methodName, err, quoteHash)
			},
		}
	default:
		return nil
	}
}

func (d *DefaultLiquidityProvider) pause() {
	d.paused = true
}
