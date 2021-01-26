package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMockHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := MockHandler()
	handler.ServeHTTP(rr, req)

	if http.StatusBadRequest != rr.Code {
		t.Errorf("expected mock handler to return status code %d, returned %d", http.StatusBadRequest, rr.Code)
	}
}
