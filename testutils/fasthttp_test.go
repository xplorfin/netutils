package testutils_test

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/xplorfin/netutils/testutils"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/valyala/fasthttp"
)

type ClientMethod int

const (
	Dial   ClientMethod = 0
	Client ClientMethod = 1
)

type TestCase struct {
	Method ClientMethod
}

// test cases to run for fasthttp and http
var testCases = []TestCase{
	{
		Method: Dial,
	},
	{
		Method: Client,
	},
}

func makeSimpleRequest() *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("/uri") // task URI
	req.Header.SetMethod("GET")
	req.Header.SetHost("hi")
	return req
}

func TestFastHttpClient(t *testing.T) {
	for _, test := range testCases {
		var client fasthttp.Client
		body := gofakeit.Sentence(gofakeit.Number(1, 3))
		server := testutils.NewFastHTTPMock(t)

		server.Start(func(ctx *fasthttp.RequestCtx) {
			ctx.Response.SetBodyString(body)
		})

		switch test.Method {
		case Dial:
			client = fasthttp.Client{
				Dial: func(addr string) (net.Conn, error) {
					return server.Dial()
				},
			}
		case Client:
			client = *server.FastHTTPMockClient()
		}

		req := makeSimpleRequest()
		resp := fasthttp.AcquireResponse()
		err := client.Do(req, resp)
		if err != nil {
			t.Error(err)
		}
		if string(resp.Body()) != body {
			t.Errorf("expected response: %s to match body: %s", resp.String(), body)
		}
	}
}

func TestHttpClient(t *testing.T) {
	for _, test := range testCases {
		var client http.Client
		body := gofakeit.Sentence(gofakeit.Number(1, 3))
		server := testutils.NewFastHTTPMock(t)

		server.Start(func(ctx *fasthttp.RequestCtx) {
			ctx.Response.SetBodyString(body)
		})

		switch test.Method {
		case Dial:
			client = http.Client{
				Transport: &http.Transport{
					DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
						return server.Dial()
					},
				},
			}

		case Client:
			client = server.HTTPMockClient()
		}

		resp, err := client.Get("http://test")
		if err != nil {
			t.Fatal(err)
		}
		resultBody, _ := ioutil.ReadAll(resp.Body)

		if string(resultBody) != body {
			t.Errorf("expected response: %s to match body: %s", resultBody, body)
		}
	}
}
