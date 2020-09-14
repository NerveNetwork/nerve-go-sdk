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
 */

//@Description 这里列举了所有NULS网络支持的交易类型
package txprotocal

const (
	//coinbase 共识奖励交易类型
	TX_TYPE_COIN_BASE = 1

	//the type of the transfer transaction
	TX_TYPE_TRANSFER = 2

	//设置账户别名
	//*Set the transaction type of account alias.
	TX_TYPE_ACCOUNT_ALIAS = 3

	//新建共识节点
	TX_TYPE_REGISTER_AGENT = 4

	//委托参与共识
	TX_TYPE_DEPOSIT = 5

	//取消委托
	TX_TYPE_CANCEL_DEPOSIT = 6

	//黄牌惩罚交易
	TX_TYPE_YELLOW_PUNISH = 7

	//红牌惩罚交易
	TX_TYPE_RED_PUNISH = 8

	//注销共识节点交易类型
	TX_TYPE_STOP_AGENT = 9

	//跨链转账交易类型
	TX_TYPE_CROSS_CHAIN = 10

	//注册平行链交易类型
	TX_TYPE_REGISTER_CHAIN_AND_ASSET = 11

	//从NULS网络中注销一条平行链的交易的类型
	TX_TYPE_DESTROY_CHAIN_AND_ASSET = 12

	//为平行链登记一种资产的交易类型
	TX_TYPE_ADD_ASSET_TO_CHAIN = 13

	//删除链上资产交易类型
	TX_TYPE_REMOVE_ASSET_FROM_CHAIN = 14

	//创建智能合约交易的类型
	TX_TYPE_CREATE_CONTRACT = 15

	//调用智能合约的交易的类型
	TX_TYPE_CALL_CONTRACT = 16

	//删除智能合约交易的类型
	TX_TYPE_DELETE_CONTRACT = 17

	//合约内部转账交易的类型
	//contract transfer tx type
	TX_TYPE_CONTRACT_TRANSFER = 18

	//合约执行手续费返还交易的类型
	//合约在调用时，会模拟执行一次，用来估算需要花费的手续费金额，实际在调用过程中，给出了更多的手续费（出于一定要执行成功的目的），
	//当合约实际执行时，可能会出现手续费富余，这部分gas将以“合约手续费返回交易”的方式返回给调用者
	TX_TYPE_CONTRACT_RETURN_GAS = 19

	//合约创建共识节点交易
	TX_TYPE_CONTRACT_CREATE_AGENT = 20

	//合约委托交易
	TX_TYPE_CONTRACT_DEPOSIT = 21

	//合约撤销委托交易
	TX_TYPE_CONTRACT_CANCEL_DEPOSIT = 22

	//合约停止节点交易
	TX_TYPE_CONTRACT_STOP_AGENT = 23

	//跨链验证人变更交易
	TX_TYPE_VERIFIER_CHANGE = 24

	//跨链验证人初始化交易
	TX_TYPE_VERIFIER_INIT = 25

	//token跨链转账
	CONTRACT_TOKEN_CROSS_TRANSFER = 26
	//账本注册本来新增资产
	LEDGER_ASSET_REG_TRANSFER = 27

	//追加节点保证金
	APPEND_AGENT_DEPOSIT = 28

	// 撤销节点保证金
	// Cancel agent deposit

	REDUCE_AGENT_DEPOSIT = 29

	// 喂价交易
	QUOTATION = 30

	// 最终喂价交易
	FINAL_QUOTATION = 31

	// 批量退出staking交易
	BATCH_WITHDRAW = 32

	// 合并活期staking记录
	BATCH_STAKING_MERGE = 33

	// 创建交易对
	COIN_TRADING = 228

	// 挂单委托
	TRADING_ORDER = 229

	// 挂单撤销
	TRADING_ORDER_CANCEL = 230

	// 挂单成交
	TRADING_DEAL = 231

	// 修改交易对
	EDIT_COIN_TRADING = 232

	// 撤单交易确认
	ORDER_CANCEL_CONFIRM = 233

	// 确认 虚拟银行变更交易
	CONFIRM_CHANGE_VIRTUAL_BANK = 40

	// 虚拟银行变更交易
	CHANGE_VIRTUAL_BANK = 41

	// 链内充值交易
	RECHARGE = 42

	// 提现交易
	WITHDRAWAL = 43

	// 确认提现成功状态交易
	CONFIRM_WITHDRAWAL = 44

	// 发起提案交易
	PROPOSAL = 45

	// 对提案进行投票交易
	VOTE_PROPOSAL = 46

	// 异构链交易手续费补贴
	DISTRIBUTION_FEE = 47

	// 虚拟银行初始化异构链
	INITIALIZE_HETEROGENEOUS = 48

	// 异构链合约资产注册等待
	HETEROGENEOUS_CONTRACT_ASSET_REG_PENDING = 49

	// 异构链合约资产注册完成
	HETEROGENEOUS_CONTRACT_ASSET_REG_COMPLETE = 50

	// 确认提案执行交易
	CONFIRM_PROPOSAL = 51

	// 重置异构链(合约)虚拟银行
	RESET_HETEROGENEOUS_VIRTUAL_BANK = 52

	// 确认重置异构链(合约)虚拟银行
	CONFIRM_HETEROGENEOUS_RESET_VIRTUAL_BANK = 53

	//修改跨链注册信息
	REGISTERED_CHAIN_CHANGE = 60
)
