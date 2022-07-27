package evmmock

import (
	"encoding/json"
	"net/http"
	"strings"
)

// StubRule ...
type StubRule struct {
	request  *Request
	response *Response
	handler  StubFunc
}

// StubFunc ...
type StubFunc func(request *ClientRequest, writer http.ResponseWriter)

// NewStub ...
func NewStub(request *Request) *StubRule {
	return &StubRule{
		request: request,
	}
}

// WillReturn ...
func (sr *StubRule) WillReturn(response *Response) *StubRule {
	sr.response = response
	return sr
}

// WillHandle ...
func (sr *StubRule) WillHandle(handler StubFunc) *StubRule {
	sr.handler = handler
	return sr
}

// IsMatched ...
func (sr *StubRule) IsMatched(request *ClientRequest) bool {
	if !strings.EqualFold(request.Method, sr.request.Method) {
		return false
	}
	if !sr.isMatchedParams(request) {
		return false
	}
	return true
}

func (sr *StubRule) isMatchedParams(request *ClientRequest) bool {
	for _, paramMatcher := range sr.request.ParamMatchers {
		if paramMatcher.Strategy() == ParamEqualToJson {
			if !sr.isEqualToJson(request.Params, paramMatcher) {
				return false
			}
		}
	}
	return true
}

func (sr *StubRule) isEqualToJson(params []interface{}, matcher ParamMatcher) bool {
	requestParamsInBytes, err := json.Marshal(params)
	if err != nil {
		return false
	}
	requestParamsString := strings.ReplaceAll(string(requestParamsInBytes), " ", "")
	matcherValue := strings.ReplaceAll(matcher.Value(), " ", "")
	return requestParamsString == matcherValue
}
