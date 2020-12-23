// @Title
// @Description
// @Author  Niels  2020/9/10
package nerve

import (
	"github.com/niels1286/nerve-go-sdk/acc"
	"github.com/niels1286/nerve-go-sdk/api"
	"github.com/niels1286/nerve-go-sdk/multisig"
	"github.com/niels1286/nerve-go-sdk/txs"
)

/**
 * 基本工具，封装了所有接口的调用
 */
type NerveSDK struct {
	apiUrl  string
	chainId uint16
	prefix  string
	acc.AccountSDK
	multisig.MultiAccountSDK
	api.ApiSDK
	api.PSSDK
	txs.TxSDK
}

//
func GetSDK(apiUrl string, chainId uint16, addressPrefix string) *NerveSDK {
	var sdk = &NerveSDK{
		apiUrl:  apiUrl,
		chainId: chainId,
		prefix:  addressPrefix,
	}
	sdk.AccountSDK = acc.GetAccountSDK(chainId, addressPrefix)
	sdk.MultiAccountSDK = multisig.GetAccountSDK(chainId, addressPrefix)
	sdk.ApiSDK = api.GetApiSDK(apiUrl, chainId, addressPrefix)
	sdk.PSSDK = api.GetPSSDK(chainId, addressPrefix)
	sdk.TxSDK = txs.GetTxSDK(apiUrl, chainId, addressPrefix)
	return sdk
}

//Get the set API service access path
func (sdk *NerveSDK) GetApiUrl() string {
	return sdk.apiUrl
}
