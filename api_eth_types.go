// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

// 同步状态
type SyncingStatus struct {
	IsSyncing     bool
	StartingBlock int64
	CurrentBlock  int64
	HighestBlock  int64
}
