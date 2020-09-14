/*
 *  MIT License
 *  Copyright (c) 2019-2020 niels.wang
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 *
 */

// @Title
// @Description
// @Author  Niels  2020/3/25
package txprotocal

import (
	"encoding/hex"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
)

//hash的字节长度，默认为32
const HashLength = 32

//交易hash和区块hash的结构体
type NulsHash struct {
	//实际字节数据
	bytes []byte
	//缓存的字符串数据
	hashHex string
}

//创建一个新的hash对象
func NewNulsHash(bytes []byte) *NulsHash {
	return &NulsHash{bytes: bytes, hashHex: ""}
}

//创建一个新的hash对象
func ImportNulsHash(hashHex string) *NulsHash {
	bytes, _ := hex.DecodeString(hashHex)
	return NewNulsHash(bytes)
}

//序列化hash字节数组，长度为32位
func (hash *NulsHash) Serialize() ([]byte, error) {
	return hash.bytes, nil
}

//从reader中读取32个字节，赋值到hash中
func (hash *NulsHash) Parse(reader *seria.ByteBufReader) error {
	bytes, err := reader.ReadBytes(HashLength)
	hash.bytes = bytes
	return err
}

func (hash *NulsHash) String() string {
	if len(hash.hashHex) == 0 {
		hash.hashHex = hex.EncodeToString(hash.bytes)
	}
	return hash.hashHex
}
