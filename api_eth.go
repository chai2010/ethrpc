// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

import (
	"encoding/json"
	"fmt"
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

// 获取存储器数据
func EthGetStorageAt(host string, address string, pos int, block string) (data string, err error) {
	err = ethrpcCall(host, "eth_getStorageAt", &data, address, fmt.Sprintf("0x%x", pos), block)
	return
}

// 获取账户发出的交易数目
func EthGetTransactionCount(host string, address, block string) (n int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_getTransactionCount", &s, address, block)
	return parseInt64(s)
}

// 获取账户发出的交易数目
func EthGetBlockTransactionCountByHash(host string, hash string) (n int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_getBlockTransactionCountByHash", &s, hash)
	return parseInt64(s)
}

// 获取账户发出的交易数目
func EthGetBlockTransactionCountByNumber(host string, number int64) (n int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_getBlockTransactionCountByNumber", &s, fmt.Sprintf("0x%x", number))
	return parseInt64(s)
}

// 获取叔叔区块的数目
func EthGetUncleCountByBlockHash(host string, hash string) (n int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_getUncleCountByBlockHash", &s, hash)
	return parseInt64(s)

}

// 获取叔叔区块的数目
func EthGetUncleCountByBlockNumber(host string, number int64) (n int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_getUncleCountByBlockNumber", &s, fmt.Sprintf("0x%x", number))
	return parseInt64(s)
}

// 获取合约字节码
func EthGetCode(host string, block string) (code string, err error) {
	err = ethrpcCall(host, "eth_getCode", &code, block)
	return
}

// 对数据进行签名
// 注意: 签名的账户必须处于解锁状态
// sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))
func EthSign(host string, message string) (sign string, err error) {
	err = ethrpcCall(host, "eth_sign", &sign, message)
	return
}

// 发送交易数据
func EthSendTransaction(host string, tx *TxIn) (hash string, err error) {
	err = ethrpcCall(host, "eth_sendTransaction", &hash, tx)
	return
}

// 发送原生的交易数据
// 离线生成交易数据并签名后的数据
func EthSendRawTransaction(host string, rawtx string) (hash string, err error) {
	err = ethrpcCall(host, "eth_sendRawTransaction", &hash, rawtx)
	return
}

/*

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
