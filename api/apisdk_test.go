// @Title
// @Description
// @Author  Niels  2020/9/29
package api

import (
	"fmt"
	"testing"
)

func TestNerveApiSDK_GetBalance(t *testing.T) {
	sdk := GetApiSDK("http://beta.api.nerve.network/jsonrpc/", 5, "TNVT")
	status, err := sdk.GetBalance("TNVTdTSPQvEngihwxqwCNPq3keQL1PwrcLbtj", 5, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(status)
}
