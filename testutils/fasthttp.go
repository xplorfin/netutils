package testutils

import (
	"context"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
	"net"
	"net/http"
	"testing"
)

// fast http handler - turns client requests into methods locally handled
// by overriding dial
// for wrapping gock/other mock servers
type httpFastHandler struct {
	// the server object
	Listener *fasthttputil.InmemoryListener
	// the test object (for throwing errors)
	Test *testing.T
}

// creates a mock server by overriding client
// this allows us to test fathttp servers without actually standing up a server
func NewFastHttpMock(t *testing.T) *httpFastHandler {
	server := fasthttputil.NewInmemoryListener()
	return &httpFastHandler{
		Listener: server,
		Test:     t,
	}
}

// fast http handler
func (server httpFastHandler) Start(handler fasthttp.RequestHandler) {
	go func() {
		err := fasthttp.Serve(server.Listener, handler)
		if err != nil {
			server.Test.Error(err)
			panic(err)
		}
	}()
}

// Create a dial object cooresponding to mock server
func (server httpFastHandler) Dial() (net.Conn, error) {
	return server.Listener.Dial()
}

// fast http client with the server name
// note this will override every request with dial
func (server httpFastHandler) FastHttpMockClient() *fasthttp.Client {
	return &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return server.Dial()
		},
	}
}

// fast http client with the server name
// note this will override every request with dial
func (server httpFastHandler) HttpMockClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return server.Dial()
			},
		},
	}
}
