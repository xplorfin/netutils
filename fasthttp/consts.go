package fasthttp

import (
	"github.com/valyala/fasthttp"
)

// Byte encoded headers for convience
var (
	// AcceptEncoding is a byte encoded header for accepting encoding
	// used to speed up fasthttp requests. Defined here: https://git.io/JYTc2
	AcceptEncoding = []byte(fasthttp.HeaderAcceptEncoding)
	// Gzip is the gzip header defined here https://git.io/JYTcg
	Gzip = []byte("gzip")
	// Authorization contains the authorization header
	Authorization = []byte(fasthttp.HeaderAuthorization)
	// ContentEncoding is the
	ContentEncoding = []byte(fasthttp.HeaderContentEncoding)
	ContentType     = []byte(fasthttp.HeaderContentType)
	// Brotli is the long brotli header
	Brotli = []byte("brotli")
	// BrotliShort is the abbreviated brotli header
	BrotliShort = []byte("br")
	// JsonEncoding is the header for when data is json encoded
	JsonEncoding = []byte("application/json")
	// GzipBrotli is the header for data encoded with both gzip and brotli
	GzipBrotli = []byte("gzip, brotli")
)
