# EVM JSONRPC Mock Server

## Installation

```shell
go get github.com/minhngoc270499/evm-mock
```

## Usage

Running a mock server

```go
package test

import "github.com/minhngoc270499/evm-mock"

func TestWithMockServer() {
	ms := evmmock.NewServer()
	ms.Start()
	defer ms.Stop()
}
```

Add stub

```go
package test

import "github.com/minhngoc270499/evm-mock"

func TestWithMockServerAndStub() {
	ms := evmmock.NewServer()
	ms.On(evmmock.NewRequest()).WillReturn(evmmock.NewResponse())
	ms.Start()
	defer ms.Stop()
}
```
