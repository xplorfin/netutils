package testutils

import (
	"net/http"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// FastRouter creates a fasthttp router that is compatible with standard
// http.Handler's in golang
type FastRouter struct {
	Router *router.Router
}

// NewRouter allows you to wrap a fast http server around regular http.handler funcs
func NewRouter() *FastRouter {
	r := router.New()
	return &FastRouter{r}
}

// Handler returns the fasthttp.RequestHandler object for use in a server/with a mock client
func (r *FastRouter) Handler() fasthttp.RequestHandler {
	return r.Router.Handler
}

// request handlers

// GETFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) GETFastHTTP(path string, handler fasthttp.RequestHandler) {
	r.Router.GET(path, handler)
}

// GET returns a request through the FastRouter.Router handler
func (r FastRouter) GET(path string, handler http.HandlerFunc) {
	r.GETFastHTTP(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// PutFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) PutFastHTTP(path string, handler fasthttp.RequestHandler) {
	r.Router.PUT(path, handler)
}

// PUT returns a request through the FastRouter.Router handler
func (r FastRouter) PUT(path string, handler http.HandlerFunc) {
	r.PutFastHTTP(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// PostFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) PostFastHTTP(path string, handler fasthttp.RequestHandler) {
	r.Router.POST(path, handler)
}

// POST returns a request through the FastRouter.Router handler
func (r FastRouter) POST(path string, handler http.HandlerFunc) {
	r.PostFastHTTP(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// PatchFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) PatchFastHTTP(path string, handler fasthttp.RequestHandler) {
	r.Router.PATCH(path, handler)
}

// PATCH returns a request through the FastRouter.Router handler
func (r FastRouter) PATCH(path string, handler http.HandlerFunc) {
	r.PatchFastHTTP(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// DeleteFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) DeleteFastHTTP(path string, handler fasthttp.RequestHandler) {
	r.Router.DELETE(path, handler)
}

// DELETE returns a request through the FastRouter.Router handler
func (r FastRouter) DELETE(path string, handler http.HandlerFunc) {
	r.DeleteFastHTTP(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// ConnectFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) ConnectFastHTTP(path string, handler fasthttp.RequestHandler) {
	r.Router.CONNECT(path, handler)
}

// CONNECT returns a request through the FastRouter.Router handler
func (r FastRouter) CONNECT(path string, handler http.HandlerFunc) {
	r.ConnectFastHTTP(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// OptionsFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) OptionsFastHTTP(path string, handler fasthttp.RequestHandler) {
	r.Router.OPTIONS(path, handler)
}

// OPTIONS returns a request through the FastRouter.Router handler
func (r FastRouter) OPTIONS(path string, handler http.HandlerFunc) {
	r.OptionsFastHTTP(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// TraceFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) TraceFastHTTP(path string, handler fasthttp.RequestHandler) {
	r.Router.TRACE(path, handler)
}

// TRACE returns a request through the FastRouter.Router handler
func (r FastRouter) TRACE(path string, handler http.HandlerFunc) {
	r.TraceFastHTTP(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// AnyFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) AnyFastHTTP(path string, handler fasthttp.RequestHandler) {
	r.Router.ANY(path, handler)
}

// ANY returns a request through the FastRouter.Router handler
func (r FastRouter) ANY(path string, handler http.HandlerFunc) {
	r.AnyFastHTTP(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// Handle returns a request through the FastRouter.Router handler
func (r FastRouter) Handle(method, path string, handler http.HandlerFunc) {
	r.Router.Handle(method, path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

// HandleFastHTTP returns a request through the FastRouter.Router handler
func (r FastRouter) HandleFastHTTP(method, path string, handler fasthttp.RequestHandler) {
	r.Router.Handle(method, path, handler)
}
