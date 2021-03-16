package testutils

import (
	"context"
	"net"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

// HTTPFastHandler - turns client requests into methods locally handled
// by overriding dial for wrapping gock/other mock servers
type HTTPFastHandler struct {
	// Listener is the server object
	Listener *fasthttputil.InmemoryListener
	// Test is the test object (for throwing errors)
	Test *testing.T
}

// NewFastHTTPMock creates a mock server by overriding client
// this allows us to test fathttp servers without actually standing up a server
func NewFastHTTPMock(t *testing.T) *HTTPFastHandler {
	server := fasthttputil.NewInmemoryListener()
	return &HTTPFastHandler{
		Listener: server,
		Test:     t,
	}
}

// Start starts the fastHTTP server and routes request to a given handler
func (server HTTPFastHandler) Start(handler fasthttp.RequestHandler) {
	go func() {
		err := fasthttp.Serve(server.Listener, handler)
		if err != nil {
			server.Test.Error(err)
			panic(err)
		}
	}()
}

// Dial creates a dial object corresponding to mock server
func (server HTTPFastHandler) Dial() (net.Conn, error) {
	return server.Listener.Dial()
}

// FastHTTPMockClient creates a fasthttp client with the server name
// note: this will override every request with dial
func (server HTTPFastHandler) FastHTTPMockClient() *fasthttp.Client {
	return &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return server.Dial()
		},
	}
}

// HTTPMockClient creates an http client with the server name
// Note: this will override every request with dial
func (server HTTPFastHandler) HTTPMockClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return server.Dial()
			},
		},
	}
}
