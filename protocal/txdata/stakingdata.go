// @Title
// @Description
// @Author  Niels  2020/9/25
package txdata

import (
	"github.com/niels1286/nerve-go-sdk/acc"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
	"math/big"
)

type Staking struct {
	//委托金额
	Amount *big.Int
	//委托者的地址
	Address       []byte
	AssetsChainId uint16
	AssetsId      uint16
	DepositType   byte
	TimeType      byte
}

//反序列化
func (d *Staking) Parse(reader *seria.ByteBufReader) error {
	var err error
	d.Amount, err = reader.ReadBigInt()
	if err != nil {
		return err
	}
	d.Address, err = reader.ReadBytes(acc.AddressBytesLength)
	if err != nil {
		return err
	}
	d.AssetsChainId, err = reader.ReadUint16()
	if err != nil {
		return err
	}
	d.AssetsId, err = reader.ReadUint16()
	if err != nil {
		return err
	}
	d.DepositType, err = reader.ReadByte()
	if err != nil {
		return err
	}
	d.TimeType, err = reader.ReadByte()
	if err != nil {
		return err
	}
	return nil
}

//序列化方法
func (d *Staking) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteBigint(d.Amount)
	writer.WriteBytes(d.Address)
	writer.WriteUInt16(d.AssetsChainId)
	writer.WriteUInt16(d.AssetsId)
	writer.WriteByte(d.DepositType)
	writer.WriteByte(d.TimeType)
	return writer.Serialize(), nil
}
