// @Title
// @Description
// @Author  Niels  2020/9/29
package api

import (
	"fmt"
	"testing"
)

func TestNervePSSDK_GetNode(t *testing.T) {
	sdk := GetPSSDK(5, "TNVT")
	val, _ := sdk.GetNode("http://beta.public.nerve.network", "e07d6195b1f757a06fcb040e29f75e5a03149fc677d88f941e4eb724da82bae8")
	fmt.Println(val.Amount)
}
