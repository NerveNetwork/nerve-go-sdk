// @Title
// @Description
// @Author  Niels  2020/9/15
package api

import (
	"encoding/json"
	"errors"
	"github.com/niels1286/nerve-go-sdk/utils/rpc"
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
