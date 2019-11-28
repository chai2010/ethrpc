// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

package ethrpc

import (
	"fmt"
	"regexp"
)

// 数据库接口
type DbAPI interface {
	PutString(dbname, key, value string) (bool, error)
	GetString(dbname, key string) (string, error)
	PutHex(dbname, key, value string) (bool, error)
	GetHex(dbname, key string) (string, error)
}

// 构造数据库接口
func NewDbAPI(host string) DbAPI {
	return &dbAPI{host}
}

// 保存字符串到本地数据库
func DbPutString(host string, dbname, key, value string) (bool, error) {
	var ok bool
	var err = ethrpcCall(host, "db_putString", &ok, dbname, key, value)
	return ok, err
}

// 从本地数据库读取字符串数据
func DbGetString(host string, dbname, key string) (string, error) {
	var s string
	var err = ethrpcCall(host, "db_getString", &s, dbname, key)
	return s, err
}

// 保存十六进制字符串到本地数据库
func DbPutHex(host string, dbname, key, value string) (bool, error) {
	if mached, _ := regexp.MatchString(`^0[xX][0-9a-fA-F]+$`, value); !mached {
		return false, fmt.Errorf("ethrpc.DbPutHex: invalid hex: %q", value)
	}

	var ok bool
	var err = ethrpcCall(host, "db_putHex", &ok, dbname, key, value)
	return ok, err
}

// 从本地数据库读取十六进制字符串数据
func DbGetHex(host string, dbname, key string) (string, error) {
	var s string
	var err = ethrpcCall(host, "db_getHex", &s, dbname, key)
	if err != nil {
		return "", err
	}
	if mached, _ := regexp.MatchString(`^0[xX][0-9a-fA-F]+$`, s); !mached {
		return s, fmt.Errorf("ethrpc.DbPutHex: invalid hex: %q", s)
	}
	return s, nil
}

type dbAPI struct{ host string }

func (p *dbAPI) PutString(dbname, key, value string) (bool, error) {
	return DbPutString(p.host, dbname, key, value)
}
func (p *dbAPI) GetString(dbname, key string) (string, error) {
	return DbGetString(p.host, dbname, key)

}
func (p *dbAPI) PutHex(dbname, key, value string) (bool, error) {
	return DbPutHex(p.host, dbname, key, value)

}
func (p *dbAPI) GetHex(dbname, key string) (string, error) {
	return DbGetHex(p.host, dbname, key)
}
