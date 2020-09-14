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
// @Author  Niels  2020/3/27
package txprotocal

import (
	"encoding/hex"
	"github.com/niels1286/nerve-go-sdk/utils/seria"
	"reflect"
	"testing"
)

func TestSignData_Parse(t *testing.T) {
	pubHex := "029d7c84d95e039529613b1eb1e6cb6915aac136d122b9329d17578dc4cb7d1a7e"
	//priHex := "e894bf09ffda5dfd808b3cea3612ff494fdb1e621bda691462e5e1d15322183d"
	signValueHex := "3045022100a246c9fa3b8a49a06584718c8f0b01dc058778739285fd0efb953c7207711cb202200fa551b1ec58ce796414ed9b1a000bdfe38832a23e44222f8376b9d424d6a0bc"
	allHex := "21029d7c84d95e039529613b1eb1e6cb6915aac136d122b9329d17578dc4cb7d1a7e473045022100a246c9fa3b8a49a06584718c8f0b01dc058778739285fd0efb953c7207711cb202200fa551b1ec58ce796414ed9b1a000bdfe38832a23e44222f8376b9d424d6a0bc"
	pub, _ := hex.DecodeString(pubHex)
	sv, _ := hex.DecodeString(signValueHex)
	all, _ := hex.DecodeString(allHex)
	type args struct {
		reader *seria.ByteBufReader
	}
	tests := []struct {
		name    string
		s       CommonSignData
		args    args
		wantErr bool
		want    CommonSignData
	}{
		{name: "test sing parse", s: CommonSignData{}, args: args{reader: seria.NewByteBufReader(all, 0)}, wantErr: false, want: CommonSignData{[]P2PHKSignature{P2PHKSignature{
			SignValue: sv,
			PublicKey: pub,
		}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Parse(tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("sign parse failed.")
			}
		})
	}
}

func TestSignData_Serialize(t *testing.T) {
	pubHex := "029d7c84d95e039529613b1eb1e6cb6915aac136d122b9329d17578dc4cb7d1a7e"
	//priHex := "e894bf09ffda5dfd808b3cea3612ff494fdb1e621bda691462e5e1d15322183d"
	signValueHex := "3045022100a246c9fa3b8a49a06584718c8f0b01dc058778739285fd0efb953c7207711cb202200fa551b1ec58ce796414ed9b1a000bdfe38832a23e44222f8376b9d424d6a0bc"
	allHex := "21029d7c84d95e039529613b1eb1e6cb6915aac136d122b9329d17578dc4cb7d1a7e473045022100a246c9fa3b8a49a06584718c8f0b01dc058778739285fd0efb953c7207711cb202200fa551b1ec58ce796414ed9b1a000bdfe38832a23e44222f8376b9d424d6a0bc"
	pub, _ := hex.DecodeString(pubHex)
	sv, _ := hex.DecodeString(signValueHex)
	all, _ := hex.DecodeString(allHex)
	tests := []struct {
		name    string
		s       CommonSignData
		want    []byte
		wantErr bool
	}{
		{name: "sign serialize", s: CommonSignData{[]P2PHKSignature{P2PHKSignature{
			SignValue: sv,
			PublicKey: pub,
		}}}, want: all, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Serialize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//{
//"address" : "TNVTdTSPPjRFY13uapY9uYRG5MG5dvZwygyuF",
//"alias" : null,
//"pubkeyHex" : "02ee054d1a7d60ca779bc2cb4a848a5e1b1647e02e8542e5567d56f16869d015a4",
//"encryptedPrikeyHex" : "60aa7b423446bdbccb7d1ea92670d0659bbee94b7de2709023b8dc02bbe77bc8a260ccdae992ae05ab7e0b3c6969114e"
//}, {
//"address" : "TNVTdTSPV14WGUW8YNjcqxnaurKEmq79zbrgH",
//"alias" : null,
//"pubkeyHex" : "02a5e952d1a783f594ad056738ac00c5af69c1cd50ff72e46e8813d75ee8937067",
//"encryptedPrikeyHex" : "13da90aeb0b0b8a00ca3155fc0468a73a74b7314a90a1a076f2757ab55e5dbe39e009612ffcce8d65cc86bbb19a75e2d"
//}, {
//"address" : "TNVTdTSPW2a2RH8JKFRyZ2ciwwF2cpaV28HXB",
//"alias" : null,
//"pubkeyHex" : "026e36699151999100146de67e2c18250907a21866325d7bf62269abb76f5d8339",
//"encryptedPrikeyHex" : "d0340a9db6260d978e2cddfb7a87e23eae09181cf78585f39736898a58463908a2e631fd6a88f83e7147a9e1abff6981"
//}, {
//"address" : "TNVTdTSPFds4D1AK5NR9YhMRvtMC174QCFZT7",
//"alias" : null,
//"pubkeyHex" : "028bd6e2202831235cfa39600ef38a8b06f9e25b3fce0a331c22fc5fe4f5e1589a",
//"encryptedPrikeyHex" : "43f5715bb9165a2dfca3a25ec846764dfe0e19e6f2eea49a0c170c6e90fadd31f2ef4934ca79cf980d5cd08fee5a5abf"
//}

func TestMultiAddressesSignData_Parse(t *testing.T) {
	type fields struct {
		M              byte
		PubkeyList     [][]byte
		CommonSignData CommonSignData
	}
	type args struct {
		reader *seria.ByteBufReader
	}
	pub1, _ := hex.DecodeString("02401b78e28d293ad840f9298c2c7e522c68776e3badf092c2dbf457af1b8ed43e")
	pub2, _ := hex.DecodeString("023d994b0452216d13ae59b4544ea168f63360cf3a2ac1a2d74cfbf37cc6fa4848")
	pub3, _ := hex.DecodeString("024e2778bc64f5a3b3c8e99cf8fb420dae66822b9434612359d0ddbdd2c64af7db")
	bytes, _ := hex.DecodeString("")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test1", fields{
			M:              2,
			PubkeyList:     [][]byte{pub1, pub2, pub3},
			CommonSignData: CommonSignData{},
		}, args{reader: seria.NewByteBufReader(bytes, 0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MultiAddressesSignData{
				M:              tt.fields.M,
				PubkeyList:     tt.fields.PubkeyList,
				CommonSignData: tt.fields.CommonSignData,
			}

			if err := s.Parse(tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
