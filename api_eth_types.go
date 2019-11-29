// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

import (
	"encoding/json"
	"fmt"
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
	From     string
	To       string
	Gas      int64
	GasPrice *big.Int
	Value    *big.Int
	Data     string
	Nonce    int64
}

// 编码为JSON
func (p *TxIn) MarshalJSON() ([]byte, error) {
	var x = struct {
		From     string `json:"from"`
		To       string
		Gas      string
		GasPrice string
		Value    string
		Data     string
		Nonce    string
	}{
		From:     p.From,
		To:       p.To,
		Gas:      fmt.Sprintf("0x%x", p.Gas),
		GasPrice: fmt.Sprintf("0x%x", p.GasPrice),
		Value:    fmt.Sprintf("0x%x", p.Value),
		Data:     p.Data,
		Nonce:    fmt.Sprintf("0x%x", p.Nonce),
	}

	return json.Marshal(x)
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
