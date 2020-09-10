// @Title
// @Description
// @Author  Niels  2020/9/10
package acc

import (
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

type AccountSDK interface {
	CreateAccount() (*Account, error)

	ValidAddress(address string) error
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

func (sdk *NerveAccountSDK) CreateAccount() (*Account, error) {
	ec, err := eckey.NewEcKey()
	if err != nil {
		return nil, err
	}
	return getAccountByEckey(ec, sdk.chainId, sdk.prefix)
}

//根据EcKey生成账户
func getAccountByEckey(ec *eckey.EcKey, chainId uint16, prefix string) (*Account, error) {
	pubBytes := ec.GetPubKeyBytes(true)
	addressBytes := GetAddressByPubBytes(pubBytes, chainId, NormalAccountType)
	address := GetStringAddress(addressBytes, prefix)
	return &Account{
		Address:      address,
		AddressBytes: addressBytes,
		ChainId:      chainId,
		AccType:      NormalAccountType,
		EcKey:        ec,
		Prefix:       prefix,
	}, nil
}

//根据公钥，生成账户地址
func GetAddressByPubBytes(bytes []byte, chainId uint16, accountType uint8) []byte {

	val := mathutils.BytesToBigInt(bytes)
	if val == nil || val.Cmp(big.NewInt(1)) <= 0 {
		return nil
	}

	hash160 := cryptoutils.Hash160(bytes)
	addressBytes := []byte{}
	addressBytes = append(addressBytes, mathutils.Uint16ToBytes(chainId)...)
	addressBytes = append(addressBytes, accountType)
	addressBytes = append(addressBytes, hash160...)
	return addressBytes
}

func (sdk *NerveAccountSDK) ValidAddress(address string) error {
	//todo
	return nil
}
