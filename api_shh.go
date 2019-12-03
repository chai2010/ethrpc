// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

func shh_version(host string) (version string, err error) {
	err = ethrpcCall(host, "shh_version", &version)
	return
}

func shh_post(host string, msg *ShhPostObject) (ok bool, err error) {
	err = ethrpcCall(host, "shh_post", &ok, msg)
	return
}

func shh_newIdentity(host string) (id string, err error) {
	err = ethrpcCall(host, "shh_newIdentity", &id)
	return
}

func shh_hasIdentity(host, id string) (exists bool, err error) {
	err = ethrpcCall(host, "shh_hasIdentity", &exists, id)
	return
}

func shh_newGroup(host string) (id string, err error) {
	err = ethrpcCall(host, "shh_newGroup", &id)
	return
}

func shh_addToGroup(host, id string) (ok bool, err error) {
	err = ethrpcCall(host, "shh_addToGroup", &ok, id)
	return
}

// todo ist

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
