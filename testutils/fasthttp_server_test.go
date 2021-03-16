package testutils

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
)

const HelloString = "Welcome!"

// Index is a request handler that writes a string to a mock server
func Index(ctx *fasthttp.RequestCtx) {
	_, _ = ctx.WriteString(HelloString)
}

// TestFastClient tests a fasthttp client omcker
func TestFastClient(t *testing.T) {
	server := NewFastHTTPMock(t)
	r := NewRouter()
	r.GETFastHTTP("/", Index)
	r.GET("/test", func(rw http.ResponseWriter, request *http.Request) {
		_, _ = rw.Write([]byte(HelloString))
	})

	server.Start(r.Handler())

	// fast http handler
	client := server.HTTPMockClient()
	resp, err := client.Get("http://test/")
	if err != nil {
		t.Fatal(err)
	}
	resultBody, _ := ioutil.ReadAll(resp.Body)
	if string(resultBody) != HelloString {
		t.Errorf("expected response %s to match %s", resultBody, HelloString)
	}

	// http handler
	resp, err = client.Get("http://test/test")
	if err != nil {
		t.Fatal(err)
	}
	resultBody, _ = ioutil.ReadAll(resp.Body)
	if string(resultBody) != HelloString {
		t.Errorf("expected response %s to match %s", resultBody, HelloString)
	}

}
