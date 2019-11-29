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
	if err != nil {
		return nil, err
	}
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
	if err != nil {
		return 0, err
	}
	return parseInt64(s)
}

// 获取账户发出的交易数目
func EthGetBlockTransactionCountByHash(host string, hash string) (n int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_getBlockTransactionCountByHash", &s, hash)
	if err != nil {
		return 0, err
	}
	return parseInt64(s)
}

// 获取账户发出的交易数目
func EthGetBlockTransactionCountByNumber(host string, number int64) (n int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_getBlockTransactionCountByNumber", &s, fmt.Sprintf("0x%x", number))
	if err != nil {
		return 0, err
	}
	return parseInt64(s)
}

// 获取叔叔区块的数目
func EthGetUncleCountByBlockHash(host string, hash string) (n int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_getUncleCountByBlockHash", &s, hash)
	if err != nil {
		return 0, err
	}
	return parseInt64(s)
}

// 获取叔叔区块的数目
func EthGetUncleCountByBlockNumber(host string, number int64) (n int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_getUncleCountByBlockNumber", &s, fmt.Sprintf("0x%x", number))
	if err != nil {
		return 0, err
	}
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

// 调用合约函数
func EthCall(host string, tx *TxIn, tag string) (result string, err error) {
	err = ethrpcCall(host, "eth_call", &result, tx, tag)
	return
}

// 估算执行合约函数需要的Gas数目
func EthEstimateGas(host string, tx *TxIn) (result int64, err error) {
	var s string
	err = ethrpcCall(host, "eth_estimateGas", &s, tx)
	if err != nil {
		return 0, err
	}
	return parseInt64(s)
}

// todo list

func eth_getBlockByHash(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getBlockByHash", &todo)
	return
}

func eth_getBlockByNumber(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getBlockByNumber", &todo)
	return
}

func eth_getTransactionByHash(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getTransactionByHash", &todo)
	return
}

func eth_getTransactionByBlockHashAndIndex(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getTransactionByBlockHashAndIndex", &todo)
	return
}

func eth_getTransactionByBlockNumberAndIndex(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getTransactionByBlockNumberAndIndex", &todo)
	return
}

func eth_getTransactionReceipt(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getTransactionReceipt", &todo)
	return
}

func eth_pendingTransactions(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_pendingTransactions", &todo)
	return
}

func eth_getUncleByBlockHashAndIndex(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getUncleByBlockHashAndIndex", &todo)
	return
}

func eth_getUncleByBlockNumberAndIndex(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getUncleByBlockNumberAndIndex", &todo)
	return
}

func eth_getCompilers(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getCompilers", &todo)
	return
}

func eth_compileLLL(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_compileLLL", &todo)
	return
}

func eth_compileSolidity(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_compileSolidity", &todo)
	return
}

func eth_compileSerpent(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_compileSerpent", &todo)
	return
}

func eth_newFilter(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_newFilter", &todo)
	return
}

func eth_newBlockFilter(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_newBlockFilter", &todo)
	return
}

func eth_newPendingTransactionFilter(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_newPendingTransactionFilter", &todo)
	return
}

func eth_uninstallFilter(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_uninstallFilter", &todo)
	return
}

func eth_getFilterChanges(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getFilterChanges", &todo)
	return
}

func eth_getFilterLogs(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getFilterLogs", &todo)
	return
}

func eth_getLogs(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getLogs", &todo)
	return
}

func eth_getWork(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getWork", &todo)
	return
}

func eth_submitWork(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_submitWork", &todo)
	return
}

func eth_submitHashrate(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_submitHashrate", &todo)
	return
}

func eth_getProof(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "eth_getProof", &todo)
	return
}
