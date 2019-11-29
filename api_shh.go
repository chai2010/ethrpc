// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

// todo ist

func shh_post(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_post", &todo)
	return
}

func shh_version(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_version", &todo)
	return
}

func shh_newIdentity(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_newIdentity", &todo)
	return
}

func shh_hasIdentity(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_hasIdentity", &todo)
	return
}

func shh_newGroup(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_newGroup", &todo)
	return
}

func shh_addToGroup(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_addToGroup", &todo)
	return
}

func shh_newFilter(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_newFilter", &todo)
	return
}

func shh_uninstallFilter(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_uninstallFilter", &todo)
	return
}

func shh_getFilterChanges(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_getFilterChanges", &todo)
	return
}

func shh_getMessages(host string) (err error) {
	var todo string
	err = ethrpcCall(host, "shh_getMessages", &todo)
	return
}
