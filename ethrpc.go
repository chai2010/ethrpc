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
	Error   *EthError       `json:"error"`
}

type EthError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *EthError) Error() string {
	return fmt.Sprintf("EthError{code:%d, message: %q}", err.Code, err.Message)
}

func jsonrpcCall(host, method string, response interface{}, args ...interface{}) (err error) {
	reqBody, err := json.Marshal(&ethRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  args,
		ID:      1,
	})
	if err != nil {
		return err
	}

	r, err := http.Post(host, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var resp = new(ethResponse)
	if err := json.Unmarshal(respBody, resp); err != nil {
		return err
	}
	if resp.Error != nil {
		return resp.Error
	}

	if err := json.Unmarshal(resp.Result, response); err != nil {
		return err
	}
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}
