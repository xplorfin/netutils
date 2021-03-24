package fasthttp_test

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v5"

	// use a seperate lib to test parity
	"net/http"
	"testing"

	"github.com/go-http-utils/headers"

	. "github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	fasthttpHelper "github.com/xplorfin/netutils/fasthttp"
	"github.com/xplorfin/netutils/testutils"
)

// FastTest is the test helper
type FastTest struct {
	// nonJSONReply is the non json reply to check against
	nonJSONReply string
	// jsonReply is the json reply to check against
	jsonReply []byte
	// serverURL is the servers URL
	serverURL string
}

// getTestServer generates a test server and data to test against
func getTestServer(t *testing.T) (fastTest FastTest) {
	var err error
	fastTest.nonJSONReply = gofakeit.Word()
	// test json reply
	fastTest.jsonReply, err = gofakeit.JSON(&gofakeit.JSONOptions{
		Type: "object",
		Fields: []gofakeit.Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "address", Function: "address"},
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
		},
		Indent: true,
	})
	Nil(t, err)

	server := testutils.NewFastHTTPMock(t)
	r := testutils.NewRouter()
	r.GET("/non-json", func(rw http.ResponseWriter, request *http.Request) {
		_, _ = rw.Write([]byte(fastTest.nonJSONReply))
	})
	r.GET("/json", func(rw http.ResponseWriter, request *http.Request) {
		rw.Header().Set(headers.ContentType, string(fasthttpHelper.JSONEncoding))
		_, _ = rw.Write(fastTest.jsonReply)
	})
	server.Start(r.Handler())
	fastTest.serverURL = fmt.Sprintf("localhost:%d", testutils.GetFreePort(t))
	go func() {
		err = fasthttp.ListenAndServe(fastTest.serverURL, r.Handler())
		Nil(t, err)
	}()
	testutils.AssertConnected(fastTest.serverURL, t)
	return fastTest
}

// TestGetRawURL tests the GetRawURL method
func TestGetRawURL(t *testing.T) {
	fastTest := getTestServer(t)
	testutils.AssertConnected(fastTest.serverURL, t)

	jsonURL := fmt.Sprintf("http://%s/json", fastTest.serverURL)
	jsonResponse, err := fasthttpHelper.GetRawURL([]byte(jsonURL))
	Nil(t, err)
	Equal(t, jsonResponse, fastTest.jsonReply)
}

// TestGetURL tests the GetURL method
func TestGetURL(t *testing.T) {
	fastTest := getTestServer(t)
	testutils.AssertConnected(fastTest.serverURL, t)

	jsonURL := fmt.Sprintf("http://%s/json", fastTest.serverURL)
	jsonResponse, err := fasthttpHelper.GetURL(jsonURL)
	Nil(t, err)
	Equal(t, jsonResponse, fastTest.jsonReply)
}

func TestIsJSON(t *testing.T) {
	fastTest := getTestServer(t)
	testutils.AssertConnected(fastTest.serverURL, t)
	jsonURL := fmt.Sprintf("http://%s/json", fastTest.serverURL)
	// test the json response
	jsonResponse, err := testutils.GetFastURLResponse([]byte(jsonURL), func(resp *fasthttp.Response) {
		True(t, fasthttpHelper.IsJSON(resp))
	})
	Nil(t, err)
	Equal(t, jsonResponse, fastTest.jsonReply)

	nonJSONURL := fmt.Sprintf("http://%s/non-json", fastTest.serverURL)
	nonJSONResponse, err := testutils.GetFastURLResponse([]byte(nonJSONURL), func(resp *fasthttp.Response) {
		False(t, fasthttpHelper.IsJSON(resp))
	})
	Nil(t, err)
	Equal(t, nonJSONResponse, []byte(fastTest.nonJSONReply))
}

func TestFastHttpExamples(t *testing.T) {
	ExampleGetRawURL()
	ExampleGetURL()
}

// gets a byte slice url using fasthttp
func ExampleGetRawURL() {
	// fetch example.org
	res, err := fasthttpHelper.GetRawURL([]byte("https://api.entropy.rocks"))
	if err != nil {
		panic(err)
	}
	// print the response body
	fmt.Println(string(res))
}

// gets a byte slice url using fasthttp
func ExampleGetURL() {
	// fetch example.org
	res, err := fasthttpHelper.GetURL("https://api.entropy.rocks")
	if err != nil {
		panic(err)
	}
	// print the response body
	fmt.Println(string(res))
}
