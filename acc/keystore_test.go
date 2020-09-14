// @Title
// @Description
// @Author  Niels  2020/9/11
package acc

import (
	"encoding/hex"
	cryptoutils "github.com/niels1286/nerve-go-sdk/crypto/utils"
	"reflect"
	"testing"
)

func TestNerveAccountSDK_CreateKeyStore(t *testing.T) {
	type fields struct {
		chainId uint16
		prefix  string
	}
	type args struct {
		account  Account
		password string
	}
	prikey1, _ := hex.DecodeString("17cbdc7e6b5924176d4616344b3150b5b8b17e7e8483428753cb5f8264e69d94")
	prikey2, _ := hex.DecodeString("1234567890")
	account1, _ := GetAccountSDK(5, "TNVT").ImportAccount(prikey1)
	account2, _ := GetAccountSDK(5, "TNVT").ImportAccount(prikey2)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KeyStore
		wantErr bool
	}{
		{"test1", fields{chainId: 5, prefix: "TNVT"}, args{
			account:  account1,
			password: "nuls123456",
		}, &KeyStore{
			Address:             "TNVTdTSPSQeYbiBQn8jjxBYPPn6FHNxx8GPF4",
			EncryptedPrivateKey: "1e95bb898ea65196942b4fc3b4d4f882dcd1895f7dc5787b5eddfd59e058d13a7e0994f7aa9eae79728a58e0ecd991dd",
			Pubkey:              "03acde2549553daea8b6f6e792fc18198625440e05f8c240e91e3f145346bfad28",
			version:             1,
		}, false},
		{"test2", fields{chainId: 5, prefix: "TNVT"}, args{
			account:  account2,
			password: "nuls123456",
		}, &KeyStore{
			Address:             account2.GetAddr(),
			EncryptedPrivateKey: hex.EncodeToString(cryptoutils.AESEncrypt(account2.GetPriKey(), cryptoutils.Sha256h([]byte("nuls123456")))),
			Pubkey:              account2.GetPubKeyHex(),
			version:             1,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &NerveAccountSDK{
				chainId: tt.fields.chainId,
				prefix:  tt.fields.prefix,
			}
			got, err := sdk.CreateKeyStore(tt.args.account, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateKeyStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateKeyStore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNerveAccountSDK_CreateKeyStoreByPrikey(t *testing.T) {
	type fields struct {
		chainId uint16
		prefix  string
	}
	type args struct {
		prikey   []byte
		password string
	}
	prikey1, _ := hex.DecodeString("17cbdc7e6b5924176d4616344b3150b5b8b17e7e8483428753cb5f8264e69d94")
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KeyStore
		wantErr bool
	}{
		{"test1", fields{
			chainId: 5,
			prefix:  "TNVT",
		}, args{
			prikey:   prikey1,
			password: "nuls123456",
		}, &KeyStore{
			Address:             "TNVTdTSPSQeYbiBQn8jjxBYPPn6FHNxx8GPF4",
			EncryptedPrivateKey: "1e95bb898ea65196942b4fc3b4d4f882dcd1895f7dc5787b5eddfd59e058d13a7e0994f7aa9eae79728a58e0ecd991dd",
			Pubkey:              "03acde2549553daea8b6f6e792fc18198625440e05f8c240e91e3f145346bfad28",
			version:             1,
		}, false},
		{"test1", fields{
			chainId: 5,
			prefix:  "TNVT",
		}, args{
			prikey:   prikey1,
			password: "qwe",
		}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &NerveAccountSDK{
				chainId: tt.fields.chainId,
				prefix:  tt.fields.prefix,
			}
			got, err := sdk.CreateKeyStoreByPrikey(tt.args.prikey, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateKeyStoreByPrikey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateKeyStoreByPrikey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNerveAccountSDK_ImportKeyStore(t *testing.T) {
	type fields struct {
		chainId uint16
		prefix  string
	}
	type args struct {
		keyStoreJson string
		password     string
	}
	prikey1, _ := hex.DecodeString("17cbdc7e6b5924176d4616344b3150b5b8b17e7e8483428753cb5f8264e69d94")
	account1, _ := GetAccountSDK(5, "TNVT").ImportAccount(prikey1)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		{"test1", fields{
			chainId: 5,
			prefix:  "TNVT",
		}, args{
			keyStoreJson: "{\"address\":\"TNVTdTSPSQeYbiBQn8jjxBYPPn6FHNxx8GPF4\",\"encryptedPrivateKey\":\"1e95bb898ea65196942b4fc3b4d4f882dcd1895f7dc5787b5eddfd59e058d13a7e0994f7aa9eae79728a58e0ecd991dd\",\"pubKey\":\"03acde2549553daea8b6f6e792fc18198625440e05f8c240e91e3f145346bfad28\",\"priKey\":null}",
			password:     "nuls123456",
		}, account1, false},
		{"test2", fields{
			chainId: 5,
			prefix:  "TNVT",
		}, args{
			keyStoreJson: "{\"address\":\"\",\"encryptedPrivateKey\":\"1e95bb898ea65196942b4fc3b4d4f882dcd1895f7dc5787b5eddfd59e058d13a7e0994f7aa9eae79728a58e0ecd991dd\",\"pubKey\":\"03acde2549553daea8b6f6e792fc18198625440e05f8c240e91e3f145346bfad28\",\"priKey\":null}",
			password:     "nuls123456",
		}, account1, true},
		{"test3", fields{
			chainId: 5,
			prefix:  "TNVT",
		}, args{
			keyStoreJson: "{\"address\":\"\",\"pubKey\":\"03acde2549553daea8b6f6e792fc18198625440e05f8c240e91e3f145346bfad28\",\"priKey\":null}",
			password:     "nuls123456",
		}, nil, true},
		{"test4", fields{
			chainId: 5,
			prefix:  "TNVT",
		}, args{
			keyStoreJson: "{\"address\":\"\",\"pubKey\":\"03acde2549553daea8b6f6e792fc18198625440e05f8c240e91e3f145346bfad28\",\"priKey\":null",
			password:     "nuls123456123123",
		}, nil, true},
		{"test5", fields{
			chainId: 5,
			prefix:  "TNVT",
		}, args{
			keyStoreJson: "{\"address\":\"\",\"encryptedPrivateKey\":\"1e95bb898eahjkyuzxc65196942b4fc3b4d4f882dcd1895f7dc5787b5eddfd59e058d13a7e0994f7aa9eae79728a58e0ecd991dd\",\"pubKey\":\"03acde2549553daea8b6f6e792fc18198625440e05f8c240e91e3f145346bfad28\",\"priKey\":null}",
			password:     "nuls123456",
		}, nil, true},
		{"test6", fields{
			chainId: 5,
			prefix:  "TNVT",
		}, args{
			keyStoreJson: "{\"address\":\"\",\"encryptedPrivateKey\":\"1e95bb898ea65196942b4fc3b4d4f882dcd1895f7dc5787b5eddfd59e058d13a7e0994f7aa9eae79728a58e0ecd991dd\",\"pubKey\":\"03acde2549553daea8b6f6e792fc18198625440e05f8c240e91e3f145346bfad28\",\"priKey\":null}",
			password:     "qwe",
		}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &NerveAccountSDK{
				chainId: tt.fields.chainId,
				prefix:  tt.fields.prefix,
			}
			got, err := sdk.ImportFromKeyStore(tt.args.keyStoreJson, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportFromKeyStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImportFromKeyStore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
