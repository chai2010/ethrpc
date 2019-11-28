// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

import "fmt"

// 网络编号
type NetworkId int

const (
	NetworkId_Nil NetworkId = iota
	NetworkId_Mainnet
	NetworkId_Morden
	NetworkId_Ropsten
	NetworkId_Rinkeby
	NetworkId_Kovan
)

func (v NetworkId) String() string {
	switch v {
	case NetworkId_Mainnet:
		return "Ethereum Mainnet"
	case NetworkId_Morden:
		return "Morden Testnet (deprecated)"
	case NetworkId_Ropsten:
		return "Ropsten Testnet"
	case NetworkId_Rinkeby:
		return "Rinkeby Testnet"
	case NetworkId_Kovan:
		return "Kovan Testnet"
	}
	return fmt.Sprintf("ethrpc.NetworkId: %d", int(v))
}
