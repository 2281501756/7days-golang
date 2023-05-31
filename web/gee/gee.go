package gee

import "net/http"

type handleFunc func(c *Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRouter(method string, pattern string, handle handleFunc) {
	e.router.addRouter(method, pattern, handle)
}
func (e *Engine) GET(pattern string, handle handleFunc) {
	e.addRouter("GET", pattern, handle)
}
func (e *Engine) POST(pattern string, handle handleFunc) {
	e.addRouter("POST", pattern, handle)
}
func (e *Engine) PUT(pattern string, handle handleFunc) {
	e.addRouter("PUT", pattern, handle)
}
func (e *Engine) DELETE(pattern string, handle handleFunc) {
	e.addRouter("DELETE", pattern, handle)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) Run(port string) error {
	return http.ListenAndServe(port, e)
}
