// test mock server
package testutils

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

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
		rw.Write(mockResponse)
	})

	httpmock.RegisterResponder("GET", "http://api.entropy.rocks/test", WrapHandler(testServer))
	rawRes, err := http.Get("http://api.entropy.rocks/test")
	if err != nil {
		t.Error(err)
	}

	res, err := ioutil.ReadAll(rawRes.Body)
	if err != nil {
		t.Error(err)
	}

	AssertJsonEquals(res, mockResponse, t)
}

func TestMockFile(t *testing.T) {
	file := MockFile(t)
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
