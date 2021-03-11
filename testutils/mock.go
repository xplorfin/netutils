package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/jarcoal/httpmock"
)

func WrapHandler(handler http.Handler) httpmock.Responder {
	return func(request *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, request)
		return w.Result(), nil
	}
}

// create a file with random contents and return the location
func MockFile(t *testing.T) string {
	return filet.TmpFile(t, "", gofakeit.Paragraph(4, 4, 4, " ")).Name()
}
