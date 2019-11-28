// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

// +build ignore

package main

import (
	"fmt"

	"github.com/chai2010/ethrpc"
)

func main() {
	// Geth/v1.9.7-omnibus-cc8122b7-20191122/linux-amd64/go1.13.4 <nil>
	version, err := ethrpc.Web3ClientVersion("https://ropsten.infura.io")
	fmt.Println(version, err)
	version, err = ethrpc.NewWeb3API("https://ropsten.infura.io").ClientVersion()
	fmt.Println(version, err)

	// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470 <nil>
	hash, err := ethrpc.Web3Sha3("https://ropsten.infura.io", []byte(""))
	fmt.Println(hash, err)
	hash, err = ethrpc.NewWeb3API("https://ropsten.infura.io").Sha3([]byte(""))
	fmt.Println(hash, err)
}
