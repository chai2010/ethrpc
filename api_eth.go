// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

import (
	"encoding/json"
	"math/big"
	"strings"
)

// 获取协议版本
func EthProtocolVersion(host string) (version string, err error) {
	err = ethrpcCall(host, "eth_protocolVersion", &version)
	return
}

// 获取同步状态
func EthSyncing(host string) (syncing *SyncingStatus, err error) {
	var s string
	err = ethrpcCall(host, "eth_syncing", &s)
	if err != nil {
		return nil, err
	}

	// 没有同步
	if strings.EqualFold(s, "false") {
		return new(SyncingStatus), nil
	}

	// 解析同步状态对象
	var x struct {
		StartingBlock string
		CurrentBlock  string
		HighestBlock  string
	}
	err = json.Unmarshal([]byte(s), &x)
	if err != nil {
		return nil, err
	}

	// 填充状态
	syncing = &SyncingStatus{
		IsSyncing:     true,
		StartingBlock: asInt64(x.StartingBlock),
		CurrentBlock:  asInt64(x.CurrentBlock),
		HighestBlock:  asInt64(x.HighestBlock),
	}
	return
}

// 获取CoinBase地址
func EthCoinbase(host string) (address string, err error) {
	err = ethrpcCall(host, "eth_coinbase", &address)
	return
}

// 是否在挖矿
func EthMining(host string) (mining bool, err error) {
	err = ethrpcCall(host, "eth_mining", &mining)
	return
}

// 矿工产生hash的速率
func EthHashrate(host string) (rate int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_hashrate", &s)
	if err != nil {
		return 0, err
	}
	return parseInt64(s)
}

// 查询gas的价格
func EthGasPrice(host string) (gasPrice int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_gasPrice", &s)
	if err != nil {
		return 0, err
	}
	return parseInt64(s)
}

// 获取结点的账户地址列表
func EthAccounts(host string) (accounts []string, err error) {
	err = ethrpcCall(host, "eth_accounts", &accounts)
	return
}

// 获取最新的区块数目
func EthBlockNumber(host string) (blockNumber int64, err error) {
	err = ethrpcCall(host, "eth_blockNumber", &blockNumber)
	return
}

// 查询账户余额
func EthGetBalance(host string, address, block string) (balance *big.Int, err error) {
	var s string
	err = ethrpcCall(host, "eth_getBalance", &s, address, block)
	return parseBigint(s)
}

/*

eth_getStorageAt
eth_getTransactionCount
eth_getBlockTransactionCountByHash
eth_getBlockTransactionCountByNumber
eth_getUncleCountByBlockHash
eth_getUncleCountByBlockNumber
eth_getCode
eth_sign
eth_sendTransaction
eth_sendRawTransaction
eth_call
eth_estimateGas
eth_getBlockByHash
eth_getBlockByNumber
eth_getTransactionByHash
eth_getTransactionByBlockHashAndIndex
eth_getTransactionByBlockNumberAndIndex
eth_getTransactionReceipt
eth_pendingTransactions
eth_getUncleByBlockHashAndIndex
eth_getUncleByBlockNumberAndIndex
eth_getCompilers
eth_compileLLL
eth_compileSolidity
eth_compileSerpent
eth_newFilter
eth_newBlockFilter
eth_newPendingTransactionFilter
eth_uninstallFilter
eth_getFilterChanges
eth_getFilterLogs
eth_getLogs
eth_getWork
eth_submitWork
eth_submitHashrate
eth_getProof
*/
