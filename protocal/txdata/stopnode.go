// @Title
// @Description
// @Author  Niels  2020/9/25
package txdata

import (
	txprotocal "github.com/niels1286/nerve-go-sdk/protocal"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
)

//停止节点的交易扩展字段的具体协议
type StopNode struct {
	Address []byte
	//节点的hash
	AgentHash *txprotocal.NulsHash
}

//反序列化
func (s *StopNode) Parse(reader *seria.ByteBufReader) error {
	var errs error
	s.Address, errs = reader.ReadBytesWithLen()
	if errs != nil {
		return errs
	}

	bytes, err := reader.ReadBytes(txprotocal.HashLength)
	if err != nil {
		return err
	}
	s.AgentHash = txprotocal.NewNulsHash(bytes)
	return nil
}

//序列化方法
func (s *StopNode) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()

	writer.WriteBytesWithLen(s.Address)

	hash, _ := s.AgentHash.Serialize()
	writer.WriteBytes(hash)
	return writer.Serialize(), nil
}
