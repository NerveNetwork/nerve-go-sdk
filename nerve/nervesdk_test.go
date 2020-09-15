// @Title
// @Description
// @Author  Niels  2020/9/10
package nerve

//获取sdk工具对象
func ExampleGetSDK() {
	apiURL := "https://api.nerve.network/jsonrpc" //可以指定字节的节点（节点必须包含api模块）
	chainId := uint16(9)                          //Nerve主网为：9，测试网为：5，NULS主网为：1，NULS测试网为：2
	prefix := "NERVE"                             //Nerve主网为：NERVE，测试网为：TNVT，NULS主网为：NULS，NULS测试网为：tNULS

	sdk := GetSDK(apiURL, chainId, prefix)

	//use it
	sdk.GetApiUrl()
}
