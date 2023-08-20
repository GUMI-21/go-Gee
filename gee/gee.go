package gee

import (
	"net/http"
)

// 框架入口

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(ctx *Context)

type Engine struct {
	router *Router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	engine.router.AddRoute(method, pattern, handlerFunc)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defined the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.Handle(c)
}
