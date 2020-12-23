// @Title
// @Description
// @Author  Niels  2020/9/11
package txs

import (
	"encoding/hex"
	"errors"
	"fmt"
	acc2 "github.com/niels1286/nerve-go-sdk/acc"
	"github.com/niels1286/nerve-go-sdk/api"
	txprotocal "github.com/niels1286/nerve-go-sdk/protocal"
	"github.com/niels1286/nerve-go-sdk/protocal/txdata"
	"math/big"
	"time"
)

const Default_asset_id = uint16(1)
const POCLockValue = 18446744073709551615

type TxSDK interface {
	Transfer(froms []*txprotocal.Coin, tos []*txprotocal.Coin, remark string, prikeyHexArray []string) (string, error)
	Staking(prikeyHex string, assetsChainId uint16, assetsId uint16, amount *big.Int, timeType uint8) (string, error)
	ExitStaking(prikeyHex, stakingTxHash string) (string, error)
}
type NerveTxSDK struct {
	chainId uint16
	prefix  string
	apiUrl  string
}

func GetTxSDK(apiUrl string, chainId uint16, prefix string) TxSDK {
	var sdk = &NerveTxSDK{
		apiUrl:  apiUrl,
		chainId: chainId,
		prefix:  prefix,
	}
	return sdk
}

func (sdk NerveTxSDK) Transfer(froms []*txprotocal.Coin, tos []*txprotocal.Coin, remark string, prikeyHexArray []string) (string, error) {
	//todo
	return "", nil
}

//timeType：0-current, 1-3 months, 2-6 months, 3-1 years, 4-2 years, 5-3 years, 6-5 years, 7-10 years
func (sdk NerveTxSDK) Staking(prikeyHex string, assetsChainId uint16, assetsId uint16, amount *big.Int, timeType uint8) (string, error) {
	accSdk := acc2.GetAccountSDK(sdk.chainId, sdk.prefix)
	apiSdk := api.GetApiSDK(sdk.apiUrl, sdk.chainId, sdk.prefix)
	prikey, err := hex.DecodeString(prikeyHex)
	if err != nil {
		return "", err
	}
	acc, err := accSdk.ImportAccount(prikey)
	if err != nil {
		return "", err
	}
	tx := &txprotocal.Transaction{
		TxType:   txprotocal.TX_TYPE_DEPOSIT,
		Time:     uint32(time.Now().Unix()),
		Remark:   nil,
		Extend:   nil,
		CoinData: nil,
		SignData: nil,
	}
	//账户nonce获取，余额获取
	accountInfo, err := apiSdk.GetBalance(acc.GetAddr(), sdk.chainId, Default_asset_id)
	if err != nil {
		return "", err
	}
	froms := []txprotocal.CoinFrom{}
	mainAssetFrom := big.NewInt(100000) //手续费
	if assetsChainId == sdk.chainId && assetsId == Default_asset_id {
		mainAssetFrom = mainAssetFrom.Add(mainAssetFrom, amount)
	} else {
		info, err := apiSdk.GetBalance(acc.GetAddr(), assetsChainId, assetsId)
		if err != nil {
			return "", err
		}
		if info.Balance.Cmp(amount) < 0 {
			return "", errors.New("Balance not emough!" + fmt.Sprintf("%d", assetsChainId) + "-" + fmt.Sprintf("%d", assetsId))
		}
		froms = append(froms, txprotocal.CoinFrom{
			Coin: txprotocal.Coin{
				Address:       acc.GetAddrBytes(),
				AssetsChainId: assetsChainId,
				AssetsId:      assetsId,
				Amount:        amount,
			},
			Nonce:  info.Nonce,
			Locked: 0,
		})
	}
	if accountInfo.Balance.Cmp(mainAssetFrom) < 0 {
		return "", errors.New("Balance not emough!" + fmt.Sprintf("%d", sdk.chainId) + "-" + fmt.Sprintf("%d", Default_asset_id))
	}
	froms = append(froms, txprotocal.CoinFrom{
		Coin: txprotocal.Coin{
			Address:       acc.GetAddrBytes(),
			AssetsChainId: sdk.chainId,
			AssetsId:      Default_asset_id,
			Amount:        mainAssetFrom,
		},
		Nonce:  accountInfo.Nonce,
		Locked: 0,
	})
	coinData := txprotocal.CoinData{
		Froms: froms,
		Tos: []txprotocal.CoinTo{{
			Coin: txprotocal.Coin{
				Address:       acc.GetAddrBytes(),
				AssetsChainId: assetsChainId,
				AssetsId:      assetsId,
				Amount:        amount,
			},
			LockValue: POCLockValue,
		}},
	}
	tx.CoinData, err = coinData.Serialize()
	if err != nil {
		return "", err
	}
	dType := uint8(0)
	tType := uint8(0)
	if timeType > 0 {
		dType = 1
		tType = timeType - 1
	}
	txData := txdata.Staking{
		Amount:        amount,
		Address:       acc.GetAddrBytes(),
		AssetsChainId: assetsChainId,
		AssetsId:      assetsId,
		DepositType:   dType,
		TimeType:      tType,
	}
	tx.Extend, err = txData.Serialize()
	if err != nil {
		return "", err
	}
	hash, _ := tx.GetHash().Serialize()
	signValue, _ := accSdk.Sign(acc, hash)
	txSign := txprotocal.CommonSignData{
		Signatures: []txprotocal.P2PHKSignature{{
			SignValue: signValue,
			PublicKey: acc.GetPubKey(),
		}},
	}
	tx.SignData, _ = txSign.Serialize()
	resultBytes, _ := tx.Serialize()
	txHex := hex.EncodeToString(resultBytes)
	bcdResult, err := apiSdk.Broadcast(txHex)
	if err != nil {
		return "", err
	}
	return bcdResult, nil
}
func (sdk NerveTxSDK) ExitStaking(prikeyHex, stakingTxHash string) (string, error) {
	//todo
	return "", nil
}
