// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

import (
	"encoding/json"
	"math/big"
)

// 同步状态
type SyncingStatus struct {
	IsSyncing     bool
	StartingBlock int64
	CurrentBlock  int64
	HighestBlock  int64
}

// 输入的交易数据
type TxIn struct {
	From     string // 不参与签名
	To       string
	Gas      string
	GasPrice string
	Value    string
	Data     string
	Nonce    string // 可以缺省
}

// 查询返回的交易数据
type TxResult struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

// 交易回执
type TxReceipt struct {
	TransactionHash   string          `json:"transactionHash"`
	TransactionIndex  string          `json:"transactionIndex"`
	BlockHash         string          `json:"blockHash"`
	BlockNumber       string          `json:"blockNumber"`
	From              string          `json:"from"`
	To                string          `json:"to"`
	CumulativeGasUsed string          `json:"cumulativeGasUsed"`
	GasUsed           string          `json:"gasUsed"`
	ContractAddress   string          `json:"contractAddress"`
	Logs              json.RawMessage `json:"logs"`
	LogsBloom         string          `json:"logsBloom"`
	Root              string          `json:"root"`
	Status            string          `json:"status"`
}

type Block struct {
	Number           int64
	Hash             string
	ParentHash       string
	Nonce            int64
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	StateRoot        string
	ReceiptsRoot     string
	Miner            string
	Difficulty       *big.Int
	TotalDifficulty  *big.Int
	ExtraData        string
	Size             int64
	GasLimit         *big.Int
	GasUsed          *big.Int
	Timestamp        int64
	//Transactions     []Transaction
	Uncles []string
}
