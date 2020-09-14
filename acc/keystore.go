/*
 * MIT License
 * Copyright (c) 2019-2020 niels.wang
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package acc

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	cryptoutils "github.com/niels1286/nerve-go-sdk/crypto/utils"
)

//keystore 结构，所有NULS体系的keystore都遵循该基本协议
//todo 下一个版本增加加密算法的说明和算法参数
type KeyStore struct {
	//地址
	Address string `json:"address"`
	//使用密码对私钥进行加密后，得到的密文
	EncryptedPrivateKey string `json:"encryptedPrivateKey"`
	//公钥Hex
	Pubkey string `json:"pubkey"`
	//版本号
	version int
}

func (sdk *NerveAccountSDK) ImportFromKeyStore(keyStoreJson string, password string) (account Account, err error) {
	defer func() {
		if err := recover(); err != nil {
			account = nil
		}
	}()
	store := KeyStore{}
	err = json.Unmarshal([]byte(keyStoreJson), &store)
	if err != nil {
		return nil, err
	}
	if store.EncryptedPrivateKey == "" {
		return nil, errors.New("Keystore is broken.")
	}
	data, err := hex.DecodeString(store.EncryptedPrivateKey)
	if err != nil {
		return nil, err
	}
	err = errors.New("password may be wrong!")
	pwd := cryptoutils.Sha256h([]byte(password))
	prikey := cryptoutils.AESDecrypt(data, pwd)
	account, err = sdk.ImportAccount(prikey)
	if nil != err {
		return account, err
	}
	if store.Address != account.GetAddr() {
		return account, errors.New("Got a different address!")
	}
	return account, nil
}

func (sdk *NerveAccountSDK) CreateKeyStoreByPrikey(prikey []byte, password string) (*KeyStore, error) {
	account, err := sdk.ImportAccount(prikey)
	if err != nil {
		return nil, err
	}
	return sdk.CreateKeyStore(account, password)
}

func (sdk *NerveAccountSDK) CreateKeyStore(account Account, password string) (*KeyStore, error) {
	if !sdk.PasswordCheck(password) {
		return nil, errors.New("Invalid password format")
	}
	key := cryptoutils.Sha256h([]byte(password))
	epk := cryptoutils.AESEncrypt(account.GetPriKey(), key)
	epkHex := hex.EncodeToString(epk)
	return &KeyStore{Address: account.GetAddr(), EncryptedPrivateKey: epkHex, Pubkey: account.GetPubKeyHex(), version: 1}, nil
}
