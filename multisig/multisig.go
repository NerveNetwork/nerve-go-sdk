// @Title
// @Description
// @Author  Niels  2020/9/14
package multisig

import (
	"encoding/hex"
	"errors"
	"github.com/niels1286/nerve-go-sdk/acc"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
	"sort"
	"strings"
)

type MultiAccount struct {
	M            int
	Pks          string
	Address      string
	AddressBytes []byte
}
type MultiAccountSDK interface {
	CreateMultiAccount(m int, pubkeysHex string) (*MultiAccount, error)
}

type NerveMultiAccountSDK struct {
	chainId uint16
	prefix  string
}

func GetAccountSDK(chainId uint16, prefix string) MultiAccountSDK {
	var sdk = &NerveMultiAccountSDK{
		chainId: chainId,
		prefix:  prefix,
	}
	return sdk
}

func (sdk *NerveMultiAccountSDK) CreateMultiAccount(m int, pubkeysHex string) (*MultiAccount, error) {
	if m < 2 || m > 15 {
		return nil, errors.New("m value valid")
	}
	pubkeysHex = strings.TrimSpace(pubkeysHex)
	pkHexSlice := strings.Split(pubkeysHex, ",")
	if len(pkHexSlice) < m || len(pkHexSlice) > 15 {
		return nil, errors.New("Incorrect public keys")
	}

	sort.Slice(pkHexSlice, func(i, j int) bool {
		return pkHexSlice[i] < pkHexSlice[j]
	})
	writer := seria.NewByteBufWriter()
	writer.WriteByte(byte(sdk.chainId))
	pks := ""
	writer.WriteByte(uint8(m))
	for i, pk := range pkHexSlice {
		bytes, err := hex.DecodeString(pk)
		if err != nil {
			return nil, err
		}
		writer.WriteBytes(bytes)
		if i == 0 {
			pks = pk
		} else {
			pks += "," + pk
		}
	}
	bytes := writer.Serialize()
	accSDK := acc.GetAccountSDK(sdk.chainId, sdk.prefix)
	addressBytes := accSDK.GetAddressByPubBytes(bytes, acc.P2SHAccountType)
	val := &MultiAccount{
		M:            m,
		Pks:          pks,
		Address:      accSDK.GetStringAddress(addressBytes),
		AddressBytes: addressBytes,
	}
	return val, nil
}
