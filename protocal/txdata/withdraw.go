// @Title
// @Description
// @Author  Niels  2020/9/25
package txdata

import (
	txprotocal "github.com/niels1286/nerve-go-sdk/protocal"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
)

//退出委托的交易扩展字段的具体协议
type Withdraw struct {
	Address []byte
	//委托交易的hash
	StakingTxHash *txprotocal.NulsHash
}

//反序列化
func (w *Withdraw) Parse(reader *seria.ByteBufReader) error {
	var err error
	w.Address, err = reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	bytes, err := reader.ReadBytes(txprotocal.HashLength)
	if err != nil {
		return err
	}
	w.StakingTxHash = txprotocal.NewNulsHash(bytes)
	return nil
}

//序列化方法
func (w *Withdraw) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteBytesWithLen(w.Address)
	hash, _ := w.StakingTxHash.Serialize()
	writer.WriteBytes(hash)
	return writer.Serialize(), nil
}
