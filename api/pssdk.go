// @Title
// @Description
// @Author  Niels  2020/9/29
package api

import (
	"encoding/json"
	"errors"
	"github.com/niels1286/nerve-go-sdk/utils/rpc"
	"math/big"
	"math/rand"
)

type PSSDK interface {
	GetNode(psApiURL string, nodeHash string) (*NodeInfo, error)
}

type NervePSSDK struct {
	chainId uint16
	prefix  string
}

func GetPSSDK(chainId uint16, prefix string) PSSDK {
	return &NervePSSDK{
		chainId: chainId,
		prefix:  prefix,
	}
}

type NodeInfo struct {
	NodeHash       string
	Amount         *big.Int
	NodeAddress    string
	RewardAddress  string
	PackingAddress string
	CreditVal      float64
	NodeId         string
	NodeAlias      string
	IsBank         bool
}

func (sdk *NervePSSDK) GetNode(psApiURL string, nodeHash string) (*NodeInfo, error) {
	client := rpc.GetJsonRPCClient(psApiURL)
	param := client.NewRequestParam(rand.Intn(10000), "getConsensusNode", []interface{}{sdk.chainId, nodeHash})
	result, err := client.Request(param)
	if err != nil {
		return nil, err
	}
	if nil == result || nil == result.Result {
		if result != nil && result.Error != nil {
			bytes, _ := json.Marshal(result.Error)
			return nil, errors.New(string(bytes))
		}
		return nil, errors.New("Get nil result.")
	}
	nodeInfo := result.Result.(map[string]interface{})

	return &NodeInfo{
		NodeHash:       nodeHash,
		Amount:         big.NewInt(int64(nodeInfo["deposit"].(float64))),
		NodeAddress:    nodeInfo["agentAddress"].(string),
		RewardAddress:  nodeInfo["rewardAddress"].(string),
		PackingAddress: nodeInfo["packingAddress"].(string),
		CreditVal:      nodeInfo["creditValue"].(float64),
		NodeId:         nodeInfo["agentId"].(string),
		NodeAlias:      nodeInfo["agentAlias"].(string),
		IsBank:         nodeInfo["bankNode"].(bool),
	}, nil
}
