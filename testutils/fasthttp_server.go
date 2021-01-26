package testutils

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
)

type fastRouter struct {
	Router *router.Router
}

// this module allows you to wrap a fast http server around regular http.handler funcs
func NewRouter() *fastRouter {
	r := router.New()
	return &fastRouter{r}
}

func (r *fastRouter) Handler() fasthttp.RequestHandler {
	return r.Router.Handler
}

// request handlers

func (r fastRouter) GETFastHttp(path string, handler fasthttp.RequestHandler) {
	r.Router.GET(path, handler)
}

func (r fastRouter) GET(path string, handler http.HandlerFunc) {
	r.GETFastHttp(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) PutFastHttp(path string, handler fasthttp.RequestHandler) {
	r.Router.PUT(path, handler)
}

func (r fastRouter) PUT(path string, handler http.HandlerFunc) {
	r.PutFastHttp(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) PostFastHttp(path string, handler fasthttp.RequestHandler) {
	r.Router.POST(path, handler)
}

func (r fastRouter) POST(path string, handler http.HandlerFunc) {
	r.PostFastHttp(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) PatchFastHttp(path string, handler fasthttp.RequestHandler) {
	r.Router.PATCH(path, handler)
}

func (r fastRouter) PATCH(path string, handler http.HandlerFunc) {
	r.PatchFastHttp(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) DeleteFastHttp(path string, handler fasthttp.RequestHandler) {
	r.Router.DELETE(path, handler)
}

func (r fastRouter) DELETE(path string, handler http.HandlerFunc) {
	r.DeleteFastHttp(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) ConnectFastHttp(path string, handler fasthttp.RequestHandler) {
	r.Router.CONNECT(path, handler)
}

func (r fastRouter) CONNECT(path string, handler http.HandlerFunc) {
	r.ConnectFastHttp(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) OptionsFastHttp(path string, handler fasthttp.RequestHandler) {
	r.Router.OPTIONS(path, handler)
}

func (r fastRouter) OPTIONS(path string, handler http.HandlerFunc) {
	r.OptionsFastHttp(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) TraceFastHttp(path string, handler fasthttp.RequestHandler) {
	r.Router.TRACE(path, handler)
}

func (r fastRouter) TRACE(path string, handler http.HandlerFunc) {
	r.TraceFastHttp(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) AnyFastHttp(path string, handler fasthttp.RequestHandler) {
	r.Router.ANY(path, handler)
}

func (r fastRouter) ANY(path string, handler http.HandlerFunc) {
	r.AnyFastHttp(path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) Handle(method, path string, handler http.HandlerFunc) {
	r.Router.Handle(method, path, fasthttpadaptor.NewFastHTTPHandlerFunc(handler))
}

func (r fastRouter) HandleFastHttp(method, path string, handler fasthttp.RequestHandler) {
	r.Router.Handle(method, path, handler)
}
