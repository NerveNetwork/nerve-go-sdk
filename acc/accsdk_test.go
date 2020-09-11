// @Title
// @Description
// @Author  Niels  2020/9/10
package acc

import (
	"encoding/hex"
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
		{"NERVE-falsea", fields{chainId: 9, prefix: "NERVE"}, args{address: "NERVEepb6FhfgWLsdfsdfsdfedfuw3huvchBus"}, true},
		{"NERVE-falseb", fields{chainId: 9, prefix: "NERVE"}, args{address: "NERVEeacdFhfgWLsdfsdfsdfedfuw3huvchBus"}, true},
		{"NERVE-falsec", fields{chainId: 9, prefix: "NERVE"}, args{address: "GJbpb6T1UHeQPxjcjCA4fAUqa9Ce16sjv55"}, true},
		{"NERVE-falsed", fields{chainId: 9, prefix: "NERVE"}, args{address: "NERVEbpb6T1UHeQPxjcjCA4fAUqa9Ce16sjv55"}, false},
		{"NERVE-falsee", fields{chainId: 9, prefix: "NERVE"}, args{address: "SSSSSepb6FhfgWLsdfsdfsdfedfuw3huvchBus"}, true},
		{"NERVE-false2", fields{chainId: 9, prefix: "NERVE"}, args{address: "ASDFASDFASDFASDFASDF"}, true},
		{"NERVE-false3", fields{chainId: 9, prefix: "NERVE"}, args{address: ""}, true},
		{"error2", fields{chainId: 1, prefix: "NULS"}, args{"NULSd12KSBnno5mKghxcPyncnKn33F7ufGjucoVyaCBVmg5NaFCV"}, true},
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

func TestNerveAccountSDK_ImportAccount(t *testing.T) {
	type fields struct {
		chainId uint16
		prefix  string
	}
	type args struct {
		prikeyHex string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"test1", fields{chainId: 9, prefix: "NERVE"}, args{"00"}, "", true},
		{"test1", fields{chainId: 9, prefix: "NERVE"}, args{"01"}, "", true},

		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"602bec1904be78c646fe4f5c00f04cab6164be5ff80a48846b4afa0c96a76c25"}, "NULSd6HgeZEv2iqBKgpbnLAF8NFDav4LFeG55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"0e50110b30b24280622f5e8e8911d5e55efb89d476b8d9242448841857f0524c"}, "NULSd6HgcnQJUwQEBGPyRAs7xJyXkw3zEUL55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"fdf0b4bdd8bfb33c641e9b621f676bcc91c7fe541f407323080d35ab21a13668"}, "NULSd6Hgb8zsiaB7sHBMwoNQ5N5UswfihAZ55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"a338a27cb2d26dc7f41aef1eeada8b288ca60ebec5133cac0eb469fd9bdfcb6c"}, "NULSd6HghRNWRaKVuqGzGsMQeB3pih832md55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"bc8d923920316ab5040a68ee133f4ff9e6226b3bd09c8ca77df29d1e7571119a"}, "NULSd6Hge7e2Wem1wNbN6EBnUjMfGdZnwk855", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"c77fd5176ffa54ad4d89c9005c2861d8fd5574202d8a860718e6d41c9341fc18"}, "NULSd6HgiyYFLi7Qiof6ZWSNMiF38fjNVKS55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"45214fcbc081bb027c1c75e31c26d9027e6023edd8cddef93ad114acf04f1f7b"}, "NULSd6HggHSWi4GX96nQqNYUGc6ZV66kJKR55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"169cf30cb0027d1d5d6e64c1c19e1815fa29a4ed5c1e1321293bb987e3fde2be"}, "NULSd6HghFJUKCoZFMwTDTbE2vvQsi1WqVm55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"d565384b3d9b42217878f759bd0a45428bbcf5dd2726b4df0e2240d12e548de6"}, "NULSd6HgbT4iA2AGsz2K6bhftqtMF6ECfcn55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"3b295c7331182ecbc48be80ab8c45525b83b967128c3b1d9fa217bdd62b49f10"}, "NULSd6HgaAvw8LM4kJBfQX4JmuPMBMrshYF55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"eabb11e9539b5f6f8bd8e78bfdaa44c472b7fd72b176675154b4968c038f0459"}, "NULSd6HghrjR6iw8VRqsV4u5e3j8ZmUSKGB55", false},
		{"test1", fields{chainId: 1, prefix: "NULS"}, args{"a91f7db84306409dd5af9cac65c0ebe6f95749f293be6f4ded1ec6e8d49bdc06"}, "NULSd6HgZaUkenVAUa4JMYamxXqoxXBsyAL55", false},
		{"test1", fields{chainId: 9, prefix: "NERVE"}, args{"3146a7fb296a246cfa1c85d69cd0772cf12d465627b20b392226b0757dcc4e73"}, "NERVEepb63LZPU4ujwiq88DRXTgFcQFEqCDLAr", false},
		{"test1", fields{chainId: 9, prefix: "NERVE"}, args{"2e7cd068db708fd55a82e0e4375add5186045ee22d28fddf8c0799fa78f0b61c"}, "NERVEepb69KyDvjCgUPx7jaVwUrivAWtT7Ci1G", false},
		{"test1", fields{chainId: 9, prefix: "NERVE"}, args{"31894cfad61cb02576542a64a135cf6d9bdc0b562defdb7bc20b5006eb86c0f6"}, "NERVEepb61KCCUU4mDYDGtL6nEJ4a8zztAPsPZ", false},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"c72bd6b30bf5c4d1a8a1e3f5b14592c415d00797fc68efd62c6511cb03472f8e"}, "TNVTdTSPES8FAtc1nuodxCvx8TaPZtVqFFfhy", false},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"162d277cbda15afe257f3a544e809a2d320132f147bba1647071a6f655f47d97"}, "TNVTdTSPTsX19Ufbh8Vu3b84sp2r5NUiE2ToB", false},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"2b869bd9c5945a9e7aebfcd317c3c1fce2f5aed5d83b9e6b34c0da352ffa0794"}, "TNVTdTSPK1WAQTLGUrZWCJNprw2r8RSbNgfaK", false},
		{"test1", fields{chainId: 2, prefix: "tNULS"}, args{"9d8e2cb9f6ad367bcb65ccc5bbcd68c7a1112f53e2a19e45f2a03bd24f5ca025"}, "tNULSeBaMnq8QJ7Kp5ngQ44JyHZvfiaqG9uouc", false},
		{"test1", fields{chainId: 2, prefix: "tNULS"}, args{"620159afad81f8d95ec0e4824868e762699042b95037ffe353f708cd9affe9cb"}, "tNULSeBaMsKjbMn5vxacp4xkP2RC9BhN7TAtCm", false},
		{"test1", fields{chainId: 2, prefix: "tNULS"}, args{"f151fe8834c1383e8b9d439827320792a679adbf01fb8a765a9e2b80d5c3b5b9"}, "tNULSeBaMvNJqTzj4yTiZoGwq3yR94sLZRinG6", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &NerveAccountSDK{
				chainId: tt.fields.chainId,
				prefix:  tt.fields.prefix,
			}
			prikey, err := hex.DecodeString(tt.args.prikeyHex)
			got, err := sdk.ImportAccount(prikey)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if nil == got && err != nil {
				return
			}
			if got.GetAddr() != tt.want {
				t.Errorf("ImportAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNerveAccountSDK_GetAddressByPubBytes(t *testing.T) {
	type fields struct {
		chainId uint16
		prefix  string
	}
	type args struct {
		bytes       []byte
		accountType uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{"test 0&1", fields{
			chainId: 9,
			prefix:  "NERVE",
		}, args{
			bytes:       []byte{1},
			accountType: 1,
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &NerveAccountSDK{
				chainId: tt.fields.chainId,
				prefix:  tt.fields.prefix,
			}
			if got := sdk.GetAddressByPubBytes(tt.args.bytes, tt.args.accountType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddressByPubBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
