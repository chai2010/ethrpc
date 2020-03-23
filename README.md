- *赞助 BTC: 1Cbd6oGAUUyBi7X7MaR4np4nTmQZXVgkCW*
- *赞助 ETH: 0x623A3C3a72186A6336C79b18Ac1eD36e1c71A8a6*

----

# 以太坊山寨RPC客户端

官方以太坊客户端太重，其中更是间接依赖了CGO特性。因此山寨一个轻量版以太坊RPC客户端。

## 例子

```go
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
```

## 版权

版权 @2019 柴树杉
