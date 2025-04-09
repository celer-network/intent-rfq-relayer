// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rfq

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RFQQuote is an auto generated low-level Go binding around an user-defined struct.
type RFQQuote struct {
	SrcChainId        uint64
	SrcToken          common.Address
	SrcAmount         *big.Int
	SrcReleaseAmount  *big.Int
	DstChainId        uint64
	DstToken          common.Address
	DstAmount         *big.Int
	Deadline          uint64
	Nonce             uint64
	Sender            common.Address
	Receiver          common.Address
	RefundTo          common.Address
	LiquidityProvider common.Address
}

// RfqMetaData contains all meta data concerning the Rfq contract.
var RfqMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DstTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"chainIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint32[]\",\"name\":\"feePercs\",\"type\":\"uint32[]\"}],\"name\":\"FeePercUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"}],\"name\":\"MessageBusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"}],\"name\":\"RefundInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"chainIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"remoteRfqContracts\",\"type\":\"address[]\"}],\"name\":\"RfqContractsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structRFQ.Quote\",\"name\":\"quote\",\"type\":\"tuple\"}],\"name\":\"SrcDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SrcReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"}],\"name\":\"TreasuryAddrUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"}],\"name\":\"dstTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"}],\"name\":\"dstTransferNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"dstTransferWithSig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_execMsgCallData\",\"type\":\"bytes\"}],\"name\":\"executeRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_execMsgCallData\",\"type\":\"bytes\"}],\"name\":\"executeRefundNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercGlobal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"feePercOverride\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"getMsgFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"}],\"name\":\"getQuoteHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getRfqFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_quoteHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"getSignerOfQuoteHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"quotes\",\"outputs\":[{\"internalType\":\"enumRFQ.QuoteStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"registerAllowedSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"remoteRfqContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"}],\"name\":\"requestRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_releaseNative\",\"type\":\"bool\"}],\"name\":\"sameChainTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_releaseNative\",\"type\":\"bool\"}],\"name\":\"sameChainTransferNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_releaseNative\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"sameChainTransferWithSig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_chainIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_feePercs\",\"type\":\"uint32[]\"}],\"name\":\"setFeePerc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeWrap\",\"type\":\"address\"}],\"name\":\"setNativeWrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_chainIds\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"_remoteRfqContracts\",\"type\":\"address[]\"}],\"name\":\"setRemoteRfqContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"}],\"name\":\"setTreasuryAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_submissionDeadline\",\"type\":\"uint64\"}],\"name\":\"srcDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_submissionDeadline\",\"type\":\"uint64\"}],\"name\":\"srcDepositNative\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_execMsgCallData\",\"type\":\"bytes\"}],\"name\":\"srcRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_execMsgCallData\",\"type\":\"bytes\"}],\"name\":\"srcReleaseNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unconsumedMsg\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_quoteHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"verifySigOfQuoteHash\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620055f5380380620055f5833981016040819052620000349162000262565b6200003f3362000088565b6001805460ff60a01b191690556200005733620000d8565b6200006233620001a2565b600180546001600160a01b0319166001600160a01b039290921691909117905562000294565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526002602052604090205460ff1615620001475760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064015b60405180910390fd5b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f891015b60405180910390a150565b6001600160a01b03811660009081526003602052604090205460ff16156200020d5760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f72000000000060448201526064016200013e565b6001600160a01b038116600081815260036020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5910162000197565b6000602082840312156200027557600080fd5b81516001600160a01b03811681146200028d57600080fd5b9392505050565b61535180620002a46000396000f3fe6080604052600436106103645760003560e01c806382dc1ec4116101c6578063bbf3b98f116100f7578063e3eece2611610095578063eecdac881161006f578063eecdac8814610a0a578063f2fde38b14610a2a578063f8b839e514610a4a578063fbe42fea14610a6a57600080fd5b8063e3eece261461098e578063e43581b8146109be578063ed9830bb146109f757600080fd5b8063cc47e400116100d1578063cc47e400146108f3578063d1ce8b4514610929578063df1f64ef14610949578063e026049c1461097957600080fd5b8063bbf3b98f1461088a578063c78e33a2146108c0578063cbac44df146108d357600080fd5b8063a5ea10cb11610164578063ae6c218e1161013e578063ae6c218e14610817578063af4ab1381461082a578063b1a7d13f1461084a578063b62b31e41461085d57600080fd5b8063a5ea10cb146107c4578063a7e05b9c146107d7578063ab9341fd146107f757600080fd5b80639c649fdf116101a05780639c649fdf14610751578063a1a227fa14610764578063a2bdb89914610784578063a2f6531e146107a457600080fd5b806382dc1ec4146106fe5780638456cb591461071e5780638da5cb5b1461073357600080fd5b806346fbf68e116102a05780636609870d1161023e5780636ef8d66d116102185780636ef8d66d146106a657806379c7efd3146106bb5780637cd2bffc146105f457806380f51c12146106ce57600080fd5b80636609870d1461064657806369b59e75146106665780636b2c0f551461068657600080fd5b80635aa15a4b1161027a5780635aa15a4b146105d45780635ab7afc6146105f45780635b5a66a7146106075780635c975abb1461062757600080fd5b806346fbf68e1461054b578063547cad121461059457806355e9e3d2146105b457600080fd5b806325329eaf1161030d5780633c4a25d0116102e75780633c4a25d0146104c35780633e07d172146104e35780633f4ba83a14610516578063457bfa2f1461052b57600080fd5b806325329eaf1461042e5780632a37cf711461046b57806330d9a62a1461048b57600080fd5b80630bd930b41161033e5780630bd930b4146103cd5780631000cd9e146104065780631e9c57481461041b57600080fd5b8063063ce4e5146103705780630a54aacd146103995780630bcb4982146103ba57600080fd5b3661036b57005b600080fd5b61038361037e3660046147c3565b610a8a565b604051610390919061487b565b60405180910390f35b6103ac6103a73660046148a7565b610af6565b604051908152602001610390565b6103836103c83660046148dd565b610cf7565b3480156103d957600080fd5b506008546103f190600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610390565b61041961041436600461494b565b610d5d565b005b610419610429366004614976565b610f9f565b34801561043a57600080fd5b5061045e6104493660046149b0565b60076020526000908152604090205460ff1681565b60405161039091906149c9565b34801561047757600080fd5b506104196104863660046149dd565b6111eb565b34801561049757600080fd5b506008546104ab906001600160a01b031681565b6040516001600160a01b039091168152602001610390565b3480156104cf57600080fd5b506104196104de366004614a36565b6112a3565b3480156104ef57600080fd5b506103f16104fe366004614a51565b60096020526000908152604090205463ffffffff1681565b34801561052257600080fd5b50610419611306565b34801561053757600080fd5b506004546104ab906001600160a01b031681565b34801561055757600080fd5b50610584610566366004614a36565b6001600160a01b031660009081526002602052604090205460ff1690565b6040519015158152602001610390565b3480156105a057600080fd5b506104196105af366004614a36565b61136f565b3480156105c057600080fd5b506103ac6105cf366004614a6c565b61141b565b3480156105e057600080fd5b506104196105ef366004614a36565b611496565b610383610602366004614aad565b6114f1565b34801561061357600080fd5b50610419610622366004614a36565b611559565b34801561063357600080fd5b50600154600160a01b900460ff16610584565b34801561065257600080fd5b50610419610661366004614b3c565b6115d2565b34801561067257600080fd5b50610419610681366004614a36565b611747565b34801561069257600080fd5b506104196106a1366004614a36565b611820565b3480156106b257600080fd5b50610419611880565b6104196106c936600461494b565b611889565b3480156106da57600080fd5b506105846106e9366004614a36565b60026020526000908152604090205460ff1681565b34801561070a57600080fd5b50610419610719366004614a36565b611b5d565b34801561072a57600080fd5b50610419611bbd565b34801561073f57600080fd5b506000546001600160a01b03166104ab565b61038361075f366004614b91565b611c24565b34801561077057600080fd5b506001546104ab906001600160a01b031681565b34801561079057600080fd5b5061041961079f366004614b3c565b611d44565b3480156107b057600080fd5b506104ab6107bf366004614bdb565b611db2565b6104196107d236600461494b565b611e6e565b3480156107e357600080fd5b506104196107f2366004614a36565b611f77565b34801561080357600080fd5b50610419610812366004614c51565b61201c565b610419610825366004614b3c565b6122a2565b34801561083657600080fd5b50610419610845366004614b3c565b612444565b610419610858366004614cb0565b61251e565b34801561086957600080fd5b506103ac610878366004614a36565b600a6020526000908152604090205481565b34801561089657600080fd5b506104ab6108a5366004614a36565b600b602052600090815260409020546001600160a01b031681565b6104196108ce366004614976565b6126cb565b3480156108df57600080fd5b506104196108ee366004614c51565b6127cc565b3480156108ff57600080fd5b506104ab61090e366004614a51565b6005602052600090815260409020546001600160a01b031681565b34801561093557600080fd5b506103ac61094436600461494b565b61294e565b34801561095557600080fd5b506105846109643660046149b0565b60066020526000908152604090205460ff1681565b34801561098557600080fd5b50610419612a42565b34801561099a57600080fd5b506105846109a9366004614a36565b60036020526000908152604090205460ff1681565b3480156109ca57600080fd5b506105846109d9366004614a36565b6001600160a01b031660009081526003602052604090205460ff1690565b6103ac610a053660046148a7565b612a4b565b348015610a1657600080fd5b50610419610a25366004614a36565b612ac3565b348015610a3657600080fd5b50610419610a45366004614a36565b612b23565b348015610a5657600080fd5b50610419610a65366004614b3c565b612bff565b348015610a7657600080fd5b506103ac610a85366004614cfe565b612ca7565b6001546000906001600160a01b03163314610aec5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b9695505050505050565b600154600090600160a01b900460ff1615610b465760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b6004546001600160a01b0316610b9e5760405162461bcd60e51b815260206004820152601860248201527f5266713a206e61746976652077726170206e6f742073657400000000000000006044820152606401610ae3565b6004546001600160a01b0316610bba6040850160208601614a36565b6001600160a01b031614610c105760405162461bcd60e51b815260206004820152601760248201527f5266713a2073726320746f6b656e206d69736d617463680000000000000000006044820152606401610ae3565b8260400135341015610c645760405162461bcd60e51b815260206004820152601860248201527f5266713a20696e73756666696369656e7420616d6f756e7400000000000000006044820152606401610ae3565b6000610c7e8484610c79604083013534614d3e565b612d05565b9050600460009054906101000a90046001600160a01b03166001600160a01b031663d0e30db085604001356040518263ffffffff1660e01b81526004016000604051808303818588803b158015610cd457600080fd5b505af1158015610ce8573d6000803e3d6000fd5b50939450505050505b92915050565b6001546000906001600160a01b03163314610d545760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610ae3565b95945050505050565b600154600160a01b900460ff1615610daa5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b6004546001600160a01b0316610dc660c0830160a08401614a36565b6001600160a01b031614610e1c5760405162461bcd60e51b815260206004820152601760248201527f5266713a2064737420746f6b656e206d69736d617463680000000000000000006044820152606401610ae3565b8060c00135341015610e705760405162461bcd60e51b815260206004820152601860248201527f5266713a20696e73756666696369656e7420616d6f756e7400000000000000006044820152606401610ae3565b600080610e7c83613175565b6000828152600760209081526040808320805460ff191660081790555193955091935091610eaf91859160019101614d51565b60408051601f198184030181528282528051602091820120908301520160408051601f198184030181529190529050610f0382610eef6020870187614a51565b83610efe60c089013534614d3e565b61340f565b610f22610f1861016086016101408701614a36565b8560c0013561342b565b7fb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc583610f5661016087016101408801614a36565b610f6660c0880160a08901614a36565b604080519384526001600160a01b03928316602085015291169082015260c086013560608201526080015b60405180910390a150505050565b600154600160a01b900460ff1615610fec5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b610ffc60a0830160808401614a51565b6001600160401b03166110126020840184614a51565b6001600160401b0316146110685760405162461bcd60e51b815260206004820152601860248201527f5266713a206e6f742073616d6520636861696e207377617000000000000000006044820152606401610ae3565b6004546001600160a01b031661108460c0840160a08501614a36565b6001600160a01b0316146110da5760405162461bcd60e51b815260206004820152601760248201527f5266713a2064737420746f6b656e206d69736d617463680000000000000000006044820152606401610ae3565b8160c0013534146111375760405162461bcd60e51b815260206004820152602160248201527f5266713a206e617469766520746f6b656e20616d6f756e74206d69736d6174636044820152600d60fb1b6064820152608401610ae3565b600061114283613175565b50905061116461115a61016085016101408601614a36565b8460c0013561342b565b61116f83828461353c565b7fb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc5816111a361016086016101408701614a36565b6111b360c0870160a08801614a36565b604080519384526001600160a01b03928316602085015291169082015260c085013560608201526080015b60405180910390a1505050565b6001600160a01b038481166000908152600b60205260408120549091161561122d576001600160a01b038086166000908152600b60205260409020541661122f565b845b905061123c848484611db2565b6001600160a01b0316816001600160a01b03161461129c5760405162461bcd60e51b815260206004820152601760248201527f5266713a206e6f7420616c6c6f776564207369676e65720000000000000000006044820152606401610ae3565b5050505050565b336112b66000546001600160a01b031690565b6001600160a01b0316146112fa5760405162461bcd60e51b815260206004820181905260248201526000805160206152fc8339815191526044820152606401610ae3565b6113038161368f565b50565b3360009081526002602052604090205460ff166113655760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f74207061757365720000000000000000000000006044820152606401610ae3565b61136d61374c565b565b336113826000546001600160a01b031690565b6001600160a01b0316146113c65760405162461bcd60e51b815260206004820181905260248201526000805160206152fc8339815191526044820152606401610ae3565b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e906020015b60405180910390a150565b60015460405163299aee5160e11b81526000916001600160a01b031690635335dca29061144e9086908690600401614d70565b602060405180830381865afa15801561146b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061148f9190614d9f565b9392505050565b6001600160a01b0381166114c457336000908152600b6020526040902080546001600160a01b031916905550565b336000908152600b6020526040902080546001600160a01b0383166001600160a01b031990911617905550565b6001546000906001600160a01b0316331461154e5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610ae3565b979650505050505050565b3361156c6000546001600160a01b031690565b6001600160a01b0316146115b05760405162461bcd60e51b815260206004820181905260248201526000805160206152fc8339815191526044820152606401610ae3565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b600154600160a01b900460ff161561161f5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b6004546001600160a01b031661163b6040850160208601614a36565b6001600160a01b0316146116915760405162461bcd60e51b815260206004820152601760248201527f5266713a2073726320746f6b656e206d69736d617463680000000000000000006044820152606401610ae3565b60008061169f8585856137f2565b6000828152600760205260409020805460ff1916600517905590925090506116dc6116d261016087016101408801614a36565b8660400135613986565b7f2e0668a62a5f556368dca9c7113e20f2852c05155548243804bf714ce72b25a6828261170f6040890160208a01614a36565b604080519384526001600160a01b03928316602085015291168282015287013560608201526080015b60405180910390a15050505050565b6008546001600160a01b031661179f5760405162461bcd60e51b815260206004820152601d60248201527f5266713a2074726561737572792061646472657373206e6f74207365740000006044820152606401610ae3565b6001600160a01b038082166000818152600a60205260408120805491905560085490926117ce92911683613aee565b600854604080516001600160a01b039283168152918416602083015281018290527ff228de527fc1b9843baac03b9a04565473a263375950e63435d4138464386f469060600160405180910390a15050565b336118336000546001600160a01b031690565b6001600160a01b0316146118775760405162461bcd60e51b815260206004820181905260248201526000805160206152fc8339815191526044820152606401610ae3565b61130381613b7e565b61136d33613b7e565b600154600160a01b900460ff16156118d65760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b426118e8610100830160e08401614a51565b6001600160401b0316106119485760405162461bcd60e51b815260206004820152602160248201527f5266713a207472616e7366657220646561646c696e65206e6f742070617373656044820152601960fa1b6064820152608401610ae3565b6001600160401b03461661196260a0830160808401614a51565b6001600160401b0316146119b85760405162461bcd60e51b815260206004820152601960248201527f5266713a2064737420636861696e4964206d69736d61746368000000000000006044820152606401610ae3565b60006005816119ca6020850185614a51565b6001600160401b031681526020810191909152604001600020546001600160a01b0316905080611a3c5760405162461bcd60e51b815260206004820152601d60248201527f5266713a207372632072667120636f6e7472616374206e6f74207365740000006044820152606401610ae3565b6000611a478361294e565b90506000808281526007602052604090205460ff166008811115611a6d57611a6d614855565b14611aba5760405162461bcd60e51b815260206004820152601b60248201527f5266713a2071756f746520616c726561647920657865637574656400000000006044820152606401610ae3565b6000818152600760209081526040808320805460ff1916600617905551611ae691849160029101614d51565b60408051601f198184030181528282528051602091820120908301520160408051601f198184030181529190529050611b2d83611b266020870187614a51565b833461340f565b6040518281527f7cdd4403cff3a09d96c1ffe4ad1cc5c195e9053463a55edfc2944644ec02211890602001610f91565b33611b706000546001600160a01b031690565b6001600160a01b031614611bb45760405162461bcd60e51b815260206004820181905260248201526000805160206152fc8339815191526044820152606401610ae3565b61130381613c37565b3360009081526002602052604090205460ff16611c1c5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f74207061757365720000000000000000000000006044820152606401610ae3565b61136d613cf4565b6001546000906001600160a01b03163314611c815760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610ae3565b60208314611cd15760405162461bcd60e51b815260206004820152601d60248201527f5266713a20696e636f7272656374206d657373616765206c656e6774680000006044820152606401610ae3565b6001600160401b0385166000908152600560205260409020546001600160a01b039081169087168114611d08576002915050610d54565b600160066000611d188789614db8565b81526020810191909152604001600020805460ff19169115159190911790555060019695505050505050565b600154600160a01b900460ff1615611d915760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b6000611d9e848484613d7c565b9050611dac8482600061353c565b50505050565b604080514660208201526bffffffffffffffffffffffff193060601b16918101919091527f416c6c6f7765645472616e7366657200000000000000000000000000000000006054820152606381018490526000908190611e2a9060830160405160208183030381529060405280519060200120613e09565b9050610d5484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508593925050613e449050565b600154600160a01b900460ff1615611ebb5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b600080611ec783613175565b6000828152600760208181526040808420805460ff1916909317909255905193955091935091611efc91859160019101614d51565b60408051601f198184030181528282528051602091820120908301520160408051601f198184030181529190529050611f3c82611b266020870187614a51565b610f2233611f5261016087016101408801614a36565b60c08701803590611f669060a08a01614a36565b6001600160a01b0316929190613e68565b33611f8a6000546001600160a01b031690565b6001600160a01b031614611fce5760405162461bcd60e51b815260206004820181905260248201526000805160206152fc8339815191526044820152606401610ae3565b600880546001600160a01b0319166001600160a01b0383169081179091556040519081527fb17659014001857e7557191ad74dc9e967b181eaed0895975325e3848debbc4290602001611410565b3360009081526003602052604090205460ff1661207b5760405162461bcd60e51b815260206004820152601660248201527f43616c6c6572206973206e6f7420676f7665726e6f72000000000000000000006044820152606401610ae3565b8281146120ca5760405162461bcd60e51b815260206004820152601460248201527f5266713a206c656e677468206d69736d617463680000000000000000000000006044820152606401610ae3565b60005b8381101561226c57620f42408383838181106120eb576120eb614dd6565b90506020020160208101906121009190614e00565b63ffffffff16106121535760405162461bcd60e51b815260206004820152601d60248201527f5266713a206665652070657263656e7461676520746f6f206c617267650000006044820152606401610ae3565b84848281811061216557612165614dd6565b905060200201602081019061217a9190614a51565b6001600160401b03166000036121d65782828281811061219c5761219c614dd6565b90506020020160208101906121b19190614e00565b600860146101000a81548163ffffffff021916908363ffffffff16021790555061225a565b8282828181106121e8576121e8614dd6565b90506020020160208101906121fd9190614e00565b6009600087878581811061221357612213614dd6565b90506020020160208101906122289190614a51565b6001600160401b031681526020810191909152604001600020805463ffffffff191663ffffffff929092169190911790555b8061226481614e1b565b9150506120cd565b507f541df5e570cf10ffe04899eebd9eebebd1c54e2bd4af9f24b23fb4a40c6ea00b84848484604051610f919493929190614e7b565b600154600160a01b900460ff16156122ef5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b6000806122fb85613175565b909250905061231d6123156101a087016101808801614a36565b8386866111eb565b6000828152600760208181526040808420805460ff1916909317909255905161234b91859160019101614d51565b60408051601f198184030181528282528051602091820120908301520160408051601f19818403018152919052905061238b82611b266020890189614a51565b6123c66123a06101a088016101808901614a36565b6123b261016089016101408a01614a36565b60c08901803590611f669060a08c01614a36565b7fb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc5836123fa61016089016101408a01614a36565b61240a60c08a0160a08b01614a36565b604080519384526001600160a01b03928316602085015291169082015260c0880135606082015260800160405180910390a1505050505050565b600154600160a01b900460ff16156124915760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b6004546001600160a01b03166124ad6040850160208601614a36565b6001600160a01b0316146125035760405162461bcd60e51b815260206004820152601760248201527f5266713a2073726320746f6b656e206d69736d617463680000000000000000006044820152606401610ae3565b6000612510848484613d7c565b9050611dac8482600161353c565b600154600160a01b900460ff161561256b5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b61257b60a0850160808601614a51565b6001600160401b03166125916020860186614a51565b6001600160401b0316146125e75760405162461bcd60e51b815260206004820152601860248201527f5266713a206e6f742073616d6520636861696e207377617000000000000000006044820152606401610ae3565b60006125f285613175565b50905061261261260a6101a087016101808801614a36565b8285856111eb565b61264d6126276101a087016101808801614a36565b61263961016088016101408901614a36565b60c08801803590611f669060a08b01614a36565b61265885828661353c565b7fb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc58161268c61016088016101408901614a36565b61269c60c0890160a08a01614a36565b604080519384526001600160a01b03928316602085015291169082015260c08701356060820152608001611738565b600154600160a01b900460ff16156127185760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b61272860a0830160808401614a51565b6001600160401b031661273e6020840184614a51565b6001600160401b0316146127945760405162461bcd60e51b815260206004820152601860248201527f5266713a206e6f742073616d6520636861696e207377617000000000000000006044820152606401610ae3565b600061279f83613175565b509050611164336127b861016086016101408701614a36565b60c08601803590611f669060a08901614a36565b336127df6000546001600160a01b031690565b6001600160a01b0316146128235760405162461bcd60e51b815260206004820181905260248201526000805160206152fc8339815191526044820152606401610ae3565b8281146128725760405162461bcd60e51b815260206004820152601460248201527f5266713a206c656e677468206d69736d617463680000000000000000000000006044820152606401610ae3565b60005b838110156129185782828281811061288f5761288f614dd6565b90506020020160208101906128a49190614a36565b600560008787858181106128ba576128ba614dd6565b90506020020160208101906128cf9190614a51565b6001600160401b03168152602081019190915260400160002080546001600160a01b0319166001600160a01b03929092169190911790558061291081614e1b565b915050612875565b507fb4739c640c5970d8ce88b6c31f3706099efca660e282d47b0a267a0bb572d8b784848484604051610f919493929190614edb565b600061295d6020830183614a51565b61296d6040840160208501614a36565b6040840135606085013561298760a0870160808801614a51565b61299760c0880160a08901614a36565b60c08801356129ad6101008a0160e08b01614a51565b6129bf6101208b016101008c01614a51565b6129d16101408c016101208d01614a36565b6129e36101608d016101408e01614a36565b6129f56101808e016101608f01614a36565b8d610180016020810190612a099190614a36565b604051602001612a259d9c9b9a99989796959493929190614f31565b604051602081830303815290604052805190602001209050919050565b61136d33613ea0565b600154600090600160a01b900460ff1615612a9b5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b6000612aa8848434612d05565b905061148f333060408701803590611f669060208a01614a36565b33612ad66000546001600160a01b031690565b6001600160a01b031614612b1a5760405162461bcd60e51b815260206004820181905260248201526000805160206152fc8339815191526044820152606401610ae3565b61130381613ea0565b33612b366000546001600160a01b031690565b6001600160a01b031614612b7a5760405162461bcd60e51b815260206004820181905260248201526000805160206152fc8339815191526044820152606401610ae3565b6001600160a01b038116612bf65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610ae3565b61130381613f59565b600154600160a01b900460ff1615612c4c5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b600080612c5a8585856137f2565b6000828152600760205260409020805492945090925060049160ff191660018302179055506116dc8160408701803590612c979060208a01614a36565b6001600160a01b03169190613aee565b6001600160401b03821660009081526009602052604081205463ffffffff16808203612cdf5750600854600160a01b900463ffffffff165b620f4240612cf363ffffffff831685615060565b612cfd9190615077565b949350505050565b600042836001600160401b0316118015612d4057506001600160401b038316612d35610100860160e08701614a51565b6001600160401b0316115b612d8c5760405162461bcd60e51b815260206004820152601b60248201527f5266713a20696e617070726f70726961746520646561646c696e6500000000006044820152606401610ae3565b6000612da061016086016101408701614a36565b6001600160a01b031614158015612dd257506000612dc66101a086016101808701614a36565b6001600160a01b031614155b612e445760405162461bcd60e51b815260206004820152602a60248201527f5266713a20696e76616c6964207265636569766572206f72206c69717569646960448201527f747950726f7669646572000000000000000000000000000000000000000000006064820152608401610ae3565b6001600160401b034616612e5b6020860186614a51565b6001600160401b031614612eb15760405162461bcd60e51b815260206004820152601960248201527f5266713a2073726320636861696e4964206d69736d61746368000000000000006044820152606401610ae3565b33612ec461014086016101208701614a36565b6001600160a01b031614612f1a5760405162461bcd60e51b815260206004820152601460248201527f5266713a2073656e646572206d69736d617463680000000000000000000000006044820152606401610ae3565b6000612f258561294e565b90506000808281526007602052604090205460ff166008811115612f4b57612f4b614855565b14612f985760405162461bcd60e51b815260206004820152601660248201527f5266713a2071756f7465206861736820657869737473000000000000000000006044820152606401610ae3565b6000612fb7612fad60a0880160808901614a51565b8760400135612ca7565b9050612fcb60608701356040880135614d3e565b81111561301a5760405162461bcd60e51b815260206004820152601e60248201527f5266713a20696e73756666696369656e742070726f746f636f6c2066656500006044820152606401610ae3565b6000828152600760205260409020805460ff1916600117905561304360a0870160808801614a51565b6001600160401b03166130596020880188614a51565b6001600160401b03161461313357600060058161307c60a08a0160808b01614a51565b6001600160401b031681526020810191909152604001600020546001600160a01b03169050806130ee5760405162461bcd60e51b815260206004820152601960248201527f5266713a2064737420636f6e7472616374206e6f7420736574000000000000006044820152606401610ae3565b60008360405160200161310391815260200190565b60408051601f1981840301815291905290506131308261312960a08b0160808c01614a51565b838961340f565b50505b7f3e4de2d1674631d426ae2a89635b421e6d40a31d27681afdf0eed67e81d07bcb8287604051613164929190615099565b60405180910390a150949350505050565b6000804261318a610100850160e08601614a51565b6001600160401b0316116131e05760405162461bcd60e51b815260206004820152601d60248201527f5266713a207472616e7366657220646561646c696e65207061737365640000006044820152606401610ae3565b6001600160401b0346166131fa60a0850160808601614a51565b6001600160401b0316146132505760405162461bcd60e51b815260206004820152601960248201527f5266713a2064737420636861696e4964206d69736d61746368000000000000006044820152606401610ae3565b600061325b8461294e565b9050600060058161326f6020880188614a51565b6001600160401b031681526020810191909152604001600020546001600160a01b031690506132a460a0860160808701614a51565b6001600160401b03166132ba6020870187614a51565b6001600160401b0316146133935760008281526007602052604081205460ff1660088111156132eb576132eb614855565b146133385760405162461bcd60e51b815260206004820152601b60248201527f5266713a2071756f746520616c726561647920657865637574656400000000006044820152606401610ae3565b6001600160a01b03811661338e5760405162461bcd60e51b815260206004820152601d60248201527f5266713a206473742072667120636f6e7472616374206e6f74207365740000006044820152606401610ae3565b613405565b600160008381526007602052604090205460ff1660088111156133b8576133b8614855565b146134055760405162461bcd60e51b815260206004820152601d60248201527f5266713a206e6f206465706f736974206f6e2073616d6520636861696e0000006044820152606401610ae3565b9094909350915050565b600154611dac908590859085906001600160a01b031685613fa9565b6004546001600160a01b03166134835760405162461bcd60e51b815260206004820152601860248201527f5266713a206e61746976652077726170206e6f742073657400000000000000006044820152606401610ae3565b6000826001600160a01b03168261c35090604051600060405180830381858888f193505050503d80600081146134d5576040519150601f19603f3d011682016040523d82523d6000602084013e6134da565b606091505b50509050806135375760405162461bcd60e51b8152602060048201526024808201527f5266713a206661696c656420746f207472616e73666572206e6174697665207460448201526337b5b2b760e11b6064820152608401610ae3565b505050565b61354e60608401356040850135614d3e565b600a60006135626040870160208801614a36565b6001600160a01b03166001600160a01b031681526020019081526020016000206000828254613591919061520b565b909155505080156135d9576000828152600760205260409020805460ff191660031790556135d46135ca6101a085016101808601614a36565b8460600135613986565b61361c565b6000828152600760205260409020805460ff1916600217905561361c6136076101a085016101808601614a36565b6060850135612c976040870160208801614a36565b7ff29b32a17c591b8b3b1216ce0ffb67c07f3478e99b50c5ccf8602878b1ee6376826136506101a086016101808701614a36565b6136606040870160208801614a36565b604080519384526001600160a01b039283166020850152911690820152606080860135908201526080016111de565b6001600160a01b03811660009081526003602052604090205460ff16156136f85760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f7200000000006044820152606401610ae3565b6001600160a01b038116600081815260036020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b59101611410565b600154600160a01b900460ff166137a55760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610ae3565b6001805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60008060006138008661294e565b9050600160008281526007602052604090205460ff16600881111561382757613827614855565b146138745760405162461bcd60e51b815260206004820152601960248201527f5266713a20696e636f72726563742071756f74652068617368000000000000006044820152606401610ae3565b61388460a0870160808801614a51565b6001600160401b031661389a6020880188614a51565b6001600160401b0316146138ba576138b58585836002614014565b61392c565b426138cc610100880160e08901614a51565b6001600160401b03161061392c5760405162461bcd60e51b815260206004820152602160248201527f5266713a207472616e7366657220646561646c696e65206e6f742070617373656044820152601960fa1b6064820152608401610ae3565b60008061394161018089016101608a01614a36565b6001600160a01b0316146139665761396161018088016101608901614a36565b613978565b61397861014088016101208901614a36565b919791965090945050505050565b6004546001600160a01b03166139de5760405162461bcd60e51b815260206004820152601860248201527f5266713a206e61746976652077726170206e6f742073657400000000000000006044820152606401610ae3565b60048054604051632e1a7d4d60e01b81529182018390526001600160a01b031690632e1a7d4d90602401600060405180830381600087803b158015613a2257600080fd5b505af1158015613a36573d6000803e3d6000fd5b505050506000826001600160a01b03168261c35090604051600060405180830381858888f193505050503d8060008114613a8c576040519150601f19603f3d011682016040523d82523d6000602084013e613a91565b606091505b50509050806135375760405162461bcd60e51b8152602060048201526024808201527f5266713a206661696c656420746f207769746864726177206e6174697665207460448201526337b5b2b760e11b6064820152608401610ae3565b6040516001600160a01b03831660248201526044810182905261353790849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152614187565b6001600160a01b03811660009081526002602052604090205460ff16613be65760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f742070617573657200000000000000000000006044820152606401610ae3565b6001600160a01b038116600081815260026020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e9101611410565b6001600160a01b03811660009081526002602052604090205460ff1615613ca05760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c726561647920706175736572000000000000006044820152606401610ae3565b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f89101611410565b600154600160a01b900460ff1615613d415760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610ae3565b6001805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586137d53390565b600080613d888561294e565b9050600160008281526007602052604090205460ff166008811115613daf57613daf614855565b14613dfc5760405162461bcd60e51b815260206004820152601960248201527f5266713a20696e636f72726563742071756f74652068617368000000000000006044820152606401610ae3565b612cfd8484836001614014565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01612a25565b6000806000613e53858561426c565b91509150613e60816142da565b509392505050565b6040516001600160a01b0380851660248301528316604482015260648101829052611dac9085906323b872dd60e01b90608401613b1a565b6001600160a01b03811660009081526003602052604090205460ff16613f085760405162461bcd60e51b815260206004820152601760248201527f4163636f756e74206973206e6f7420676f7665726e6f720000000000000000006044820152606401610ae3565b6001600160a01b038116600081815260036020908152604091829020805460ff1916905590519182527f1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b9101611410565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b604051634f9e72ad60e11b81526001600160a01b03831690639f3ce55a908390613fdb9089908990899060040161526e565b6000604051808303818588803b158015613ff457600080fd5b505af1158015614008573d6000803e3d6000fd5b50505050505050505050565b60008282604051602001614029929190614d51565b60408051601f1981840301815291815281516020928301206000818152600690935291205490915060ff1661410d576001546040516000916001600160a01b031690614078908890889061529f565b6000604051808303816000865af19150503d80600081146140b5576040519150601f19603f3d011682016040523d82523d6000602084013e6140ba565b606091505b505090508061410b5760405162461bcd60e51b815260206004820152601260248201527f65786563757465206d7367206661696c656400000000000000000000000000006044820152606401610ae3565b505b60008181526006602052604090205460ff1661416b5760405162461bcd60e51b815260206004820152601060248201527f5266713a20696e76616c6964206d7367000000000000000000000000000000006044820152606401610ae3565b6000908152600660205260409020805460ff1916905550505050565b60006141dc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166144909092919063ffffffff16565b80519091501561353757808060200190518101906141fa91906152af565b6135375760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610ae3565b60008082516041036142a25760208301516040840151606085015160001a6142968782858561449f565b945094505050506142d3565b82516040036142cb57602083015160408401516142c086838361458c565b9350935050506142d3565b506000905060025b9250929050565b60008160048111156142ee576142ee614855565b036142f65750565b600181600481111561430a5761430a614855565b036143575760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610ae3565b600281600481111561436b5761436b614855565b036143b85760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610ae3565b60038160048111156143cc576143cc614855565b036144245760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610ae3565b600481600481111561443857614438614855565b036113035760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610ae3565b6060612cfd84846000856145de565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156144d65750600090506003614583565b8460ff16601b141580156144ee57508460ff16601c14155b156144ff5750600090506004614583565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614553573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661457c57600060019250925050614583565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316816145c260ff86901c601b61520b565b90506145d08782888561449f565b935093505050935093915050565b6060824710156146565760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610ae3565b6001600160a01b0385163b6146ad5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610ae3565b600080866001600160a01b031685876040516146c991906152cc565b60006040518083038185875af1925050503d8060008114614706576040519150601f19603f3d011682016040523d82523d6000602084013e61470b565b606091505b509150915061154e8282866060831561472557508161148f565b8251156147355782518084602001fd5b8160405162461bcd60e51b8152600401610ae391906152e8565b60008083601f84011261476157600080fd5b5081356001600160401b0381111561477857600080fd5b6020830191508360208285010111156142d357600080fd5b80356001600160401b03811681146147a757600080fd5b919050565b80356001600160a01b03811681146147a757600080fd5b600080600080600080608087890312156147dc57600080fd5b86356001600160401b03808211156147f357600080fd5b6147ff8a838b0161474f565b909850965086915061481360208a01614790565b9550604089013591508082111561482957600080fd5b5061483689828a0161474f565b90945092506148499050606088016147ac565b90509295509295509295565b634e487b7160e01b600052602160045260246000fd5b6003811061130357611303614855565b602081016148888361486b565b91905290565b60006101a082840312156148a157600080fd5b50919050565b6000806101c083850312156148bb57600080fd5b6148c5848461488e565b91506148d46101a08401614790565b90509250929050565b6000806000806000608086880312156148f557600080fd5b6148fe866147ac565b94506020860135935060408601356001600160401b0381111561492057600080fd5b61492c8882890161474f565b909450925061493f9050606087016147ac565b90509295509295909350565b60006101a0828403121561495e57600080fd5b61148f838361488e565b801515811461130357600080fd5b6000806101c0838503121561498a57600080fd5b614994848461488e565b91506101a08301356149a581614968565b809150509250929050565b6000602082840312156149c257600080fd5b5035919050565b602081016009831061488857614888614855565b600080600080606085870312156149f357600080fd5b6149fc856147ac565b93506020850135925060408501356001600160401b03811115614a1e57600080fd5b614a2a8782880161474f565b95989497509550505050565b600060208284031215614a4857600080fd5b61148f826147ac565b600060208284031215614a6357600080fd5b61148f82614790565b60008060208385031215614a7f57600080fd5b82356001600160401b03811115614a9557600080fd5b614aa18582860161474f565b90969095509350505050565b600080600080600080600060c0888a031215614ac857600080fd5b614ad1886147ac565b9650614adf602089016147ac565b955060408801359450614af460608901614790565b935060808801356001600160401b03811115614b0f57600080fd5b614b1b8a828b0161474f565b9094509250614b2e905060a089016147ac565b905092959891949750929550565b60008060006101c08486031215614b5257600080fd5b614b5c858561488e565b92506101a08401356001600160401b03811115614b7857600080fd5b614b848682870161474f565b9497909650939450505050565b600080600080600060808688031215614ba957600080fd5b614bb2866147ac565b9450614bc060208701614790565b935060408601356001600160401b0381111561492057600080fd5b600080600060408486031215614bf057600080fd5b8335925060208401356001600160401b03811115614b7857600080fd5b60008083601f840112614c1f57600080fd5b5081356001600160401b03811115614c3657600080fd5b6020830191508360208260051b85010111156142d357600080fd5b60008060008060408587031215614c6757600080fd5b84356001600160401b0380821115614c7e57600080fd5b614c8a88838901614c0d565b90965094506020870135915080821115614ca357600080fd5b50614a2a87828801614c0d565b6000806000806101e08587031215614cc757600080fd5b614cd1868661488e565b93506101a0850135614ce281614968565b92506101c08501356001600160401b03811115614a1e57600080fd5b60008060408385031215614d1157600080fd5b614d1a83614790565b946020939093013593505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610cf157610cf1614d28565b828152614d5d8261486b565b60f89190911b6020820152602101919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600060208284031215614db157600080fd5b5051919050565b80356020831015610cf157600019602084900360031b1b1692915050565b634e487b7160e01b600052603260045260246000fd5b803563ffffffff811681146147a757600080fd5b600060208284031215614e1257600080fd5b61148f82614dec565b600060018201614e2d57614e2d614d28565b5060010190565b8183526000602080850194508260005b85811015614e70576001600160401b03614e5d83614790565b1687529582019590820190600101614e44565b509495945050505050565b604081526000614e8f604083018688614e34565b8281036020848101919091528482528591810160005b86811015614ece5763ffffffff614ebb85614dec565b1682529282019290820190600101614ea5565b5098975050505050505050565b604081526000614eef604083018688614e34565b8281036020848101919091528482528591810160005b86811015614ece576001600160a01b03614f1e856147ac565b1682529282019290820190600101614f05565b6001600160c01b03198e60c01b1681526bffffffffffffffffffffffff198d60601b1660088201528b601c8201528a603c820152614f7e605c82018b60c01b6001600160c01b0319169052565b614f9c606482018a60601b6bffffffffffffffffffffffff19169052565b876078820152614fbb609882018860c01b6001600160c01b0319169052565b614fd460a082018760c01b6001600160c01b0319169052565b614ff260a882018660601b6bffffffffffffffffffffffff19169052565b61501060bc82018560601b6bffffffffffffffffffffffff19169052565b61502e60d082018460601b6bffffffffffffffffffffffff19169052565b61504c60e482018360601b6bffffffffffffffffffffffff19169052565b60f8019d9c50505050505050505050505050565b8082028115828204841417610cf157610cf1614d28565b60008261509457634e487b7160e01b600052601260045260246000fd5b500490565b8281526101c081016150be602083016150b185614790565b6001600160401b03169052565b6150ca602084016147ac565b6001600160a01b03811660408401525060408301356060830152606083013560808301526150fa60808401614790565b6001600160401b03811660a08401525061511660a084016147ac565b6001600160a01b03811660c08401525060c083013560e083015261513c60e08401614790565b610100615153818501836001600160401b03169052565b61515e818601614790565b915050610120615178818501836001600160401b03169052565b6151838186016147ac565b91505061014061519d818501836001600160a01b03169052565b6151a88186016147ac565b9150506101606151c2818501836001600160a01b03169052565b6151cd8186016147ac565b9150506101806151e7818501836001600160a01b03169052565b6151f28186016147ac565b915050613e606101a08401826001600160a01b03169052565b80820180821115610cf157610cf1614d28565b60005b83811015615239578181015183820152602001615221565b50506000910152565b6000815180845261525a81602086016020860161521e565b601f01601f19169290920160200192915050565b6001600160a01b03841681526001600160401b0383166020820152606060408201526000610d546060830184615242565b8183823760009101908152919050565b6000602082840312156152c157600080fd5b815161148f81614968565b600082516152de81846020870161521e565b9190910192915050565b60208152600061148f602083018461524256fe4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a26469706673582212202bcb7d6c1e2f7962c5b8b8d31f87ca0d09869be33786852a6840ef1de3d68e8c64736f6c63430008110033",
}

// RfqABI is the input ABI used to generate the binding from.
// Deprecated: Use RfqMetaData.ABI instead.
var RfqABI = RfqMetaData.ABI

// RfqBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RfqMetaData.Bin instead.
var RfqBin = RfqMetaData.Bin

// DeployRfq deploys a new Ethereum contract, binding an instance of Rfq to it.
func DeployRfq(auth *bind.TransactOpts, backend bind.ContractBackend, _messageBus common.Address) (common.Address, *types.Transaction, *Rfq, error) {
	parsed, err := RfqMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RfqBin), backend, _messageBus)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rfq{RfqCaller: RfqCaller{contract: contract}, RfqTransactor: RfqTransactor{contract: contract}, RfqFilterer: RfqFilterer{contract: contract}}, nil
}

// Rfq is an auto generated Go binding around an Ethereum contract.
type Rfq struct {
	RfqCaller     // Read-only binding to the contract
	RfqTransactor // Write-only binding to the contract
	RfqFilterer   // Log filterer for contract events
}

// RfqCaller is an auto generated read-only Go binding around an Ethereum contract.
type RfqCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RfqTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RfqTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RfqFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RfqFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RfqSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RfqSession struct {
	Contract     *Rfq              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RfqCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RfqCallerSession struct {
	Contract *RfqCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RfqTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RfqTransactorSession struct {
	Contract     *RfqTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RfqRaw is an auto generated low-level Go binding around an Ethereum contract.
type RfqRaw struct {
	Contract *Rfq // Generic contract binding to access the raw methods on
}

// RfqCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RfqCallerRaw struct {
	Contract *RfqCaller // Generic read-only contract binding to access the raw methods on
}

// RfqTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RfqTransactorRaw struct {
	Contract *RfqTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRfq creates a new instance of Rfq, bound to a specific deployed contract.
func NewRfq(address common.Address, backend bind.ContractBackend) (*Rfq, error) {
	contract, err := bindRfq(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rfq{RfqCaller: RfqCaller{contract: contract}, RfqTransactor: RfqTransactor{contract: contract}, RfqFilterer: RfqFilterer{contract: contract}}, nil
}

// NewRfqCaller creates a new read-only instance of Rfq, bound to a specific deployed contract.
func NewRfqCaller(address common.Address, caller bind.ContractCaller) (*RfqCaller, error) {
	contract, err := bindRfq(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RfqCaller{contract: contract}, nil
}

// NewRfqTransactor creates a new write-only instance of Rfq, bound to a specific deployed contract.
func NewRfqTransactor(address common.Address, transactor bind.ContractTransactor) (*RfqTransactor, error) {
	contract, err := bindRfq(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RfqTransactor{contract: contract}, nil
}

// NewRfqFilterer creates a new log filterer instance of Rfq, bound to a specific deployed contract.
func NewRfqFilterer(address common.Address, filterer bind.ContractFilterer) (*RfqFilterer, error) {
	contract, err := bindRfq(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RfqFilterer{contract: contract}, nil
}

// bindRfq binds a generic wrapper to an already deployed contract.
func bindRfq(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RfqABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rfq *RfqRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rfq.Contract.RfqCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rfq *RfqRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.Contract.RfqTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rfq *RfqRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rfq.Contract.RfqTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rfq *RfqCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rfq.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rfq *RfqTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rfq *RfqTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rfq.Contract.contract.Transact(opts, method, params...)
}

// AllowedSigner is a free data retrieval call binding the contract method 0xbbf3b98f.
//
// Solidity: function allowedSigner(address ) view returns(address)
func (_Rfq *RfqCaller) AllowedSigner(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "allowedSigner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedSigner is a free data retrieval call binding the contract method 0xbbf3b98f.
//
// Solidity: function allowedSigner(address ) view returns(address)
func (_Rfq *RfqSession) AllowedSigner(arg0 common.Address) (common.Address, error) {
	return _Rfq.Contract.AllowedSigner(&_Rfq.CallOpts, arg0)
}

// AllowedSigner is a free data retrieval call binding the contract method 0xbbf3b98f.
//
// Solidity: function allowedSigner(address ) view returns(address)
func (_Rfq *RfqCallerSession) AllowedSigner(arg0 common.Address) (common.Address, error) {
	return _Rfq.Contract.AllowedSigner(&_Rfq.CallOpts, arg0)
}

// FeePercGlobal is a free data retrieval call binding the contract method 0x0bd930b4.
//
// Solidity: function feePercGlobal() view returns(uint32)
func (_Rfq *RfqCaller) FeePercGlobal(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "feePercGlobal")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeePercGlobal is a free data retrieval call binding the contract method 0x0bd930b4.
//
// Solidity: function feePercGlobal() view returns(uint32)
func (_Rfq *RfqSession) FeePercGlobal() (uint32, error) {
	return _Rfq.Contract.FeePercGlobal(&_Rfq.CallOpts)
}

// FeePercGlobal is a free data retrieval call binding the contract method 0x0bd930b4.
//
// Solidity: function feePercGlobal() view returns(uint32)
func (_Rfq *RfqCallerSession) FeePercGlobal() (uint32, error) {
	return _Rfq.Contract.FeePercGlobal(&_Rfq.CallOpts)
}

// FeePercOverride is a free data retrieval call binding the contract method 0x3e07d172.
//
// Solidity: function feePercOverride(uint64 ) view returns(uint32)
func (_Rfq *RfqCaller) FeePercOverride(opts *bind.CallOpts, arg0 uint64) (uint32, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "feePercOverride", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeePercOverride is a free data retrieval call binding the contract method 0x3e07d172.
//
// Solidity: function feePercOverride(uint64 ) view returns(uint32)
func (_Rfq *RfqSession) FeePercOverride(arg0 uint64) (uint32, error) {
	return _Rfq.Contract.FeePercOverride(&_Rfq.CallOpts, arg0)
}

// FeePercOverride is a free data retrieval call binding the contract method 0x3e07d172.
//
// Solidity: function feePercOverride(uint64 ) view returns(uint32)
func (_Rfq *RfqCallerSession) FeePercOverride(arg0 uint64) (uint32, error) {
	return _Rfq.Contract.FeePercOverride(&_Rfq.CallOpts, arg0)
}

// GetMsgFee is a free data retrieval call binding the contract method 0x55e9e3d2.
//
// Solidity: function getMsgFee(bytes _message) view returns(uint256)
func (_Rfq *RfqCaller) GetMsgFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "getMsgFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMsgFee is a free data retrieval call binding the contract method 0x55e9e3d2.
//
// Solidity: function getMsgFee(bytes _message) view returns(uint256)
func (_Rfq *RfqSession) GetMsgFee(_message []byte) (*big.Int, error) {
	return _Rfq.Contract.GetMsgFee(&_Rfq.CallOpts, _message)
}

// GetMsgFee is a free data retrieval call binding the contract method 0x55e9e3d2.
//
// Solidity: function getMsgFee(bytes _message) view returns(uint256)
func (_Rfq *RfqCallerSession) GetMsgFee(_message []byte) (*big.Int, error) {
	return _Rfq.Contract.GetMsgFee(&_Rfq.CallOpts, _message)
}

// GetQuoteHash is a free data retrieval call binding the contract method 0xd1ce8b45.
//
// Solidity: function getQuoteHash((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) pure returns(bytes32)
func (_Rfq *RfqCaller) GetQuoteHash(opts *bind.CallOpts, _quote RFQQuote) ([32]byte, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "getQuoteHash", _quote)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetQuoteHash is a free data retrieval call binding the contract method 0xd1ce8b45.
//
// Solidity: function getQuoteHash((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) pure returns(bytes32)
func (_Rfq *RfqSession) GetQuoteHash(_quote RFQQuote) ([32]byte, error) {
	return _Rfq.Contract.GetQuoteHash(&_Rfq.CallOpts, _quote)
}

// GetQuoteHash is a free data retrieval call binding the contract method 0xd1ce8b45.
//
// Solidity: function getQuoteHash((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) pure returns(bytes32)
func (_Rfq *RfqCallerSession) GetQuoteHash(_quote RFQQuote) ([32]byte, error) {
	return _Rfq.Contract.GetQuoteHash(&_Rfq.CallOpts, _quote)
}

// GetRfqFee is a free data retrieval call binding the contract method 0xfbe42fea.
//
// Solidity: function getRfqFee(uint64 _chainId, uint256 _amount) view returns(uint256)
func (_Rfq *RfqCaller) GetRfqFee(opts *bind.CallOpts, _chainId uint64, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "getRfqFee", _chainId, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRfqFee is a free data retrieval call binding the contract method 0xfbe42fea.
//
// Solidity: function getRfqFee(uint64 _chainId, uint256 _amount) view returns(uint256)
func (_Rfq *RfqSession) GetRfqFee(_chainId uint64, _amount *big.Int) (*big.Int, error) {
	return _Rfq.Contract.GetRfqFee(&_Rfq.CallOpts, _chainId, _amount)
}

// GetRfqFee is a free data retrieval call binding the contract method 0xfbe42fea.
//
// Solidity: function getRfqFee(uint64 _chainId, uint256 _amount) view returns(uint256)
func (_Rfq *RfqCallerSession) GetRfqFee(_chainId uint64, _amount *big.Int) (*big.Int, error) {
	return _Rfq.Contract.GetRfqFee(&_Rfq.CallOpts, _chainId, _amount)
}

// GetSignerOfQuoteHash is a free data retrieval call binding the contract method 0xa2f6531e.
//
// Solidity: function getSignerOfQuoteHash(bytes32 _quoteHash, bytes _sig) view returns(address)
func (_Rfq *RfqCaller) GetSignerOfQuoteHash(opts *bind.CallOpts, _quoteHash [32]byte, _sig []byte) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "getSignerOfQuoteHash", _quoteHash, _sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSignerOfQuoteHash is a free data retrieval call binding the contract method 0xa2f6531e.
//
// Solidity: function getSignerOfQuoteHash(bytes32 _quoteHash, bytes _sig) view returns(address)
func (_Rfq *RfqSession) GetSignerOfQuoteHash(_quoteHash [32]byte, _sig []byte) (common.Address, error) {
	return _Rfq.Contract.GetSignerOfQuoteHash(&_Rfq.CallOpts, _quoteHash, _sig)
}

// GetSignerOfQuoteHash is a free data retrieval call binding the contract method 0xa2f6531e.
//
// Solidity: function getSignerOfQuoteHash(bytes32 _quoteHash, bytes _sig) view returns(address)
func (_Rfq *RfqCallerSession) GetSignerOfQuoteHash(_quoteHash [32]byte, _sig []byte) (common.Address, error) {
	return _Rfq.Contract.GetSignerOfQuoteHash(&_Rfq.CallOpts, _quoteHash, _sig)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Rfq *RfqCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Rfq *RfqSession) Governors(arg0 common.Address) (bool, error) {
	return _Rfq.Contract.Governors(&_Rfq.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Rfq *RfqCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _Rfq.Contract.Governors(&_Rfq.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Rfq *RfqCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Rfq *RfqSession) IsGovernor(_account common.Address) (bool, error) {
	return _Rfq.Contract.IsGovernor(&_Rfq.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Rfq *RfqCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _Rfq.Contract.IsGovernor(&_Rfq.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Rfq *RfqCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Rfq *RfqSession) IsPauser(account common.Address) (bool, error) {
	return _Rfq.Contract.IsPauser(&_Rfq.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Rfq *RfqCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Rfq.Contract.IsPauser(&_Rfq.CallOpts, account)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_Rfq *RfqCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_Rfq *RfqSession) MessageBus() (common.Address, error) {
	return _Rfq.Contract.MessageBus(&_Rfq.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_Rfq *RfqCallerSession) MessageBus() (common.Address, error) {
	return _Rfq.Contract.MessageBus(&_Rfq.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Rfq *RfqCaller) NativeWrap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "nativeWrap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Rfq *RfqSession) NativeWrap() (common.Address, error) {
	return _Rfq.Contract.NativeWrap(&_Rfq.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Rfq *RfqCallerSession) NativeWrap() (common.Address, error) {
	return _Rfq.Contract.NativeWrap(&_Rfq.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rfq *RfqCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rfq *RfqSession) Owner() (common.Address, error) {
	return _Rfq.Contract.Owner(&_Rfq.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rfq *RfqCallerSession) Owner() (common.Address, error) {
	return _Rfq.Contract.Owner(&_Rfq.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rfq *RfqCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rfq *RfqSession) Paused() (bool, error) {
	return _Rfq.Contract.Paused(&_Rfq.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rfq *RfqCallerSession) Paused() (bool, error) {
	return _Rfq.Contract.Paused(&_Rfq.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Rfq *RfqCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Rfq *RfqSession) Pausers(arg0 common.Address) (bool, error) {
	return _Rfq.Contract.Pausers(&_Rfq.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Rfq *RfqCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _Rfq.Contract.Pausers(&_Rfq.CallOpts, arg0)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb62b31e4.
//
// Solidity: function protocolFee(address ) view returns(uint256)
func (_Rfq *RfqCaller) ProtocolFee(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "protocolFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb62b31e4.
//
// Solidity: function protocolFee(address ) view returns(uint256)
func (_Rfq *RfqSession) ProtocolFee(arg0 common.Address) (*big.Int, error) {
	return _Rfq.Contract.ProtocolFee(&_Rfq.CallOpts, arg0)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb62b31e4.
//
// Solidity: function protocolFee(address ) view returns(uint256)
func (_Rfq *RfqCallerSession) ProtocolFee(arg0 common.Address) (*big.Int, error) {
	return _Rfq.Contract.ProtocolFee(&_Rfq.CallOpts, arg0)
}

// Quotes is a free data retrieval call binding the contract method 0x25329eaf.
//
// Solidity: function quotes(bytes32 ) view returns(uint8)
func (_Rfq *RfqCaller) Quotes(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "quotes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Quotes is a free data retrieval call binding the contract method 0x25329eaf.
//
// Solidity: function quotes(bytes32 ) view returns(uint8)
func (_Rfq *RfqSession) Quotes(arg0 [32]byte) (uint8, error) {
	return _Rfq.Contract.Quotes(&_Rfq.CallOpts, arg0)
}

// Quotes is a free data retrieval call binding the contract method 0x25329eaf.
//
// Solidity: function quotes(bytes32 ) view returns(uint8)
func (_Rfq *RfqCallerSession) Quotes(arg0 [32]byte) (uint8, error) {
	return _Rfq.Contract.Quotes(&_Rfq.CallOpts, arg0)
}

// RemoteRfqContracts is a free data retrieval call binding the contract method 0xcc47e400.
//
// Solidity: function remoteRfqContracts(uint64 ) view returns(address)
func (_Rfq *RfqCaller) RemoteRfqContracts(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "remoteRfqContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteRfqContracts is a free data retrieval call binding the contract method 0xcc47e400.
//
// Solidity: function remoteRfqContracts(uint64 ) view returns(address)
func (_Rfq *RfqSession) RemoteRfqContracts(arg0 uint64) (common.Address, error) {
	return _Rfq.Contract.RemoteRfqContracts(&_Rfq.CallOpts, arg0)
}

// RemoteRfqContracts is a free data retrieval call binding the contract method 0xcc47e400.
//
// Solidity: function remoteRfqContracts(uint64 ) view returns(address)
func (_Rfq *RfqCallerSession) RemoteRfqContracts(arg0 uint64) (common.Address, error) {
	return _Rfq.Contract.RemoteRfqContracts(&_Rfq.CallOpts, arg0)
}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_Rfq *RfqCaller) TreasuryAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "treasuryAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_Rfq *RfqSession) TreasuryAddr() (common.Address, error) {
	return _Rfq.Contract.TreasuryAddr(&_Rfq.CallOpts)
}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_Rfq *RfqCallerSession) TreasuryAddr() (common.Address, error) {
	return _Rfq.Contract.TreasuryAddr(&_Rfq.CallOpts)
}

// UnconsumedMsg is a free data retrieval call binding the contract method 0xdf1f64ef.
//
// Solidity: function unconsumedMsg(bytes32 ) view returns(bool)
func (_Rfq *RfqCaller) UnconsumedMsg(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "unconsumedMsg", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnconsumedMsg is a free data retrieval call binding the contract method 0xdf1f64ef.
//
// Solidity: function unconsumedMsg(bytes32 ) view returns(bool)
func (_Rfq *RfqSession) UnconsumedMsg(arg0 [32]byte) (bool, error) {
	return _Rfq.Contract.UnconsumedMsg(&_Rfq.CallOpts, arg0)
}

// UnconsumedMsg is a free data retrieval call binding the contract method 0xdf1f64ef.
//
// Solidity: function unconsumedMsg(bytes32 ) view returns(bool)
func (_Rfq *RfqCallerSession) UnconsumedMsg(arg0 [32]byte) (bool, error) {
	return _Rfq.Contract.UnconsumedMsg(&_Rfq.CallOpts, arg0)
}

// VerifySigOfQuoteHash is a free data retrieval call binding the contract method 0x2a37cf71.
//
// Solidity: function verifySigOfQuoteHash(address _liquidityProvider, bytes32 _quoteHash, bytes _sig) view returns()
func (_Rfq *RfqCaller) VerifySigOfQuoteHash(opts *bind.CallOpts, _liquidityProvider common.Address, _quoteHash [32]byte, _sig []byte) error {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "verifySigOfQuoteHash", _liquidityProvider, _quoteHash, _sig)

	if err != nil {
		return err
	}

	return err

}

// VerifySigOfQuoteHash is a free data retrieval call binding the contract method 0x2a37cf71.
//
// Solidity: function verifySigOfQuoteHash(address _liquidityProvider, bytes32 _quoteHash, bytes _sig) view returns()
func (_Rfq *RfqSession) VerifySigOfQuoteHash(_liquidityProvider common.Address, _quoteHash [32]byte, _sig []byte) error {
	return _Rfq.Contract.VerifySigOfQuoteHash(&_Rfq.CallOpts, _liquidityProvider, _quoteHash, _sig)
}

// VerifySigOfQuoteHash is a free data retrieval call binding the contract method 0x2a37cf71.
//
// Solidity: function verifySigOfQuoteHash(address _liquidityProvider, bytes32 _quoteHash, bytes _sig) view returns()
func (_Rfq *RfqCallerSession) VerifySigOfQuoteHash(_liquidityProvider common.Address, _quoteHash [32]byte, _sig []byte) error {
	return _Rfq.Contract.VerifySigOfQuoteHash(&_Rfq.CallOpts, _liquidityProvider, _quoteHash, _sig)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Rfq *RfqTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Rfq *RfqSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.AddGovernor(&_Rfq.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Rfq *RfqTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.AddGovernor(&_Rfq.TransactOpts, _account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Rfq *RfqTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Rfq *RfqSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.AddPauser(&_Rfq.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Rfq *RfqTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.AddPauser(&_Rfq.TransactOpts, account)
}

// CollectFee is a paid mutator transaction binding the contract method 0x69b59e75.
//
// Solidity: function collectFee(address _token) returns()
func (_Rfq *RfqTransactor) CollectFee(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "collectFee", _token)
}

// CollectFee is a paid mutator transaction binding the contract method 0x69b59e75.
//
// Solidity: function collectFee(address _token) returns()
func (_Rfq *RfqSession) CollectFee(_token common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.CollectFee(&_Rfq.TransactOpts, _token)
}

// CollectFee is a paid mutator transaction binding the contract method 0x69b59e75.
//
// Solidity: function collectFee(address _token) returns()
func (_Rfq *RfqTransactorSession) CollectFee(_token common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.CollectFee(&_Rfq.TransactOpts, _token)
}

// DstTransfer is a paid mutator transaction binding the contract method 0xa5ea10cb.
//
// Solidity: function dstTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactor) DstTransfer(opts *bind.TransactOpts, _quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "dstTransfer", _quote)
}

// DstTransfer is a paid mutator transaction binding the contract method 0xa5ea10cb.
//
// Solidity: function dstTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqSession) DstTransfer(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransfer(&_Rfq.TransactOpts, _quote)
}

// DstTransfer is a paid mutator transaction binding the contract method 0xa5ea10cb.
//
// Solidity: function dstTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactorSession) DstTransfer(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransfer(&_Rfq.TransactOpts, _quote)
}

// DstTransferNative is a paid mutator transaction binding the contract method 0x1000cd9e.
//
// Solidity: function dstTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactor) DstTransferNative(opts *bind.TransactOpts, _quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "dstTransferNative", _quote)
}

// DstTransferNative is a paid mutator transaction binding the contract method 0x1000cd9e.
//
// Solidity: function dstTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqSession) DstTransferNative(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransferNative(&_Rfq.TransactOpts, _quote)
}

// DstTransferNative is a paid mutator transaction binding the contract method 0x1000cd9e.
//
// Solidity: function dstTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactorSession) DstTransferNative(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransferNative(&_Rfq.TransactOpts, _quote)
}

// DstTransferWithSig is a paid mutator transaction binding the contract method 0xae6c218e.
//
// Solidity: function dstTransferWithSig((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _sig) payable returns()
func (_Rfq *RfqTransactor) DstTransferWithSig(opts *bind.TransactOpts, _quote RFQQuote, _sig []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "dstTransferWithSig", _quote, _sig)
}

// DstTransferWithSig is a paid mutator transaction binding the contract method 0xae6c218e.
//
// Solidity: function dstTransferWithSig((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _sig) payable returns()
func (_Rfq *RfqSession) DstTransferWithSig(_quote RFQQuote, _sig []byte) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransferWithSig(&_Rfq.TransactOpts, _quote, _sig)
}

// DstTransferWithSig is a paid mutator transaction binding the contract method 0xae6c218e.
//
// Solidity: function dstTransferWithSig((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _sig) payable returns()
func (_Rfq *RfqTransactorSession) DstTransferWithSig(_quote RFQQuote, _sig []byte) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransferWithSig(&_Rfq.TransactOpts, _quote, _sig)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessage(&_Rfq.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessage(&_Rfq.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address ) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessage0(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessage0", _sender, _srcChainId, _message, arg3)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address ) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessage0(_sender common.Address, _srcChainId uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessage0(&_Rfq.TransactOpts, _sender, _srcChainId, _message, arg3)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address ) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessage0(_sender common.Address, _srcChainId uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessage0(&_Rfq.TransactOpts, _sender, _srcChainId, _message, arg3)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransfer(&_Rfq.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransfer(&_Rfq.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransferFallback(&_Rfq.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransferFallback(&_Rfq.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransferRefund(&_Rfq.TransactOpts, _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransferRefund(&_Rfq.TransactOpts, _token, _amount, _message, _executor)
}

// ExecuteRefund is a paid mutator transaction binding the contract method 0xf8b839e5.
//
// Solidity: function executeRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactor) ExecuteRefund(opts *bind.TransactOpts, _quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeRefund", _quote, _execMsgCallData)
}

// ExecuteRefund is a paid mutator transaction binding the contract method 0xf8b839e5.
//
// Solidity: function executeRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqSession) ExecuteRefund(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteRefund(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// ExecuteRefund is a paid mutator transaction binding the contract method 0xf8b839e5.
//
// Solidity: function executeRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactorSession) ExecuteRefund(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteRefund(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// ExecuteRefundNative is a paid mutator transaction binding the contract method 0x6609870d.
//
// Solidity: function executeRefundNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactor) ExecuteRefundNative(opts *bind.TransactOpts, _quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeRefundNative", _quote, _execMsgCallData)
}

// ExecuteRefundNative is a paid mutator transaction binding the contract method 0x6609870d.
//
// Solidity: function executeRefundNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqSession) ExecuteRefundNative(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteRefundNative(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// ExecuteRefundNative is a paid mutator transaction binding the contract method 0x6609870d.
//
// Solidity: function executeRefundNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactorSession) ExecuteRefundNative(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteRefundNative(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rfq *RfqTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rfq *RfqSession) Pause() (*types.Transaction, error) {
	return _Rfq.Contract.Pause(&_Rfq.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rfq *RfqTransactorSession) Pause() (*types.Transaction, error) {
	return _Rfq.Contract.Pause(&_Rfq.TransactOpts)
}

// RegisterAllowedSigner is a paid mutator transaction binding the contract method 0x5aa15a4b.
//
// Solidity: function registerAllowedSigner(address _signer) returns()
func (_Rfq *RfqTransactor) RegisterAllowedSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "registerAllowedSigner", _signer)
}

// RegisterAllowedSigner is a paid mutator transaction binding the contract method 0x5aa15a4b.
//
// Solidity: function registerAllowedSigner(address _signer) returns()
func (_Rfq *RfqSession) RegisterAllowedSigner(_signer common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RegisterAllowedSigner(&_Rfq.TransactOpts, _signer)
}

// RegisterAllowedSigner is a paid mutator transaction binding the contract method 0x5aa15a4b.
//
// Solidity: function registerAllowedSigner(address _signer) returns()
func (_Rfq *RfqTransactorSession) RegisterAllowedSigner(_signer common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RegisterAllowedSigner(&_Rfq.TransactOpts, _signer)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Rfq *RfqTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Rfq *RfqSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RemoveGovernor(&_Rfq.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Rfq *RfqTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RemoveGovernor(&_Rfq.TransactOpts, _account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Rfq *RfqTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Rfq *RfqSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RemovePauser(&_Rfq.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Rfq *RfqTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RemovePauser(&_Rfq.TransactOpts, account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Rfq *RfqTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Rfq *RfqSession) RenounceGovernor() (*types.Transaction, error) {
	return _Rfq.Contract.RenounceGovernor(&_Rfq.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Rfq *RfqTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _Rfq.Contract.RenounceGovernor(&_Rfq.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Rfq *RfqTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Rfq *RfqSession) RenouncePauser() (*types.Transaction, error) {
	return _Rfq.Contract.RenouncePauser(&_Rfq.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Rfq *RfqTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Rfq.Contract.RenouncePauser(&_Rfq.TransactOpts)
}

// RequestRefund is a paid mutator transaction binding the contract method 0x79c7efd3.
//
// Solidity: function requestRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactor) RequestRefund(opts *bind.TransactOpts, _quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "requestRefund", _quote)
}

// RequestRefund is a paid mutator transaction binding the contract method 0x79c7efd3.
//
// Solidity: function requestRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqSession) RequestRefund(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.RequestRefund(&_Rfq.TransactOpts, _quote)
}

// RequestRefund is a paid mutator transaction binding the contract method 0x79c7efd3.
//
// Solidity: function requestRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactorSession) RequestRefund(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.RequestRefund(&_Rfq.TransactOpts, _quote)
}

// SameChainTransfer is a paid mutator transaction binding the contract method 0xc78e33a2.
//
// Solidity: function sameChainTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqTransactor) SameChainTransfer(opts *bind.TransactOpts, _quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "sameChainTransfer", _quote, _releaseNative)
}

// SameChainTransfer is a paid mutator transaction binding the contract method 0xc78e33a2.
//
// Solidity: function sameChainTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqSession) SameChainTransfer(_quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransfer(&_Rfq.TransactOpts, _quote, _releaseNative)
}

// SameChainTransfer is a paid mutator transaction binding the contract method 0xc78e33a2.
//
// Solidity: function sameChainTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqTransactorSession) SameChainTransfer(_quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransfer(&_Rfq.TransactOpts, _quote, _releaseNative)
}

// SameChainTransferNative is a paid mutator transaction binding the contract method 0x1e9c5748.
//
// Solidity: function sameChainTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqTransactor) SameChainTransferNative(opts *bind.TransactOpts, _quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "sameChainTransferNative", _quote, _releaseNative)
}

// SameChainTransferNative is a paid mutator transaction binding the contract method 0x1e9c5748.
//
// Solidity: function sameChainTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqSession) SameChainTransferNative(_quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransferNative(&_Rfq.TransactOpts, _quote, _releaseNative)
}

// SameChainTransferNative is a paid mutator transaction binding the contract method 0x1e9c5748.
//
// Solidity: function sameChainTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqTransactorSession) SameChainTransferNative(_quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransferNative(&_Rfq.TransactOpts, _quote, _releaseNative)
}

// SameChainTransferWithSig is a paid mutator transaction binding the contract method 0xb1a7d13f.
//
// Solidity: function sameChainTransferWithSig((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative, bytes _sig) payable returns()
func (_Rfq *RfqTransactor) SameChainTransferWithSig(opts *bind.TransactOpts, _quote RFQQuote, _releaseNative bool, _sig []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "sameChainTransferWithSig", _quote, _releaseNative, _sig)
}

// SameChainTransferWithSig is a paid mutator transaction binding the contract method 0xb1a7d13f.
//
// Solidity: function sameChainTransferWithSig((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative, bytes _sig) payable returns()
func (_Rfq *RfqSession) SameChainTransferWithSig(_quote RFQQuote, _releaseNative bool, _sig []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransferWithSig(&_Rfq.TransactOpts, _quote, _releaseNative, _sig)
}

// SameChainTransferWithSig is a paid mutator transaction binding the contract method 0xb1a7d13f.
//
// Solidity: function sameChainTransferWithSig((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative, bytes _sig) payable returns()
func (_Rfq *RfqTransactorSession) SameChainTransferWithSig(_quote RFQQuote, _releaseNative bool, _sig []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransferWithSig(&_Rfq.TransactOpts, _quote, _releaseNative, _sig)
}

// SetFeePerc is a paid mutator transaction binding the contract method 0xab9341fd.
//
// Solidity: function setFeePerc(uint64[] _chainIds, uint32[] _feePercs) returns()
func (_Rfq *RfqTransactor) SetFeePerc(opts *bind.TransactOpts, _chainIds []uint64, _feePercs []uint32) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "setFeePerc", _chainIds, _feePercs)
}

// SetFeePerc is a paid mutator transaction binding the contract method 0xab9341fd.
//
// Solidity: function setFeePerc(uint64[] _chainIds, uint32[] _feePercs) returns()
func (_Rfq *RfqSession) SetFeePerc(_chainIds []uint64, _feePercs []uint32) (*types.Transaction, error) {
	return _Rfq.Contract.SetFeePerc(&_Rfq.TransactOpts, _chainIds, _feePercs)
}

// SetFeePerc is a paid mutator transaction binding the contract method 0xab9341fd.
//
// Solidity: function setFeePerc(uint64[] _chainIds, uint32[] _feePercs) returns()
func (_Rfq *RfqTransactorSession) SetFeePerc(_chainIds []uint64, _feePercs []uint32) (*types.Transaction, error) {
	return _Rfq.Contract.SetFeePerc(&_Rfq.TransactOpts, _chainIds, _feePercs)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_Rfq *RfqTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_Rfq *RfqSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetMessageBus(&_Rfq.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_Rfq *RfqTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetMessageBus(&_Rfq.TransactOpts, _messageBus)
}

// SetNativeWrap is a paid mutator transaction binding the contract method 0x5b5a66a7.
//
// Solidity: function setNativeWrap(address _nativeWrap) returns()
func (_Rfq *RfqTransactor) SetNativeWrap(opts *bind.TransactOpts, _nativeWrap common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "setNativeWrap", _nativeWrap)
}

// SetNativeWrap is a paid mutator transaction binding the contract method 0x5b5a66a7.
//
// Solidity: function setNativeWrap(address _nativeWrap) returns()
func (_Rfq *RfqSession) SetNativeWrap(_nativeWrap common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetNativeWrap(&_Rfq.TransactOpts, _nativeWrap)
}

// SetNativeWrap is a paid mutator transaction binding the contract method 0x5b5a66a7.
//
// Solidity: function setNativeWrap(address _nativeWrap) returns()
func (_Rfq *RfqTransactorSession) SetNativeWrap(_nativeWrap common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetNativeWrap(&_Rfq.TransactOpts, _nativeWrap)
}

// SetRemoteRfqContracts is a paid mutator transaction binding the contract method 0xcbac44df.
//
// Solidity: function setRemoteRfqContracts(uint64[] _chainIds, address[] _remoteRfqContracts) returns()
func (_Rfq *RfqTransactor) SetRemoteRfqContracts(opts *bind.TransactOpts, _chainIds []uint64, _remoteRfqContracts []common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "setRemoteRfqContracts", _chainIds, _remoteRfqContracts)
}

// SetRemoteRfqContracts is a paid mutator transaction binding the contract method 0xcbac44df.
//
// Solidity: function setRemoteRfqContracts(uint64[] _chainIds, address[] _remoteRfqContracts) returns()
func (_Rfq *RfqSession) SetRemoteRfqContracts(_chainIds []uint64, _remoteRfqContracts []common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetRemoteRfqContracts(&_Rfq.TransactOpts, _chainIds, _remoteRfqContracts)
}

// SetRemoteRfqContracts is a paid mutator transaction binding the contract method 0xcbac44df.
//
// Solidity: function setRemoteRfqContracts(uint64[] _chainIds, address[] _remoteRfqContracts) returns()
func (_Rfq *RfqTransactorSession) SetRemoteRfqContracts(_chainIds []uint64, _remoteRfqContracts []common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetRemoteRfqContracts(&_Rfq.TransactOpts, _chainIds, _remoteRfqContracts)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Rfq *RfqTransactor) SetTreasuryAddr(opts *bind.TransactOpts, _treasuryAddr common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "setTreasuryAddr", _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Rfq *RfqSession) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetTreasuryAddr(&_Rfq.TransactOpts, _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Rfq *RfqTransactorSession) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetTreasuryAddr(&_Rfq.TransactOpts, _treasuryAddr)
}

// SrcDeposit is a paid mutator transaction binding the contract method 0xed9830bb.
//
// Solidity: function srcDeposit((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqTransactor) SrcDeposit(opts *bind.TransactOpts, _quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "srcDeposit", _quote, _submissionDeadline)
}

// SrcDeposit is a paid mutator transaction binding the contract method 0xed9830bb.
//
// Solidity: function srcDeposit((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqSession) SrcDeposit(_quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.Contract.SrcDeposit(&_Rfq.TransactOpts, _quote, _submissionDeadline)
}

// SrcDeposit is a paid mutator transaction binding the contract method 0xed9830bb.
//
// Solidity: function srcDeposit((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqTransactorSession) SrcDeposit(_quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.Contract.SrcDeposit(&_Rfq.TransactOpts, _quote, _submissionDeadline)
}

// SrcDepositNative is a paid mutator transaction binding the contract method 0x0a54aacd.
//
// Solidity: function srcDepositNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqTransactor) SrcDepositNative(opts *bind.TransactOpts, _quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "srcDepositNative", _quote, _submissionDeadline)
}

// SrcDepositNative is a paid mutator transaction binding the contract method 0x0a54aacd.
//
// Solidity: function srcDepositNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqSession) SrcDepositNative(_quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.Contract.SrcDepositNative(&_Rfq.TransactOpts, _quote, _submissionDeadline)
}

// SrcDepositNative is a paid mutator transaction binding the contract method 0x0a54aacd.
//
// Solidity: function srcDepositNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqTransactorSession) SrcDepositNative(_quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.Contract.SrcDepositNative(&_Rfq.TransactOpts, _quote, _submissionDeadline)
}

// SrcRelease is a paid mutator transaction binding the contract method 0xa2bdb899.
//
// Solidity: function srcRelease((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactor) SrcRelease(opts *bind.TransactOpts, _quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "srcRelease", _quote, _execMsgCallData)
}

// SrcRelease is a paid mutator transaction binding the contract method 0xa2bdb899.
//
// Solidity: function srcRelease((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqSession) SrcRelease(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SrcRelease(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// SrcRelease is a paid mutator transaction binding the contract method 0xa2bdb899.
//
// Solidity: function srcRelease((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactorSession) SrcRelease(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SrcRelease(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// SrcReleaseNative is a paid mutator transaction binding the contract method 0xaf4ab138.
//
// Solidity: function srcReleaseNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactor) SrcReleaseNative(opts *bind.TransactOpts, _quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "srcReleaseNative", _quote, _execMsgCallData)
}

// SrcReleaseNative is a paid mutator transaction binding the contract method 0xaf4ab138.
//
// Solidity: function srcReleaseNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqSession) SrcReleaseNative(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SrcReleaseNative(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// SrcReleaseNative is a paid mutator transaction binding the contract method 0xaf4ab138.
//
// Solidity: function srcReleaseNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactorSession) SrcReleaseNative(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SrcReleaseNative(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rfq *RfqTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rfq *RfqSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.TransferOwnership(&_Rfq.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rfq *RfqTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.TransferOwnership(&_Rfq.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Rfq *RfqTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Rfq *RfqSession) Unpause() (*types.Transaction, error) {
	return _Rfq.Contract.Unpause(&_Rfq.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Rfq *RfqTransactorSession) Unpause() (*types.Transaction, error) {
	return _Rfq.Contract.Unpause(&_Rfq.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rfq *RfqTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rfq *RfqSession) Receive() (*types.Transaction, error) {
	return _Rfq.Contract.Receive(&_Rfq.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rfq *RfqTransactorSession) Receive() (*types.Transaction, error) {
	return _Rfq.Contract.Receive(&_Rfq.TransactOpts)
}

// RfqDstTransferredIterator is returned from FilterDstTransferred and is used to iterate over the raw logs and unpacked data for DstTransferred events raised by the Rfq contract.
type RfqDstTransferredIterator struct {
	Event *RfqDstTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqDstTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqDstTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqDstTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqDstTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqDstTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqDstTransferred represents a DstTransferred event raised by the Rfq contract.
type RfqDstTransferred struct {
	QuoteHash [32]byte
	Receiver  common.Address
	DstToken  common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDstTransferred is a free log retrieval operation binding the contract event 0xb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc5.
//
// Solidity: event DstTransferred(bytes32 quoteHash, address receiver, address dstToken, uint256 amount)
func (_Rfq *RfqFilterer) FilterDstTransferred(opts *bind.FilterOpts) (*RfqDstTransferredIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "DstTransferred")
	if err != nil {
		return nil, err
	}
	return &RfqDstTransferredIterator{contract: _Rfq.contract, event: "DstTransferred", logs: logs, sub: sub}, nil
}

// WatchDstTransferred is a free log subscription operation binding the contract event 0xb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc5.
//
// Solidity: event DstTransferred(bytes32 quoteHash, address receiver, address dstToken, uint256 amount)
func (_Rfq *RfqFilterer) WatchDstTransferred(opts *bind.WatchOpts, sink chan<- *RfqDstTransferred) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "DstTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqDstTransferred)
				if err := _Rfq.contract.UnpackLog(event, "DstTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDstTransferred is a log parse operation binding the contract event 0xb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc5.
//
// Solidity: event DstTransferred(bytes32 quoteHash, address receiver, address dstToken, uint256 amount)
func (_Rfq *RfqFilterer) ParseDstTransferred(log types.Log) (*RfqDstTransferred, error) {
	event := new(RfqDstTransferred)
	if err := _Rfq.contract.UnpackLog(event, "DstTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the Rfq contract.
type RfqFeeCollectedIterator struct {
	Event *RfqFeeCollected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqFeeCollected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqFeeCollected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqFeeCollected represents a FeeCollected event raised by the Rfq contract.
type RfqFeeCollected struct {
	TreasuryAddr common.Address
	Token        common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeCollected is a free log retrieval operation binding the contract event 0xf228de527fc1b9843baac03b9a04565473a263375950e63435d4138464386f46.
//
// Solidity: event FeeCollected(address treasuryAddr, address token, uint256 amount)
func (_Rfq *RfqFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*RfqFeeCollectedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &RfqFeeCollectedIterator{contract: _Rfq.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0xf228de527fc1b9843baac03b9a04565473a263375950e63435d4138464386f46.
//
// Solidity: event FeeCollected(address treasuryAddr, address token, uint256 amount)
func (_Rfq *RfqFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *RfqFeeCollected) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqFeeCollected)
				if err := _Rfq.contract.UnpackLog(event, "FeeCollected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeCollected is a log parse operation binding the contract event 0xf228de527fc1b9843baac03b9a04565473a263375950e63435d4138464386f46.
//
// Solidity: event FeeCollected(address treasuryAddr, address token, uint256 amount)
func (_Rfq *RfqFilterer) ParseFeeCollected(log types.Log) (*RfqFeeCollected, error) {
	event := new(RfqFeeCollected)
	if err := _Rfq.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqFeePercUpdatedIterator is returned from FilterFeePercUpdated and is used to iterate over the raw logs and unpacked data for FeePercUpdated events raised by the Rfq contract.
type RfqFeePercUpdatedIterator struct {
	Event *RfqFeePercUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqFeePercUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqFeePercUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqFeePercUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqFeePercUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqFeePercUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqFeePercUpdated represents a FeePercUpdated event raised by the Rfq contract.
type RfqFeePercUpdated struct {
	ChainIds []uint64
	FeePercs []uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeePercUpdated is a free log retrieval operation binding the contract event 0x541df5e570cf10ffe04899eebd9eebebd1c54e2bd4af9f24b23fb4a40c6ea00b.
//
// Solidity: event FeePercUpdated(uint64[] chainIds, uint32[] feePercs)
func (_Rfq *RfqFilterer) FilterFeePercUpdated(opts *bind.FilterOpts) (*RfqFeePercUpdatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "FeePercUpdated")
	if err != nil {
		return nil, err
	}
	return &RfqFeePercUpdatedIterator{contract: _Rfq.contract, event: "FeePercUpdated", logs: logs, sub: sub}, nil
}

// WatchFeePercUpdated is a free log subscription operation binding the contract event 0x541df5e570cf10ffe04899eebd9eebebd1c54e2bd4af9f24b23fb4a40c6ea00b.
//
// Solidity: event FeePercUpdated(uint64[] chainIds, uint32[] feePercs)
func (_Rfq *RfqFilterer) WatchFeePercUpdated(opts *bind.WatchOpts, sink chan<- *RfqFeePercUpdated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "FeePercUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqFeePercUpdated)
				if err := _Rfq.contract.UnpackLog(event, "FeePercUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeePercUpdated is a log parse operation binding the contract event 0x541df5e570cf10ffe04899eebd9eebebd1c54e2bd4af9f24b23fb4a40c6ea00b.
//
// Solidity: event FeePercUpdated(uint64[] chainIds, uint32[] feePercs)
func (_Rfq *RfqFilterer) ParseFeePercUpdated(log types.Log) (*RfqFeePercUpdated, error) {
	event := new(RfqFeePercUpdated)
	if err := _Rfq.contract.UnpackLog(event, "FeePercUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the Rfq contract.
type RfqGovernorAddedIterator struct {
	Event *RfqGovernorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqGovernorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqGovernorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqGovernorAdded represents a GovernorAdded event raised by the Rfq contract.
type RfqGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Rfq *RfqFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*RfqGovernorAddedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &RfqGovernorAddedIterator{contract: _Rfq.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Rfq *RfqFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *RfqGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqGovernorAdded)
				if err := _Rfq.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Rfq *RfqFilterer) ParseGovernorAdded(log types.Log) (*RfqGovernorAdded, error) {
	event := new(RfqGovernorAdded)
	if err := _Rfq.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the Rfq contract.
type RfqGovernorRemovedIterator struct {
	Event *RfqGovernorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqGovernorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqGovernorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqGovernorRemoved represents a GovernorRemoved event raised by the Rfq contract.
type RfqGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Rfq *RfqFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*RfqGovernorRemovedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &RfqGovernorRemovedIterator{contract: _Rfq.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Rfq *RfqFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *RfqGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqGovernorRemoved)
				if err := _Rfq.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Rfq *RfqFilterer) ParseGovernorRemoved(log types.Log) (*RfqGovernorRemoved, error) {
	event := new(RfqGovernorRemoved)
	if err := _Rfq.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqMessageBusUpdatedIterator is returned from FilterMessageBusUpdated and is used to iterate over the raw logs and unpacked data for MessageBusUpdated events raised by the Rfq contract.
type RfqMessageBusUpdatedIterator struct {
	Event *RfqMessageBusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqMessageBusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqMessageBusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqMessageBusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqMessageBusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqMessageBusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqMessageBusUpdated represents a MessageBusUpdated event raised by the Rfq contract.
type RfqMessageBusUpdated struct {
	MessageBus common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageBusUpdated is a free log retrieval operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_Rfq *RfqFilterer) FilterMessageBusUpdated(opts *bind.FilterOpts) (*RfqMessageBusUpdatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "MessageBusUpdated")
	if err != nil {
		return nil, err
	}
	return &RfqMessageBusUpdatedIterator{contract: _Rfq.contract, event: "MessageBusUpdated", logs: logs, sub: sub}, nil
}

// WatchMessageBusUpdated is a free log subscription operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_Rfq *RfqFilterer) WatchMessageBusUpdated(opts *bind.WatchOpts, sink chan<- *RfqMessageBusUpdated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "MessageBusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqMessageBusUpdated)
				if err := _Rfq.contract.UnpackLog(event, "MessageBusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageBusUpdated is a log parse operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_Rfq *RfqFilterer) ParseMessageBusUpdated(log types.Log) (*RfqMessageBusUpdated, error) {
	event := new(RfqMessageBusUpdated)
	if err := _Rfq.contract.UnpackLog(event, "MessageBusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rfq contract.
type RfqOwnershipTransferredIterator struct {
	Event *RfqOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqOwnershipTransferred represents a OwnershipTransferred event raised by the Rfq contract.
type RfqOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rfq *RfqFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RfqOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RfqOwnershipTransferredIterator{contract: _Rfq.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rfq *RfqFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RfqOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqOwnershipTransferred)
				if err := _Rfq.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rfq *RfqFilterer) ParseOwnershipTransferred(log types.Log) (*RfqOwnershipTransferred, error) {
	event := new(RfqOwnershipTransferred)
	if err := _Rfq.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Rfq contract.
type RfqPausedIterator struct {
	Event *RfqPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqPaused represents a Paused event raised by the Rfq contract.
type RfqPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rfq *RfqFilterer) FilterPaused(opts *bind.FilterOpts) (*RfqPausedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RfqPausedIterator{contract: _Rfq.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rfq *RfqFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RfqPaused) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqPaused)
				if err := _Rfq.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rfq *RfqFilterer) ParsePaused(log types.Log) (*RfqPaused, error) {
	event := new(RfqPaused)
	if err := _Rfq.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Rfq contract.
type RfqPauserAddedIterator struct {
	Event *RfqPauserAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqPauserAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqPauserAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqPauserAdded represents a PauserAdded event raised by the Rfq contract.
type RfqPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Rfq *RfqFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*RfqPauserAddedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &RfqPauserAddedIterator{contract: _Rfq.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Rfq *RfqFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *RfqPauserAdded) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqPauserAdded)
				if err := _Rfq.contract.UnpackLog(event, "PauserAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Rfq *RfqFilterer) ParsePauserAdded(log types.Log) (*RfqPauserAdded, error) {
	event := new(RfqPauserAdded)
	if err := _Rfq.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Rfq contract.
type RfqPauserRemovedIterator struct {
	Event *RfqPauserRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqPauserRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqPauserRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqPauserRemoved represents a PauserRemoved event raised by the Rfq contract.
type RfqPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Rfq *RfqFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*RfqPauserRemovedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &RfqPauserRemovedIterator{contract: _Rfq.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Rfq *RfqFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *RfqPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqPauserRemoved)
				if err := _Rfq.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Rfq *RfqFilterer) ParsePauserRemoved(log types.Log) (*RfqPauserRemoved, error) {
	event := new(RfqPauserRemoved)
	if err := _Rfq.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqRefundInitiatedIterator is returned from FilterRefundInitiated and is used to iterate over the raw logs and unpacked data for RefundInitiated events raised by the Rfq contract.
type RfqRefundInitiatedIterator struct {
	Event *RfqRefundInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqRefundInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqRefundInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqRefundInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqRefundInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqRefundInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqRefundInitiated represents a RefundInitiated event raised by the Rfq contract.
type RfqRefundInitiated struct {
	QuoteHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundInitiated is a free log retrieval operation binding the contract event 0x7cdd4403cff3a09d96c1ffe4ad1cc5c195e9053463a55edfc2944644ec022118.
//
// Solidity: event RefundInitiated(bytes32 quoteHash)
func (_Rfq *RfqFilterer) FilterRefundInitiated(opts *bind.FilterOpts) (*RfqRefundInitiatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "RefundInitiated")
	if err != nil {
		return nil, err
	}
	return &RfqRefundInitiatedIterator{contract: _Rfq.contract, event: "RefundInitiated", logs: logs, sub: sub}, nil
}

// WatchRefundInitiated is a free log subscription operation binding the contract event 0x7cdd4403cff3a09d96c1ffe4ad1cc5c195e9053463a55edfc2944644ec022118.
//
// Solidity: event RefundInitiated(bytes32 quoteHash)
func (_Rfq *RfqFilterer) WatchRefundInitiated(opts *bind.WatchOpts, sink chan<- *RfqRefundInitiated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "RefundInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqRefundInitiated)
				if err := _Rfq.contract.UnpackLog(event, "RefundInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundInitiated is a log parse operation binding the contract event 0x7cdd4403cff3a09d96c1ffe4ad1cc5c195e9053463a55edfc2944644ec022118.
//
// Solidity: event RefundInitiated(bytes32 quoteHash)
func (_Rfq *RfqFilterer) ParseRefundInitiated(log types.Log) (*RfqRefundInitiated, error) {
	event := new(RfqRefundInitiated)
	if err := _Rfq.contract.UnpackLog(event, "RefundInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the Rfq contract.
type RfqRefundedIterator struct {
	Event *RfqRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqRefunded represents a Refunded event raised by the Rfq contract.
type RfqRefunded struct {
	QuoteHash [32]byte
	RefundTo  common.Address
	SrcToken  common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x2e0668a62a5f556368dca9c7113e20f2852c05155548243804bf714ce72b25a6.
//
// Solidity: event Refunded(bytes32 quoteHash, address refundTo, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) FilterRefunded(opts *bind.FilterOpts) (*RfqRefundedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return &RfqRefundedIterator{contract: _Rfq.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x2e0668a62a5f556368dca9c7113e20f2852c05155548243804bf714ce72b25a6.
//
// Solidity: event Refunded(bytes32 quoteHash, address refundTo, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *RfqRefunded) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqRefunded)
				if err := _Rfq.contract.UnpackLog(event, "Refunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefunded is a log parse operation binding the contract event 0x2e0668a62a5f556368dca9c7113e20f2852c05155548243804bf714ce72b25a6.
//
// Solidity: event Refunded(bytes32 quoteHash, address refundTo, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) ParseRefunded(log types.Log) (*RfqRefunded, error) {
	event := new(RfqRefunded)
	if err := _Rfq.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqRfqContractsUpdatedIterator is returned from FilterRfqContractsUpdated and is used to iterate over the raw logs and unpacked data for RfqContractsUpdated events raised by the Rfq contract.
type RfqRfqContractsUpdatedIterator struct {
	Event *RfqRfqContractsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqRfqContractsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqRfqContractsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqRfqContractsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqRfqContractsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqRfqContractsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqRfqContractsUpdated represents a RfqContractsUpdated event raised by the Rfq contract.
type RfqRfqContractsUpdated struct {
	ChainIds           []uint64
	RemoteRfqContracts []common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRfqContractsUpdated is a free log retrieval operation binding the contract event 0xb4739c640c5970d8ce88b6c31f3706099efca660e282d47b0a267a0bb572d8b7.
//
// Solidity: event RfqContractsUpdated(uint64[] chainIds, address[] remoteRfqContracts)
func (_Rfq *RfqFilterer) FilterRfqContractsUpdated(opts *bind.FilterOpts) (*RfqRfqContractsUpdatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "RfqContractsUpdated")
	if err != nil {
		return nil, err
	}
	return &RfqRfqContractsUpdatedIterator{contract: _Rfq.contract, event: "RfqContractsUpdated", logs: logs, sub: sub}, nil
}

// WatchRfqContractsUpdated is a free log subscription operation binding the contract event 0xb4739c640c5970d8ce88b6c31f3706099efca660e282d47b0a267a0bb572d8b7.
//
// Solidity: event RfqContractsUpdated(uint64[] chainIds, address[] remoteRfqContracts)
func (_Rfq *RfqFilterer) WatchRfqContractsUpdated(opts *bind.WatchOpts, sink chan<- *RfqRfqContractsUpdated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "RfqContractsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqRfqContractsUpdated)
				if err := _Rfq.contract.UnpackLog(event, "RfqContractsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRfqContractsUpdated is a log parse operation binding the contract event 0xb4739c640c5970d8ce88b6c31f3706099efca660e282d47b0a267a0bb572d8b7.
//
// Solidity: event RfqContractsUpdated(uint64[] chainIds, address[] remoteRfqContracts)
func (_Rfq *RfqFilterer) ParseRfqContractsUpdated(log types.Log) (*RfqRfqContractsUpdated, error) {
	event := new(RfqRfqContractsUpdated)
	if err := _Rfq.contract.UnpackLog(event, "RfqContractsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqSrcDepositedIterator is returned from FilterSrcDeposited and is used to iterate over the raw logs and unpacked data for SrcDeposited events raised by the Rfq contract.
type RfqSrcDepositedIterator struct {
	Event *RfqSrcDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqSrcDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqSrcDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqSrcDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqSrcDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqSrcDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqSrcDeposited represents a SrcDeposited event raised by the Rfq contract.
type RfqSrcDeposited struct {
	QuoteHash [32]byte
	Quote     RFQQuote
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSrcDeposited is a free log retrieval operation binding the contract event 0x3e4de2d1674631d426ae2a89635b421e6d40a31d27681afdf0eed67e81d07bcb.
//
// Solidity: event SrcDeposited(bytes32 quoteHash, (uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) quote)
func (_Rfq *RfqFilterer) FilterSrcDeposited(opts *bind.FilterOpts) (*RfqSrcDepositedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "SrcDeposited")
	if err != nil {
		return nil, err
	}
	return &RfqSrcDepositedIterator{contract: _Rfq.contract, event: "SrcDeposited", logs: logs, sub: sub}, nil
}

// WatchSrcDeposited is a free log subscription operation binding the contract event 0x3e4de2d1674631d426ae2a89635b421e6d40a31d27681afdf0eed67e81d07bcb.
//
// Solidity: event SrcDeposited(bytes32 quoteHash, (uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) quote)
func (_Rfq *RfqFilterer) WatchSrcDeposited(opts *bind.WatchOpts, sink chan<- *RfqSrcDeposited) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "SrcDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqSrcDeposited)
				if err := _Rfq.contract.UnpackLog(event, "SrcDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSrcDeposited is a log parse operation binding the contract event 0x3e4de2d1674631d426ae2a89635b421e6d40a31d27681afdf0eed67e81d07bcb.
//
// Solidity: event SrcDeposited(bytes32 quoteHash, (uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) quote)
func (_Rfq *RfqFilterer) ParseSrcDeposited(log types.Log) (*RfqSrcDeposited, error) {
	event := new(RfqSrcDeposited)
	if err := _Rfq.contract.UnpackLog(event, "SrcDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqSrcReleasedIterator is returned from FilterSrcReleased and is used to iterate over the raw logs and unpacked data for SrcReleased events raised by the Rfq contract.
type RfqSrcReleasedIterator struct {
	Event *RfqSrcReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqSrcReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqSrcReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqSrcReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqSrcReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqSrcReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqSrcReleased represents a SrcReleased event raised by the Rfq contract.
type RfqSrcReleased struct {
	QuoteHash         [32]byte
	LiquidityProvider common.Address
	SrcToken          common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSrcReleased is a free log retrieval operation binding the contract event 0xf29b32a17c591b8b3b1216ce0ffb67c07f3478e99b50c5ccf8602878b1ee6376.
//
// Solidity: event SrcReleased(bytes32 quoteHash, address liquidityProvider, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) FilterSrcReleased(opts *bind.FilterOpts) (*RfqSrcReleasedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "SrcReleased")
	if err != nil {
		return nil, err
	}
	return &RfqSrcReleasedIterator{contract: _Rfq.contract, event: "SrcReleased", logs: logs, sub: sub}, nil
}

// WatchSrcReleased is a free log subscription operation binding the contract event 0xf29b32a17c591b8b3b1216ce0ffb67c07f3478e99b50c5ccf8602878b1ee6376.
//
// Solidity: event SrcReleased(bytes32 quoteHash, address liquidityProvider, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) WatchSrcReleased(opts *bind.WatchOpts, sink chan<- *RfqSrcReleased) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "SrcReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqSrcReleased)
				if err := _Rfq.contract.UnpackLog(event, "SrcReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSrcReleased is a log parse operation binding the contract event 0xf29b32a17c591b8b3b1216ce0ffb67c07f3478e99b50c5ccf8602878b1ee6376.
//
// Solidity: event SrcReleased(bytes32 quoteHash, address liquidityProvider, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) ParseSrcReleased(log types.Log) (*RfqSrcReleased, error) {
	event := new(RfqSrcReleased)
	if err := _Rfq.contract.UnpackLog(event, "SrcReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqTreasuryAddrUpdatedIterator is returned from FilterTreasuryAddrUpdated and is used to iterate over the raw logs and unpacked data for TreasuryAddrUpdated events raised by the Rfq contract.
type RfqTreasuryAddrUpdatedIterator struct {
	Event *RfqTreasuryAddrUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqTreasuryAddrUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqTreasuryAddrUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqTreasuryAddrUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqTreasuryAddrUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqTreasuryAddrUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqTreasuryAddrUpdated represents a TreasuryAddrUpdated event raised by the Rfq contract.
type RfqTreasuryAddrUpdated struct {
	TreasuryAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTreasuryAddrUpdated is a free log retrieval operation binding the contract event 0xb17659014001857e7557191ad74dc9e967b181eaed0895975325e3848debbc42.
//
// Solidity: event TreasuryAddrUpdated(address treasuryAddr)
func (_Rfq *RfqFilterer) FilterTreasuryAddrUpdated(opts *bind.FilterOpts) (*RfqTreasuryAddrUpdatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "TreasuryAddrUpdated")
	if err != nil {
		return nil, err
	}
	return &RfqTreasuryAddrUpdatedIterator{contract: _Rfq.contract, event: "TreasuryAddrUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryAddrUpdated is a free log subscription operation binding the contract event 0xb17659014001857e7557191ad74dc9e967b181eaed0895975325e3848debbc42.
//
// Solidity: event TreasuryAddrUpdated(address treasuryAddr)
func (_Rfq *RfqFilterer) WatchTreasuryAddrUpdated(opts *bind.WatchOpts, sink chan<- *RfqTreasuryAddrUpdated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "TreasuryAddrUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqTreasuryAddrUpdated)
				if err := _Rfq.contract.UnpackLog(event, "TreasuryAddrUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasuryAddrUpdated is a log parse operation binding the contract event 0xb17659014001857e7557191ad74dc9e967b181eaed0895975325e3848debbc42.
//
// Solidity: event TreasuryAddrUpdated(address treasuryAddr)
func (_Rfq *RfqFilterer) ParseTreasuryAddrUpdated(log types.Log) (*RfqTreasuryAddrUpdated, error) {
	event := new(RfqTreasuryAddrUpdated)
	if err := _Rfq.contract.UnpackLog(event, "TreasuryAddrUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Rfq contract.
type RfqUnpausedIterator struct {
	Event *RfqUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RfqUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RfqUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RfqUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqUnpaused represents a Unpaused event raised by the Rfq contract.
type RfqUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rfq *RfqFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RfqUnpausedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RfqUnpausedIterator{contract: _Rfq.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rfq *RfqFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RfqUnpaused) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqUnpaused)
				if err := _Rfq.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rfq *RfqFilterer) ParseUnpaused(log types.Log) (*RfqUnpaused, error) {
	event := new(RfqUnpaused)
	if err := _Rfq.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
