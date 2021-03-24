// test mock server
package testutils_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/xplorfin/netutils/testutils"

	"github.com/jarcoal/httpmock"
)

var mockResponse = []byte(`{"id": 1, "name": "My Great Article"}`)

// allows calling a custom handler func
func TestCustomHandlerFunc(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testServer := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		// Send response to be tested
		_, _ = rw.Write(mockResponse)
	})

	httpmock.RegisterResponder("GET", "http://api.entropy.rocks/test", testutils.WrapHandler(testServer))
	rawRes, err := http.Get("http://api.entropy.rocks/test")
	if err != nil {
		t.Error(err)
	}

	res, err := ioutil.ReadAll(rawRes.Body)
	if err != nil {
		t.Error(err)
	}

	testutils.AssertJSONEquals(res, mockResponse, t)
}

func TestMockFile(t *testing.T) {
	file := testutils.MockFile(t)
	if !FileExists(file) {
		t.Errorf("expected file %s created by mockfile to exist", file)
	}
}

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func TestMockExamples(t *testing.T) {
	ExampleMockHTTPServer()
}

func ExampleMockHTTPServer() {
	// turn on http mocking
	httpmock.Activate()
	defer httpmock.Deactivate()
	requestCount := 0
	testServer := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestCount++
		rw.WriteHeader(200)
		_, _ = rw.Write([]byte(strconv.Itoa(requestCount)))
	})

	httpmock.RegisterResponder("GET", "https://requestcounter.com", testutils.WrapHandler(testServer))

	resp, err := http.Get("https://requestcounter.com")
	if err != nil {
		panic(err)
	}
	// print the response count, should be 1
	fmt.Print(resp)
}
