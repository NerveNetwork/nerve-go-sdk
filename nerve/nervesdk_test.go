// @Title
// @Description
// @Author  Niels  2020/9/10
package nerve

import "testing"

func TestNerveSDK_GetApiUrl(t *testing.T) {

	tests := []struct {
		name string
		want string
	}{
		{"normal", "https://api.nerve.network/jsonrpc"},
		{"normal", "test"},
		{"normal", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := GetSDK(tt.want, 9, "NERVE")
			if got := sdk.GetApiUrl(); got != tt.want {
				t.Errorf("GetApiUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
