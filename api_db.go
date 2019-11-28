// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

// 数据库接口
type DbAPI interface {
	PutString(host string, dbname, key, value string) (bool, error)
	GetString(host string, dbname, key string) (string, error)
	PutHex(host string, dbname, key, value string) (bool, error)
	GetHex(host string, dbname, key string) (string, error)
}

// 构造数据库接口
func NewDbAPI(host string) DbAPI {
	return nil
}

// 保存字符串到本地数据库
func DbPutString(host string, dbname, key, value string) (bool, error) {
	// db_putString
	panic("todo")
}

// 从本地数据库读取字符串数据
func DbGetString(host string, dbname, key string) (string, error) {
	// db_getString
	panic("todo")
}

// 保存十六进制字符串到本地数据库
func DbPutHex(host string, dbname, key, value string) (bool, error) {
	// ^0[xX][0-9a-fA-F]+$
	// db_putHex
	panic("todo")
}

// 从本地数据库读取十六进制字符串数据
func DbGetHex(host string, dbname, key string) (string, error) {
	// db_getHex
	panic("todo")
}
