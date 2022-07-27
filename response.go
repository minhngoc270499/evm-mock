package evmmock

import (
	"encoding/json"
	"fmt"
)

// Response represents RPC response
type Response struct {
	JSONRPC string      `json:"json_rpc"`
	ID      int64       `json:"id"`
	Result  interface{} `json:"result"`
}

// NewResponse return a new instance of request
func NewResponse() *Response {
	return &Response{
		ID:      1,
		JSONRPC: "2.0",
		Result:  nil,
	}
}

// WithID ...
func (r *Response) WithID(id int64) *Response {
	r.ID = id
	return r
}

// WithJsonRpc ...
func (r *Response) WithJsonRpc(version string) *Response {
	r.JSONRPC = version
	return r
}

// WithResult ...
func (r *Response) WithResult(result interface{}) *Response {
	if v, ok := result.(string); ok {
		var js json.RawMessage
		if json.Unmarshal([]byte(v), &js) == nil {
			r.Result = js
		} else {
			r.Result = result
		}
	} else {
		r.Result = result
	}
	return r
}

func (r *Response) ToString() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		panic(fmt.Errorf("could not marshall: %v", err))
	}
	return string(marshal)
}
