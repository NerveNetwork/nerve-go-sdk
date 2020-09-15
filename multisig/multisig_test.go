// @Title
// @Description
// @Author  Niels  2020/9/14
package multisig

import (
	"reflect"
	"testing"
)

func TestNerveMultiAccountSDK_CreateMultiAccount(t *testing.T) {
	type fields struct {
		chainId uint16
		prefix  string
	}
	type args struct {
		m          int
		pubkeysHex string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *MultiAccount
		wantErr bool
	}{
		{"test1", fields{5, "TNVT"},
			args{2, "02401b78e28d293ad840f9298c2c7e522c68776e3badf092c2dbf457af1b8ed43e,023d994b0452216d13ae59b4544ea168f63360cf3a2ac1a2d74cfbf37cc6fa4848,024e2778bc64f5a3b3c8e99cf8fb420dae66822b9434612359d0ddbdd2c64af7db"},
			&MultiAccount{
				M:       2,
				Pks:     "023d994b0452216d13ae59b4544ea168f63360cf3a2ac1a2d74cfbf37cc6fa4848,02401b78e28d293ad840f9298c2c7e522c68776e3badf092c2dbf457af1b8ed43e,024e2778bc64f5a3b3c8e99cf8fb420dae66822b9434612359d0ddbdd2c64af7db",
				Address: "TNVTdTSPrbqXUBnC9C7JLH4HysLLNCD6q48XB",
			}, false},
		{"test2", fields{5, "TNVT"},
			args{1, "02401b78e28d293ad840f9298c2c7e522c68776e3badf092c2dbf457af1b8ed43e,023d994b0452216d13ae59b4544ea168f63360cf3a2ac1a2d74cfbf37cc6fa4848,024e2778bc64f5a3b3c8e99cf8fb420dae66822b9434612359d0ddbdd2c64af7db"},
			nil, true},
		{"test3", fields{5, "TNVT"},
			args{18, "02401b78e28d293ad840f9298c2c7e522c68776e3badf092c2dbf457af1b8ed43e,023d994b0452216d13ae59b4544ea168f63360cf3a2ac1a2d74cfbf37cc6fa4848,024e2778bc64f5a3b3c8e99cf8fb420dae66822b9434612359d0ddbdd2c64af7db"},
			nil, true},
		{"test3", fields{5, "TNVT"},
			args{5, "02401b78e28d293ad840f9298c2c7e522c68776e3badf092c2dbf457af1b8ed43e,023d994b0452216d13ae59b4544ea168f63360cf3a2ac1a2d74cfbf37cc6fa4848,024e2778bc64f5a3b3c8e99cf8fb420dae66822b9434612359d0ddbdd2c64af7db"},
			nil, true},
		{"test4", fields{5, "TNVT"},
			args{2, "zxcfdsdf78e28d293ad840f9298c2c7e522c68776e3badf092c2dbf457af1b8ed43e,023d994b0452216d13ae59b4544ea168f63360cf3a2ac1a2d74cfbf37cc6fa4848,024e2778bc64f5a3b3c8e99cf8fb420dae66822b9434612359d0ddbdd2c64af7db"},
			nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := GetAccountSDK(tt.fields.chainId, tt.fields.prefix)
			got, err := sdk.CreateMultiAccount(tt.args.m, tt.args.pubkeysHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMultiAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMultiAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
