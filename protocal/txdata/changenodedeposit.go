// @Title
// @Description
// @Author  Niels  2020/9/25
package txdata

import (
	"github.com/niels1286/nerve-go-sdk/acc"
	txprotocal "github.com/niels1286/nerve-go-sdk/protocal"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
	"math/big"
)

//停止节点的交易扩展字段的具体协议
type ChangeNodeDeposit struct {
	Address []byte
	Amount  *big.Int
	//节点的hash
	NodeHash *txprotocal.NulsHash
}

//反序列化
func (s *ChangeNodeDeposit) Parse(reader *seria.ByteBufReader) error {
	var errs error
	s.Address, errs = reader.ReadBytes(acc.AddressBytesLength)
	if errs != nil {
		return errs
	}

	s.Amount, errs = reader.ReadBigInt()
	if errs != nil {
		return errs
	}

	bytes, err := reader.ReadBytes(txprotocal.HashLength)
	if err != nil {
		return err
	}
	s.NodeHash = txprotocal.NewNulsHash(bytes)
	return nil
}

//序列化方法
func (s *ChangeNodeDeposit) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()

	writer.WriteBytes(s.Address)
	writer.WriteBigint(s.Amount)

	hash, _ := s.NodeHash.Serialize()
	writer.WriteBytes(hash)
	return writer.Serialize(), nil
}
