// @Title
// @Description
// @Author  Niels  2020/9/15
package txdata

import "github.com/niels1286/nerve-go-sdk/utils/seria"

//设置别名交易的拓展字段的详细协议
type Alias struct {
	//地址
	Address []byte
	//别名
	Alias string
}

//反序列化
func (a *Alias) Parse(reader *seria.ByteBufReader) error {
	var err error
	a.Address, err = reader.ReadBytesWithLen()
	if err != nil {
		return err
	}
	a.Alias, err = reader.ReadStringWithLen()
	if err != nil {
		return err
	}
	return nil
}

//序列化方法
func (a *Alias) Serialize() ([]byte, error) {
	writer := seria.NewByteBufWriter()
	writer.WriteBytesWithLen(a.Address)
	writer.WriteString(a.Alias)
	return writer.Serialize(), nil
}
