// @Title
// @Description
// @Author  Niels  2020/9/10
package acc

import (
	"encoding/hex"
	"fmt"
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

func TestNerveAccountSDK_PasswordCheck(t *testing.T) {
	type fields struct {
		chainId uint16
		prefix  string
	}
	type args struct {
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"qwer1234"}, true},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"qwe"}, false},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"qweqweqwe"}, false},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"12345678"}, false},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{""}, false},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"!@#$%^&*()"}, false},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"nuls123456"}, true},
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{"nuls123!@#$"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &NerveAccountSDK{
				chainId: tt.fields.chainId,
				prefix:  tt.fields.prefix,
			}
			if got := sdk.PasswordCheck(tt.args.password); got != tt.want {
				t.Errorf("PasswordCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

//创建账户，支持NULS生态内任意区块链的账户创建
func ExampleNerveAccountSDK_CreateAccount() {
	chainId := uint16(9) //Nerve主网为：9，测试网为：5，NULS主网为：1，NULS测试网为：2
	prefix := "NERVE"    //Nerve主网为：NERVE，测试网为：TNVT，NULS主网为：NULS，NULS测试网为：tNULS

	sdk := GetAccountSDK(chainId, prefix)
	//创建账户
	account, err := sdk.CreateAccount()
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Address:", account.GetAddr())
	fmt.Println("Private Key:", account.GetPriKeyHex())
	fmt.Println("Public Key:", account.GetPubKeyHex())
	fmt.Println("Chain ID:", account.GetChainId())
}

//使用私钥导入账户，支持NULS生态内任意区块链的账户创建
func ExampleNerveAccountSDK_ImportAccount() {
	chainId := uint16(9) //Nerve主网为：9，测试网为：5，NULS主网为：1，NULS测试网为：2
	prefix := "NERVE"    //Nerve主网为：NERVE，测试网为：TNVT，NULS主网为：NULS，NULS测试网为：tNULS

	sdk := GetAccountSDK(chainId, prefix)
	//创建账户
	prikey, _ := hex.DecodeString("602bec1904be78c646fe4f5c00f04cab6164be5ff80a48846b4afa0c96a76c25")
	account, err := sdk.ImportAccount(prikey)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Address:", account.GetAddr())
	fmt.Println("Private Key:", account.GetPriKeyHex())
	fmt.Println("Public Key:", account.GetPubKeyHex())
	fmt.Println("Chain ID:", account.GetChainId())
}

func TestNerveAccountSDK_GetBytesAddress(t *testing.T) {
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
		want    string
		wantErr bool
	}{
		{"test-1", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb67PB9xWGjpw1wzgoV9W11CZgsuee15"}, "09000169c97d30b4fdbdd40770b64336ef8bbdff87d545", false},
		{"test-2", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb6GTLWpJvzBG2qmQayVBQivgTTBU6E3"}, "090001fd50bf11292f6cd4dda1d9a2d086bd2600ba6762", false},
		{"test-3", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb6DLRrjJs78hu87sRsaxDoqyyEs3BMi"}, "090001ca9763378cbcb94740029be3a8cc631e60bffa60", false},
		{"test-4", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb69wSvKfYvVbqQ8LbHzuwrbJMf9h9cY"}, "090001935c44c8802ab4cff4c776f0c40868c27acfe839", false},
		{"test-5", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb69ufDrd6ZyXkoaef245TRBYVDmXrVV"}, "09000192dbef97aa568017a95f725c47ba162dc80887c3", false},
		{"test-6", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb67mKfhRAC9dPqHY8fHUYhPW7tYRZvy"}, "0900016fff28a6320ed738c68a72cab898b101121918f8", false},
		{"test-7", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb68aveD5dMBWHq4kvU2TrfuPcQHYcJK"}, "0900017d580c36e399d747dbc72faba1225319c88aac92", false},
		{"test-8", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb6C4cQ777E3afTPyoA6TdMDgDpbwVHq"}, "090001b5e4cc45ae0e2e24d61db3190455f579583ae175", false},
		{"test-9", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb65U7Nyg28BXzur59Ctg27AzGFxQLaH"}, "0900014aa55d4abdfdc789222828de132c927f2e934e1d", false},
		{"test-0", fields{chainId: 9, prefix: "NERVE"}, args{"NERVEepb61mKgUDpt8ELK9h2gZ4i1Bd9H5Wxz7"}, "0900010e6c22d6a520601b3855ef9814a60c9c3ddadd16", false},
		{"test-5-1", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPRYuXYDgWZVMLDwr6bRWzwjBk5mwAH"}, "050001b5c59b0ee084a8c604e67da1b4148b5342ecd9e8", false},
		{"test-5-2", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPJYNowzMYVjVnhbMTHqVK6hwi4EPNu"}, "05000143c960ffcd75b318f0b6a3e9fdcccaa537b52296", false},
		{"test-5-3", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPRzFex3N1vJXyKEi9Fp6iARphNMFYi"}, "050001bce0fc81aa29a0e03f1788acd0669897b44ac43e", false},
		{"test-5-4", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPP5Nk4ouLGfMq2hQJje5dgi9CS7DaJ"}, "0500018d872215cb7d52f897eb1a859145abd6618ca73e", false},
		{"test-5-5", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPPbT9Sqpu9bcy5SYUFEKU1w2NFTNZf"}, "05000195f5f29de657ee60573b2461d2315b00bc035beb", false},
		{"test-5-6", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPMLAwFpT59qgEe6YyAgfUhJw3prx9V"}, "0500017126db812887ff027d2f8c8d9cf41ef15a7db991", false},
		{"test-5-7", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPRG85jT4r3JGnGgst2XMN9xgvEx4sD"}, "050001b110e72af1d9951fe89c0a92fca34b9c1d49aac4", false},
		{"test-5-8", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPLbnCkPv4DzRPDYUu2TkpxhiqVA1Ev"}, "0500016544032a201f775d9143a5f1ce6dd8907e6f1df7", false},
		{"test-5-9", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPNtTnhe8JYrxJyMjNpM6Rm32WW7UDG"}, "0500018a77cf53b1a1846f6ee64569917b2364fa81346b", false},
		{"test-5-0", fields{chainId: 5, prefix: "TNVT"}, args{"TNVTdTSPQ3xKrPtQiqSEbWt6vDrNTMnnxxNQQ"}, "0500019d644f2a381a7eeafd6bacf4f36a984c79e754ce", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &NerveAccountSDK{
				chainId: tt.fields.chainId,
				prefix:  tt.fields.prefix,
			}
			got, err := sdk.GetBytesAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBytesAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != hex.EncodeToString(got) {
				t.Errorf("GetBytesAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
