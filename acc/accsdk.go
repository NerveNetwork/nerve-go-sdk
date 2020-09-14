// @Title
// @Description
// @Author  Niels  2020/9/10
package acc

import (
	"encoding/hex"
	"errors"
	"github.com/niels1286/nerve-go-sdk/crypto/base58"
	"github.com/niels1286/nerve-go-sdk/crypto/eckey"
	cryptoutils "github.com/niels1286/nerve-go-sdk/crypto/utils"
	"github.com/niels1286/nerve-go-sdk/utils/mathutils"
	"math/big"
	"regexp"
)

const (
	//默认的地址字节长度
	AddressBytesLength = 23
	//默认账户类型
	NormalAccountType = uint8(1)
	//合约账户类型
	ContractAccountType = uint8(2)
	//多签账户类型
	P2SHAccountType = uint8(3)
)

//用于前缀和实际地址的分隔符
var PrefixTable = [...]string{"", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

type AccountSDK interface {
	CreateAccount() (Account, error)
	ImportAccount(prikey []byte) (Account, error)
	ValidAddress(address string) error
	ImportFromKeyStore(keyStoreJson string, password string) (Account, error)
	CreateKeyStoreByPrikey(prikey []byte, password string) (*KeyStore, error)
	CreateKeyStore(account Account, password string) (*KeyStore, error)
	GetAddressByPubBytes(bytes []byte, accountType uint8) []byte
	GetAccountByEckey(ec *eckey.EcKey) (Account, error)
	GetStringAddress(bytes []byte) string
}

type NerveAccountSDK struct {
	chainId uint16
	prefix  string
}

func GetAccountSDK(chainId uint16, prefix string) AccountSDK {
	var sdk = &NerveAccountSDK{
		chainId: chainId,
		prefix:  prefix,
	}
	return sdk
}

func (sdk *NerveAccountSDK) CreateAccount() (Account, error) {
	ec, _ := eckey.NewEcKey()
	return sdk.GetAccountByEckey(ec)
}

func (sdk *NerveAccountSDK) ImportAccount(prikey []byte) (Account, error) {
	ec, err := eckey.FromPriKeyBytes(prikey)
	if err != nil {
		return nil, err
	}
	return sdk.GetAccountByEckey(ec)
}

//根据EcKey生成账户
func (sdk *NerveAccountSDK) GetAccountByEckey(ec *eckey.EcKey) (Account, error) {
	pubBytes := ec.GetPubKeyBytes(true)
	addressBytes := sdk.GetAddressByPubBytes(pubBytes, NormalAccountType)
	address := sdk.GetStringAddress(addressBytes)
	return &NerveAccount{
		addr:      address,
		prefix:    sdk.prefix,
		chainId:   sdk.chainId,
		prikeyHex: ec.GetPriKeyHex(),
		pubkeyHex: hex.EncodeToString(pubBytes),
		accType:   NormalAccountType,
		pubkey:    pubBytes,
		prikey:    ec.GetPriKeyBytes(),
		addrBytes: addressBytes,
	}, nil
}

//计算异或字节
func calcXor(bytes []byte) byte {
	xor := byte(0)
	for _, one := range bytes {
		xor ^= one
	}
	return xor
}

//根据地址字节数组，生成可以阅读的字符串地址
func (sdk *NerveAccountSDK) GetStringAddress(bytes []byte) string {
	//将之前得到的所有字节，进行异或操作，得到结果追加到
	xor := calcXor(bytes)
	newbytes := append(bytes, xor)
	return sdk.prefix + PrefixTable[len(sdk.prefix)] + base58.Encode(newbytes)
}

//根据公钥，生成账户地址
func (sdk *NerveAccountSDK) GetAddressByPubBytes(bytes []byte, accountType uint8) []byte {
	val := mathutils.BytesToBigInt(bytes)
	if val == nil || val.Cmp(big.NewInt(1)) <= 0 {
		return nil
	}
	hash160 := cryptoutils.Hash160(bytes)
	addressBytes := []byte{}
	addressBytes = append(addressBytes, mathutils.Uint16ToBytes(sdk.chainId)...)
	addressBytes = append(addressBytes, accountType)
	addressBytes = append(addressBytes, hash160...)
	return addressBytes
}

const message = "The address is not valid"

func (sdk *NerveAccountSDK) ValidAddress(address string) error {
	if address == "" {
		return errors.New(message)
	}
	prefix, realAddressStr, err := getRealAddress(address)
	if nil != err {
		return err
	}
	bytes := base58.Decode(realAddressStr)
	//长度必须正确，默认长度+一个校验位（xor）
	if len(bytes) != AddressBytesLength+1 {
		return errors.New(message)
	}
	chainId := mathutils.BytesToUint16(bytes[0:2])
	//验证已知链的前缀是否正确
	if chainId == sdk.chainId && prefix != sdk.prefix {
		return errors.New(message)
	}
	accountType := bytes[2]
	if accountType > P2SHAccountType {
		return errors.New(message)
	}
	addressBytes := bytes[0 : len(bytes)-1]
	xor := calcXor(addressBytes)
	if xor != bytes[len(bytes)-1] {
		//校验位不正确
		return errors.New(message)
	}
	return nil
}

//去除前缀，获得真正的地址字符串
func getRealAddress(address string) (string, string, error) {
	for index, c := range address {
		if index == 0 {
			continue
		}
		if c >= 97 {
			return address[0:index], address[index+1:], nil
		}
	}
	return "", "", errors.New(message)
}

//校验密码是否满足格式要求，如果不满足则返回false。
//密码至少8位，必须包含字母和数字
func (sdk *NerveAccountSDK) PasswordCheck(password string) bool {
	if password == "" {
		return false
	}
	length := len(password)
	if length < 8 || length > 20 {
		return false
	}
	reg, _ := regexp.Compile("(.*)[a-zA-Z](.*)")
	if !reg.MatchString(password) {
		return false
	}
	reg, _ = regexp.Compile("(.*)\\d+(.*)")
	if !reg.MatchString(password) {
		return false
	}
	reg, _ = regexp.Compile("(.*)\\s+(.*)")
	if reg.MatchString(password) {
		return false
	}
	reg, _ = regexp.Compile("(.*)[\u4e00-\u9fa5\u3000]+(.*)")
	if reg.MatchString(password) {
		return false
	}
	return true
}
