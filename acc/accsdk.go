// @Title
// @Description
// @Author  Niels  2020/9/10
package acc

import (
	"encoding/hex"
	"github.com/niels1286/nerve-go-sdk/crypto/base58"
	"github.com/niels1286/nerve-go-sdk/crypto/eckey"
	cryptoutils "github.com/niels1286/nerve-go-sdk/crypto/utils"
	"github.com/niels1286/nerve-go-sdk/utils/mathutils"
	"math/big"
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

	ValidAddress(address string) error

	GetAddressByPubBytes(bytes []byte, accountType uint8) []byte

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
	ec, err := eckey.NewEcKey()
	if err != nil {
		return nil, err
	}
	return sdk.getAccountByEckey(ec, sdk.chainId, sdk.prefix)
}

//根据EcKey生成账户
func (sdk *NerveAccountSDK) getAccountByEckey(ec *eckey.EcKey, chainId uint16, prefix string) (Account, error) {
	pubBytes := ec.GetPubKeyBytes(true)
	addressBytes := sdk.GetAddressByPubBytes(pubBytes, NormalAccountType)
	address := sdk.GetStringAddress(addressBytes)
	return &NerveAccount{
		addr:      address,
		prefix:    prefix,
		chainId:   chainId,
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

func (sdk *NerveAccountSDK) ValidAddress(address string) error {
	//todo
	return nil
}
