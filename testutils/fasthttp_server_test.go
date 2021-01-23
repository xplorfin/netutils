package testutils

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
)

const HelloString = "Welcome!"

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString(HelloString)
}

func TestFastClient(t *testing.T) {
	server := NewFastHttpMock(t)
	r := NewRouter()
	r.GETFastHttp("/", Index)
	r.GET("/test", func(rw http.ResponseWriter, request *http.Request) {
		rw.Write([]byte(HelloString))
	})

	server.Start(r.Handler())

	// fast http handler
	client := server.HttpMockClient()
	resp, err := client.Get("http://test/")
	if err != nil {
		t.Fatal(err)
	}
	resultBody, err := ioutil.ReadAll(resp.Body)
	if string(resultBody) != HelloString {
		t.Errorf("expected response %s to match %s", resultBody, HelloString)
	}

	// http handler
	resp, err = client.Get("http://test/test")
	if err != nil {
		t.Fatal(err)
	}
	resultBody, err = ioutil.ReadAll(resp.Body)
	if string(resultBody) != HelloString {
		t.Errorf("expected response %s to match %s", resultBody, HelloString)
	}

}
