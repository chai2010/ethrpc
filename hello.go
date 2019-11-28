// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

// +build ignore

package main

import (
	"fmt"

	"github.com/chai2010/ethrpc"
)

func main() {
	const host = "https://ropsten.infura.io"

	// Geth/v1.9.7-omnibus-cc8122b7-20191122/linux-amd64/go1.13.4 <nil>
	version, err := ethrpc.Web3ClientVersion(host)
	fmt.Println(version, err)
	version, err = ethrpc.NewWeb3API(host).ClientVersion()
	fmt.Println(version, err)

	// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470 <nil>
	hash, err := ethrpc.Web3Sha3(host, []byte(""))
	fmt.Println(hash, err)
	hash, err = ethrpc.NewWeb3API(host).Sha3([]byte(""))
	fmt.Println(hash, err)

	// 网络ID
	netversion, err := ethrpc.NetVersion(host)
	fmt.Println(netversion, err)

	// 是否在监听
	listening, err := ethrpc.NetListening(host)
	fmt.Println(listening, err)

	// 连接的结点数目
	peercount, err := ethrpc.NetPeerCount(host)
	fmt.Println(peercount, err)
}
