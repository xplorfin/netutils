package fasthttp_test

import (
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	fasthttpHelper "github.com/xplorfin/netutils/fasthttp"
	"github.com/xplorfin/netutils/testutils"
)

// setupTests sets up the http tests
func setupTests(t *testing.T) (host string) {
	port := testutils.GetFreePort(t)
	go func() {
		err := testutils.MockHTTPServer(port)
		Nil(t, err)
	}()
	host = fmt.Sprintf("http://localhost:%d", port)
	testutils.AssertConnected(host, t)
	return host
}

// TestEmptyHttpClient tests a client with no data
func TestEmptyHttpClient(t *testing.T) {
	host := setupTests(t)

	testutils.AssertConnected(host, t)

	client := fasthttpHelper.NewFastClient()
	response, err := client.Request(host)
	if err != nil {
		t.Error(err)
	}
	_ = response
}

// TestHttpClientHooks tests the client hooks
func TestHttpClientHooks(t *testing.T) {
	var modifyCalled, processCalled bool

	host := setupTests(t)

	testutils.AssertConnected(host, t)

	client := fasthttpHelper.NewFastClient()

	// since this is called after seturi and we set an empty url below
	// this should test whether or not request is being modified
	client.ModifyRequest = func(request *fasthttp.Request) {
		request.SetRequestURI(host)
		modifyCalled = true
	}

	client.ProcessResponse = func(response *fasthttp.Response) {
		processCalled = true
		Equal(t, response.StatusCode(), 400)
	}

	response, err := client.Request("")
	if err != nil {
		t.Error(err)
	}
	_ = response
	True(t, modifyCalled)
	True(t, processCalled)
}
