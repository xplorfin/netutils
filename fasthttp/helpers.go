package fasthttp

import (
	"bytes"

	"github.com/valyala/fasthttp"
)

func IsJson(resp *fasthttp.Response) bool {
	contentType := resp.Header.PeekBytes(ContentType)
	return bytes.Index(contentType, []byte("application/json")) == 0
}

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
