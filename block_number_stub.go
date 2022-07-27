package evmmock

import "fmt"

func NewBlockNumberStub(blockNumber int64) *StubRule {
	result := fmt.Sprintf("0x%X", blockNumber)
	return NewStub(NewRequest().WithMethod("eth_blockNumber")).WillReturn(NewResponse().WithResult(result))
}
