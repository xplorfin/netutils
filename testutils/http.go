package testutils

import (
	"fmt"
	"net/http"
)

// MockHandler mocks an http handler that's not nil
func MockHandler() http.HandlerFunc {
	// handle `/` route to `http.DefaultServeMux`
	return func(res http.ResponseWriter, req *http.Request) {

		// get response headers
		header := res.Header()

		// set content type header
		header.Set("Content-Type", "application/json")

		// reset date header (inline call)
		res.Header().Set("Date", "01/01/2020")

		// set status header
		res.WriteHeader(http.StatusBadRequest) // http.StatusBadRequest == 400

		// respond with a JSON string
		_, _ = fmt.Fprint(res, `{"status":"FAILURE"}`)
	}
}

// MockHTTPServer runs a mock http server on a given port
func MockHTTPServer(port int) error {
	// handle `/` route to `http.DefaultServeMux`
	handler := MockHandler()
	router := http.NewServeMux()

	router.HandleFunc("/", handler)

	// listen and serve using `http.DefaultServeMux`
	return http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
