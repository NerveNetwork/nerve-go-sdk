// @Title
// @Description
// @Author  Niels  2020/9/15
package rpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonRPCClient struct {
	client *http.Client
	url    string
}

func GetJsonRPCClient(url string) *JsonRPCClient {
	return &JsonRPCClient{
		client: &http.Client{},
		url:    url,
	}
}

//请求参数封装
type RequestParam struct {
	//固定值：2.0
	Jsonrpc string `json:"jsonrpc"`
	//接口名称
	Method string `json:"method"`
	//可以是slice，也可以是结构体
	Params interface{} `json:"params"`
	//请求的唯一标识，返回的结果中也包含该id
	Id int `json:"id"`
}

func (pp *RequestParam) ToJson() string {
	data, err := json.Marshal(pp)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return string(data)
}

func (pp *RequestParam) ToJsonBytes() []byte {
	data, err := json.Marshal(pp)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return data
}

//请求返回结构体
type RequestResult struct {
	//固定值：2.0
	Jsonrpc string `json:"jsonrpc"`
	//对应请求的唯一标识
	Id string `json:"id"`

	Result interface{} `json:"result"`

	Error interface{} `json:"error"`
}

//组装请求参数
//@id,请求的唯一标识，在返回的结构体中，也会包含本字段，并且等于请求时的值
//@method,请求的具体方法
//@params,实际参数，请参照接口文档进行组装
func (client *JsonRPCClient) NewRequestParam(id int, method string, params interface{}) *RequestParam {
	return &RequestParam{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      id,
	}
}

func (client *JsonRPCClient) Request(param *RequestParam) (*RequestResult, error) {
	req, err := http.NewRequest("POST", client.url, bytes.NewReader(param.ToJsonBytes()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &RequestResult{}
	json.Unmarshal(body, result)
	return result, nil
}
