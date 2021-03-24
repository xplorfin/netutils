package fasthttp

import (
	"bytes"

	"github.com/valyala/fasthttp"
)

// IsJSON determines whether a fasthttp response is brotli or json
func IsJSON(resp *fasthttp.Response) bool {
	contentType := resp.Header.PeekBytes(ContentType)
	return bytes.Index(contentType, []byte("application/json")) == 0
}

// UnzipBody unzips a fasthttp body that uses brotli, gzip, or borth
func UnzipBody(resp *fasthttp.Response) []byte {
	contentEncoding := resp.Header.PeekBytes(ContentEncoding)
	var body []byte
	if bytes.EqualFold(contentEncoding, Gzip) {
		body, _ = resp.BodyGunzip()
	} else if bytes.EqualFold(contentEncoding, Brotli) || bytes.EqualFold(contentEncoding, BrotliShort) {
		body, _ = resp.BodyUnbrotli()
	} else {
		body = resp.Body()
	}
	return body
}
