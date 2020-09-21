// @Title
// @Description
// @Author  Niels  2020/9/15
package api

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/niels1286/nerve-go-sdk/utils/mathutils"
	"github.com/niels1286/nerve-go-sdk/utils/rpc"
	"math/big"
	"math/rand"
	"time"
)

type ApiSDK interface {
	Broadcast(txHex string) (string, error)
	ValidateTx(txHex string) (string, error)
}

type NerveApiSDK struct {
	client  *rpc.JsonRPCClient
	apiURL  string
	chainId uint16
	prefix  string
}

func GetApiSDK(apiURL string, chainId uint16, prefix string) ApiSDK {
	return &NerveApiSDK{
		client:  rpc.GetJsonRPCClient(apiURL),
		apiURL:  apiURL,
		chainId: chainId,
		prefix:  prefix,
	}
}

type AccountStatus struct {
	//账户地址
	Address       string
	assetsChainId uint16
	assetsId      uint16
	//可用余额
	Balance *big.Int
	//当前nonce值
	Nonce []byte
	//该nonce值的类型，1：已确认的nonce值,0：未确认的nonce值
	NonceType int
	//总的余额：可用+锁定
	TotalBalance *big.Int
}

func (sdk *NerveApiSDK) GetBalance(address string, chainId uint16, assetsId uint16) (*AccountStatus, error) {
	if address == "" {
		return nil, errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := sdk.client.NewRequestParam(rand.Intn(10000), "getAccountBalance", []interface{}{sdk.chainId, chainId, assetsId, address})
	result, err := sdk.ApiRequest(param)
	if err != nil {
		return nil, err
	}
	if nil == result || nil == result.Result {
		return nil, errors.New("Get nil result.")
	}
	resultMap := result.Result.(map[string]interface{})
	balance, err := mathutils.StringToBigInt(resultMap["balance"].(string))
	if err != nil {
		return nil, err
	}
	nonceHex := resultMap["nonce"].(string)
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		return nil, err
	}
	nonceType := resultMap["nonceType"].(float64)
	totalBalance, err := mathutils.StringToBigInt(resultMap["totalBalance"].(string))
	if err != nil {
		return nil, err
	}
	return &AccountStatus{
		Address:       address,
		assetsChainId: chainId,
		assetsId:      assetsId,
		Balance:       balance,
		Nonce:         nonce,
		NonceType:     int(nonceType),
		TotalBalance:  totalBalance,
	}, nil
}

//广播交易
func (sdk *NerveApiSDK) Broadcast(txHex string) (string, error) {
	if txHex == "" {
		return "", errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := sdk.client.NewRequestParam(rand.Intn(10000), "broadcastTx", []interface{}{sdk.chainId, txHex})
	result, err := sdk.ApiRequest(param)
	if err != nil {
		return "", err
	}
	if nil == result || nil == result.Result {
		if result != nil && result.Error != nil {
			bytes, _ := json.Marshal(result.Error)
			return "", errors.New(string(bytes))
		}
		return "", errors.New("Get nil result.")
	}
	resultMap := result.Result.(map[string]interface{})
	value := resultMap["value"].(bool)
	if !value {
		return "", errors.New("broadcast tx failed.")
	}
	hash := resultMap["hash"].(string)
	return hash, nil
}

func (sdk *NerveApiSDK) ValidateTx(txhex string) (string, error) {
	if txhex == "" {
		return "", errors.New("parameter wrong.")
	}
	rand.Seed(time.Now().Unix())
	param := sdk.client.NewRequestParam(rand.Intn(10000), "validateTx", []interface{}{sdk.chainId, txhex})
	result, err := sdk.ApiRequest(param)
	if err != nil {
		return "", err
	}
	if nil == result || nil == result.Result {
		if result != nil && result.Error != nil {
			bytes, _ := json.Marshal(result.Error)
			return "", errors.New(string(bytes))
		}
		return "", errors.New("Get nil result.")
	}
	resultMap := result.Result.(map[string]interface{})
	value := resultMap["value"].(string)
	return value, nil
}

//接口请求
//请求的地址是client中的默认地址
//本工具针对NULS的api模块的jsonrpc接口进行设计，适用范围有限
func (sdk *NerveApiSDK) ApiRequest(param *rpc.RequestParam) (*rpc.RequestResult, error) {
	return sdk.client.Request(param)
}
