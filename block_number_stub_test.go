package evmmock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBlockNumberStub(t *testing.T) {
	testCases := []struct {
		blockNumber int64
		hex         string
	}{
		{
			blockNumber: 0,
			hex:         "0x0",
		},
		{
			blockNumber: 1,
			hex:         "0x1",
		},
		{
			blockNumber: 255,
			hex:         "0xFF",
		},
	}
	for _, testCase := range testCases {
		stub := NewBlockNumberStub(testCase.blockNumber)
		assert.Equal(t, testCase.hex, stub.response.Result)
	}
}
