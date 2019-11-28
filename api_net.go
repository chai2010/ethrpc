// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

import (
	"fmt"
	"strconv"
)

// net接口
type NetAPI interface {
	Version() (NetworkId, error)
	Listening() (bool, error)
	PeerCount() (int, error)
}

// 构造Net接口对象
func NewNetAPI(host string) NetAPI {
	return &netAPI{host: host}
}

// 获取网络ID
func NetVersion(host string) (NetworkId, error) {
	var version string
	var err = ethrpcCall(host, "net_version", &version)
	if err != nil {
		return NetworkId_Nil, err
	}

	var id NetworkId
	if _, err := fmt.Sscan(version, &id); err != nil {
		return NetworkId_Nil, err
	}

	return id, nil
}

// 远程客户端是否在监听
func NetListening(host string) (bool, error) {
	var listening bool
	var err = ethrpcCall(host, "net_listening", &listening)
	return listening, err
}

// 返回连接的结点数目
func NetPeerCount(host string) (int, error) {
	var s string
	var err = ethrpcCall(host, "net_peerCount", &s)
	if err != nil {
		return 0, err
	}

	// 数字格式为十六进制 0xABCD
	if len(s) > len("0x") && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		i, _ := strconv.ParseInt(s[len("0x"):], 16, 64)
		return int(i), nil
	} else {
		return 0, fmt.Errorf("ethrpc.NetPeerCount: invalid result: %s", s)
	}
}

type netAPI struct{ host string }

func (p *netAPI) Version() (NetworkId, error) {
	return NetVersion(p.host)
}
func (p *netAPI) Listening() (bool, error) {
	return NetListening(p.host)
}
func (p *netAPI) PeerCount() (int, error) {
	return NetPeerCount(p.host)
}
