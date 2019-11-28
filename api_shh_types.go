// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

type ShhPostObject struct {
	From     string
	To       string
	Topics   []string
	Payload  string
	Priority string
	Ttl      string
}

type ShhChangedMsg struct {
	Hash       string
	From       string
	To         string
	Expiry     string
	Sent       string
	Ttl        string
	Topics     []string
	Payload    string
	WorkProved string
}
