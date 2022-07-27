package evmmock

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
)

const defaultAddress = "127.0.0.1:8545"

// Server ...
type Server struct {
	ts        *httptest.Server
	stubRules []*StubRule
}

// NewServer ...
func NewServer(address ...string) *Server {
	s := &Server{
		stubRules: make([]*StubRule, 0),
	}
	adr := defaultAddress
	if len(address) > 0 {
		adr = address[0]
	}
	l, err := net.Listen("tcp", adr)
	if err != nil {
		panic(fmt.Errorf("rpc server: listen: %v", err))
	}
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request ClientRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w)
			return
		}
		for _, stubRule := range s.stubRules {
			if stubRule.IsMatched(&request) {
				if stubRule.handler != nil {
					stubRule.handler(&request, w)
				}
				w.WriteHeader(http.StatusOK)
				_, _ = fmt.Fprintln(w, stubRule.response.ToString())
				return
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "request %+v is not matched", request)
		return
	}))
	if err := ts.Listener.Close(); err != nil {
		panic(fmt.Errorf("rpc server: close server: %v", err))
	}
	ts.Listener = l
	s.ts = ts
	return s
}

// On ...
func (s *Server) On(stubRule *StubRule) *Server {
	s.stubRules = append(s.stubRules, stubRule)
	return s
}

// Start ...
func (s *Server) Start() {
	s.ts.Start()
}

// Stop ...
func (s *Server) Stop() {
	s.ts.Close()
}
