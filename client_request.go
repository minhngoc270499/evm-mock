package evmmock

type ClientRequest struct {
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      uint64        `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
}
