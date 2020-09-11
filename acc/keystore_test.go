// @Title
// @Description
// @Author  Niels  2020/9/11
package acc

import (
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KeyStore
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KeyStore
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdk := &NerveAccountSDK{
				chainId: tt.fields.chainId,
				prefix:  tt.fields.prefix,
			}
			got, err := sdk.ImportKeyStore(tt.args.keyStoreJson, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportKeyStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImportKeyStore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
