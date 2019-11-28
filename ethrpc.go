// 以太坊山寨RPC客户端 版权 @2019 柴树杉。

// 以太坊山寨RPC客户端
//
// 详细API请参考 https://github.com/ethereum/wiki/wiki/JSON-RPC
package ethrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ethRequest struct {
	ID      int           `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type ethResponse struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *Error          `json:"error"`
}

// JSONRPC错误
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("ethrpc.Error{code:%d, message: %q}", err.Code, err.Message)
}

// 原生调用方法
func Call(host, method string, args ...interface{}) (result json.RawMessage, err error) {
	return jsonrpcCall(host, method, args...)
}

// jsonrpc调用
func jsonrpcCall(host, method string, args ...interface{}) (result json.RawMessage, err error) {
	reqBody, err := json.Marshal(&ethRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  args,
		ID:      1,
	})
	if err != nil {
		return nil, err
	}

	r, err := http.Post(host, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var resp = new(ethResponse)
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Result, nil
}

// 以太坊RPC调用, 需要通过response提供返回类型
func ethrpcCall(host, method string, response interface{}, args ...interface{}) (err error) {
	resp_Result, err := jsonrpcCall(host, method, args...)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resp_Result, response); err != nil {
		return err
	}
	return nil
}
