package evmmock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStubRule_IsMatched_WhenEqualToJson(t *testing.T) {
	stubRule := NewStub(NewRequest().WithMethod(EthGetLogsMethod).WithParams(EqualToJson(`[{"fromBlock":"0x01","toBlock":"0x01"}]`)))
	matched := stubRule.IsMatched(&ClientRequest{
		Method:  "eth_getLogs",
		Params:  []interface{}{map[string]string{"fromBlock": "0x01", "toBlock": "0x01"}},
		ID:      1,
		JsonRpc: "2.0",
	})
	assert.True(t, matched)
}

func TestStubRule_IsMatched_WhenNotEqualToJson(t *testing.T) {
	stubRule := NewStub(NewRequest().WithMethod(EthGetLogsMethod).WithParams(EqualToJson(`[{"fromBlock":"0x01","toBlock":"0x01"}]`)))
	matched := stubRule.IsMatched(&ClientRequest{
		Method:  "eth_getLogs",
		Params:  []interface{}{map[string]string{"fromBlock": "0x01", "toBlock": "0x02"}},
		ID:      1,
		JsonRpc: "2.0",
	})
	assert.False(t, matched)
}
