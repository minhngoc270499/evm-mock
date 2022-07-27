package evmmock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponse_ToString(t *testing.T) {
	t.Run("WithoutNilResult", func(t *testing.T) {
		response := NewResponse().WithResult(nil).ToString()
		assert.Equal(t, `{"json_rpc":"2.0","id":1,"result":null}`, response)
	})
	t.Run("WithString", func(t *testing.T) {
		response := NewResponse().WithResult("string").ToString()
		assert.Equal(t, `{"json_rpc":"2.0","id":1,"result":"string"}`, response)
	})
	t.Run("WithNumeric", func(t *testing.T) {
		response := NewResponse().WithResult(1).ToString()
		assert.Equal(t, `{"json_rpc":"2.0","id":1,"result":1}`, response)
	})
	t.Run("WithStruct", func(t *testing.T) {
		response := NewResponse().WithResult(struct {
			Example string `json:"example"`
		}{
			Example: "example",
		}).ToString()
		assert.Equal(t, `{"json_rpc":"2.0","id":1,"result":{"example":"example"}}`, response)
		response = NewResponse().WithResult(struct {
			Example     string      `json:"example"`
			OtherStruct interface{} `json:"other_struct"`
		}{
			Example: "example",
			OtherStruct: struct {
				OtherKey string `json:"other_key"`
			}{
				OtherKey: "other_key",
			},
		}).ToString()
		assert.Equal(t, `{"json_rpc":"2.0","id":1,"result":{"example":"example","other_struct":{"other_key":"other_key"}}}`, response)
	})
	t.Run("WithJsonString", func(t *testing.T) {
		response := NewResponse().WithResult(`{"example":"example"}`).ToString()
		assert.Equal(t, `{"json_rpc":"2.0","id":1,"result":{"example":"example"}}`, response)
		response = NewResponse().WithResult(`[{"example":"example"},{"example":"example"}]`).ToString()
		assert.Equal(t, `{"json_rpc":"2.0","id":1,"result":[{"example":"example"},{"example":"example"}]}`, response)
	})
}
