// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

import "fmt"

// Web3接口
type Web3API interface {
	ClientVersion() (version string, err error)
	Sha3(data []byte) (hash string, err error)
}

// 构造Web3接口对象
func NewWeb3API(host string) Web3API {
	return &web3API{host: host}
}

// 返回远程客户端的版本信息
func Web3ClientVersion(host string) (version string, err error) {
	err = jsonrpcCall(host, "web3_clientVersion", &version)
	return
}

// 返回Keccak-256哈希
func Web3Sha3(host string, data []byte) (hash string, err error) {
	err = jsonrpcCall(host, "web3_sha3", &hash, fmt.Sprintf("0x%x", data))
	return
}

type web3API struct{ host string }

func (p *web3API) ClientVersion() (version string, err error) {
	return Web3ClientVersion(p.host)
}
func (p *web3API) Sha3(data []byte) (hash string, err error) {
	return Web3Sha3(p.host, data)
}
