// @Title
// @Description
// @Author  Niels  2020/9/10
package acc

import (
	"github.com/niels1286/nerve-go-sdk/crypto/eckey"
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
			ecKey, _ := eckey.FromPriKeyBytes(got.GetPriKey())
			newWant, _ := sdk.GetAccountByEckey(ecKey)
			if !reflect.DeepEqual(got, newWant) {
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
		{"NULS", fields{chainId: 1, prefix: "NULS"}, args{address: "NULSd6HgcCrxrHjARG6Uhhu1kjEpXWqysfD7J"}, false},
		{"NULS-false", fields{chainId: 1, prefix: "NULS"}, args{address: "NULSd6HgcCrxrHjARG6UhwerkjEpXWqysfD7J"}, true},
		{"NERVE", fields{chainId: 9, prefix: "NERVE"}, args{address: "NERVEepb6FhfgWLyqHUQgHoHwRGuw3huvchBus"}, false},
		{"NERVE-false", fields{chainId: 9, prefix: "NERVE"}, args{address: "NERVEepb6FhfgWLsdfsdfsdfedfuw3huvchBus"}, true},
		{"NERVE-false2", fields{chainId: 9, prefix: "NERVE"}, args{address: "ASDFASDFASDFASDFASDF"}, true},
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

func Test_getRealAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{"NULS", args{"NULSd6HgeTTm2pssKHakh29fNW4t98DDZ4KVR"}, "NULS", "6HgeTTm2pssKHakh29fNW4t98DDZ4KVR", false},
		{"Nerve", args{"NERVEepb6Chtj1NEaxu8VC5LqojAoxknX4RExF"}, "NERVE", "pb6Chtj1NEaxu8VC5LqojAoxknX4RExF", false},
		{"tNerve", args{"TNVTdTSPVtySemdyntG2UHCbTGWyHycn2aHET"}, "TNVT", "TSPVtySemdyntG2UHCbTGWyHycn2aHET", false},
		{"error", args{"ASDFASDFASDFASDFSADFASDF"}, "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getRealAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRealAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRealAddress() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRealAddress() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
