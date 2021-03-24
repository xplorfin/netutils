package fasthttp

import (
	"github.com/valyala/fasthttp"
)

const UserAgent = "Mozilla/5.0 (compatible; bingbot/2.0; +https://github.com/xplorfin/netutils)"

var UserAgentBytes = []byte(UserAgent)

func GetRawUrl(uri []byte) ([]byte, error) {
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

// for testing, get url using a string
func GetUrl(uri string) ([]byte, error) {
	return GetRawUrl([]byte(uri))
}
