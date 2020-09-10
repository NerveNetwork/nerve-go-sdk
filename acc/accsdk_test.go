// @Title
// @Description
// @Author  Niels  2020/9/10
package acc

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestNerveAccountSDK_CreateAccount(t *testing.T) {
	type fields struct {
		chainId uint16
		prefix  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"NULS", fields{
			chainId: 1,
			prefix:  "NULS",
		}, "NULSd", false},
		{"tNULS", fields{
			chainId: 2,
			prefix:  "tNULS",
		}, "tNULSe", false},
		{"NERVE", fields{
			chainId: 1,
			prefix:  "NERVE",
		}, "NERVEe", false},
		{"TNVT", fields{
			chainId: 1,
			prefix:  "TNVT",
		}, "TNVTd", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := GetAccountSDK(tt.fields.chainId, tt.fields.prefix)
			got, err := sdk.CreateAccount()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.HasPrefix(got.GetAddr(), tt.want) {
				t.Errorf("CreateAccount() got = %v, want %v", got, tt.want)
			}
			if got.GetPriKeyHex() == "" {
				t.Errorf("CreateAccount() got  GetPriKeyHex is nil")
			}
			if got.GetPubKeyHex() == "" {
				t.Errorf("CreateAccount() got  GetPubKeyHex is nil")
			}
			if got.GetPrefix() == "" {
				t.Errorf("CreateAccount() got  GetPrefix is nil")
			}
			if got.GetChainId() != tt.fields.chainId {
				t.Errorf("CreateAccount() got  GetChainId is nil")
			}
			if got.GetPubKey() == nil {
				t.Errorf("CreateAccount() got  GetPubKey is nil")
			}
			if got.GetPriKey() == nil {
				t.Errorf("CreateAccount() got  GetPriKey is nil")
			}
			if got.GetAddr() != sdk.GetStringAddress(got.GetAddrBytes()) {
				t.Errorf("CreateAccount()  GetAddrBytes is wrong")
			}
			if got.GetType() != NormalAccountType {
				t.Errorf("CreateAccount() got GetType is wrong")
			}
			if !reflect.DeepEqual(sdk.GetAddressByPubBytes(got.GetPubKey(), NormalAccountType), got.GetAddrBytes()) {
				t.Errorf("CreateAccount() got is wrong")
			}

			log.Println(got.GetAddr())
		})
	}
}

func TestNerveAccountSDK_ValidAddress(t *testing.T) {
	type fields struct {
		chainId uint16
		prefix  string
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := GetAccountSDK(tt.fields.chainId, tt.fields.prefix)
			if err := sdk.ValidAddress(tt.args.address); (err != nil) != tt.wantErr {
				t.Errorf("ValidAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_calcXor(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"test1", args{bytes: []byte{1, 0, 1, 1, 1, 1, 1}}, 0},
		{"test2", args{bytes: []byte{1, 0, 1, 1, 1, 6, 1}}, 7},
		{"test3", args{bytes: []byte{1, 0, 1, 3, 1, 1, 1}}, 2},
		{"test4", args{bytes: []byte{1, 0, 120, 100, 69, 41, 21}}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcXor(tt.args.bytes); got != tt.want {
				t.Errorf("calcXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
