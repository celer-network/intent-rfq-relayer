package rfqmm

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"sync"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/peti-rfq-relayer/relayer/common"
	"github.com/celer-network/peti-rfq-relayer/relayer/eth"
	"github.com/celer-network/peti-rfq-relayer/relayer/service/rfqmm/proto"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

type LiqManager struct {
	LP  *LiqProvider
	LPs map[uint64]*LiqProvider
}

func (d *LiqManager) ReleaseNative(chainId uint64) (bool, error) {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return false, err
	}
	return lp.releaseNative, nil
}

func (d *LiqManager) GetLP(chainId uint64) (*LiqProvider, error) {
	if lp, found := d.LPs[chainId]; !found {
		return nil, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("no liquidity provider on %d", chainId))
	} else {
		return lp, nil
	}
}

type LiqProvider struct {
	mux     sync.RWMutex
	signer  ethutils.Signer
	chainId uint64
	address eth.Addr
	liqs    map[string]*Liquidity
	// sorted slice by LiqOpDetail.Until in ascending order
	liqOps []*LiqOpDetail
	// to minimize searching cost when doing unfreeze
	hashToUntil   map[eth.Hash]int64
	releaseNative bool
}

type LiqOpDetail struct {
	Type   int
	Until  int64
	Token  string
	Amount *big.Int
	Hash   eth.Hash
}

type LPConfig struct {
	Keystore      string
	Passphrase    string
	ReleaseNative bool
}

type LiquidityConfig struct {
	Address    string
	Symbol     string
	Amount     string
	Approve    string
	Decimals   int32
	FreezeTime int64
}

var _ RequestSigner = &LiqProvider{}

func (lp *LiqProvider) Sign(data []byte) ([]byte, error) {
	sig, err := lp.signer.SignEthMessage(data)
	if err != nil {
		return nil, proto.NewErr(proto.ErrCode_ERROR_REQUEST_SIGNER, err.Error())
	}
	return sig, nil
}

func (lp *LiqProvider) Verify(data, sig []byte) bool {
	addr, err := ethutils.RecoverSigner(data, sig)
	if err != nil {
		return false
	}
	if lp.address != addr {
		return false
	}
	return true
}

type Liquidity struct {
	amount     *big.Int
	reserved   *big.Int
	confirmed  *big.Int
	approved   *big.Int
	token      *common.Token
	freezeTime int64
}

func createSigner(ksfile, passphrase string, chainid *big.Int) (ethutils.Signer, eth.Addr, error) {
	if strings.HasPrefix(ksfile, "awskms") {
		kmskeyinfo := strings.SplitN(ksfile, ":", 3)
		if len(kmskeyinfo) != 3 {
			return nil, eth.ZeroAddr, fmt.Errorf("%s has wrong format", ksfile)
		}
		awskeysec := []string{"", ""}
		if passphrase != "" {
			awskeysec = strings.SplitN(passphrase, ":", 2)
			if len(awskeysec) != 2 {
				return nil, eth.ZeroAddr, fmt.Errorf("%s has wrong format", passphrase)
			}
		}
		kmsSigner, err := ethutils.NewKmsSigner(kmskeyinfo[1], kmskeyinfo[2], awskeysec[0], awskeysec[1], chainid)
		if err != nil {
			return nil, eth.ZeroAddr, err
		}
		return kmsSigner, kmsSigner.Addr, nil
	}
	ksBytes, err := ioutil.ReadFile(ksfile)
	if err != nil {
		return nil, eth.ZeroAddr, err
	}
	key, err := keystore.DecryptKey(ksBytes, passphrase)
	if err != nil {
		return nil, eth.ZeroAddr, err
	}
	signer, err := ethutils.NewSigner(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)), chainid)
	return signer, key.Address, err
}
