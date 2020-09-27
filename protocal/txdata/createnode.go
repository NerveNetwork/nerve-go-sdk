// @Title
// @Description
// @Author  Niels  2020/9/25
package txdata

import (
	"github.com/niels1286/nerve-go-sdk/acc"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
	"math/big"
)

//创建节点的易扩展字段的具体协议
type CreateNode struct {
	//节点保证金金额
	Amount *big.Int
	//节点创建地址
	AgentAddress []byte
	//节点奖励接收地址
	RewardAddress []byte
	//节点打包地址
	PackingAddress []byte
}

//反序列化
func (a *CreateNode) Parse(reader *seria.ByteBufReader) error {
	var err error
	a.Amount, err = reader.ReadBigInt()
	if err != nil {
		return err
	}
	a.AgentAddress, err = reader.ReadBytes(acc.AddressBytesLength)
	if err != nil {
		return err
	}
	a.PackingAddress, err = reader.ReadBytes(acc.AddressBytesLength)
	if err != nil {
		return err
	}
	a.RewardAddress, err = reader.ReadBytes(acc.AddressBytesLength)
	if err != nil {
		return err
	}
	return nil
}

//序列化方法
func (a *CreateNode) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteBigint(a.Amount)
	writer.WriteBytes(a.AgentAddress)
	writer.WriteBytes(a.PackingAddress)
	writer.WriteBytes(a.RewardAddress)
	return writer.Serialize(), nil
}
