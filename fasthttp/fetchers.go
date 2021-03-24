package fasthttp

import (
	"github.com/valyala/fasthttp"
)

// UserAgent contains the netutils user agent
const UserAgent = "Mozilla/5.0 (compatible; netutils/1.0; +https://github.com/xplorfin/netutils)"

// UserAgentBytes contains the raw user agent encoded as a byte slice
var UserAgentBytes = []byte(UserAgent)

// GetRawURL gets a rawurl using fasthttp
func GetRawURL(uri []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURIBytes(uri)
	req.Header.SetUserAgentBytes(UserAgentBytes)
	req.Header.SetBytesKV(AcceptEncoding, GzipBrotli)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil, err
	}

	// handle gzip
	return UnzipBody(resp), nil
}

// GetURL gets a string url and unzips body. This is mostly a utility function
func GetURL(uri string) ([]byte, error) {
	return GetRawURL([]byte(uri))
}
