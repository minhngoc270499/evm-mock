package evmmock

// Request ...
type Request struct {
	Method        string
	ParamMatchers []ParamMatcher
}

// NewRequest ...
func NewRequest() *Request {
	return &Request{}
}

// WithMethod ...
func (r *Request) WithMethod(method string) *Request {
	r.Method = method
	return r
}

// WithParams ...
func (r *Request) WithParams(param ParamMatcher) *Request {
	r.ParamMatchers = append(r.ParamMatchers, param)
	return r
}
