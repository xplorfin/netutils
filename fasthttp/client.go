package fasthttp

import (
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/valyala/fasthttp"
)

// ModifyRequestFunc represents a function that allows
// a request to be modified before it is set
type ModifyRequestFunc func(request *fasthttp.Request)

// ProcessResponseFunc represents a function that allows
// a response to be modified before it is set
type ProcessResponseFunc func(response *fasthttp.Response)

// FastClient represents a client used for requests
type FastClient struct {
	//modify a fast http request before it's send
	ModifyRequest ModifyRequestFunc
	// process a response before it's received
	ProcessResponse ProcessResponseFunc
	// user agent to use
	UserAgent string
}

// create a generic http client with a random user agent
// and modifiers/processors set to do nothing

// NewFastClient creates a fast client with a user agent
// this can be changed if needed in ModifyRequest
func NewFastClient() FastClient {
	return FastClient{
		ModifyRequest:   nil,
		ProcessResponse: nil,
		UserAgent:       browser.Random(),
	}
}

// Request requests a url
func (h FastClient) Request(url string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetUserAgent(h.UserAgent)
	req.Header.SetBytesKV(AcceptEncoding, GzipBrotli)

	if h.ModifyRequest != nil {
		h.ModifyRequest(req)
	}

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	if h.ProcessResponse != nil {
		h.ProcessResponse(resp)
	}

	respBody := UnzipBody(resp)

	return respBody, nil
}
