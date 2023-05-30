package gee

import "net/http"

type handleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]handleFunc
}

func New() *Engine {
	return &Engine{router: map[string]handleFunc{}}
}

func (e *Engine) addRouter(method string, pattern string, handle handleFunc) {
	key := method + "-" + pattern
	e.router[key] = handle
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
	key := r.Method + "-" + r.URL.Path
	if handle, ok := e.router[key]; ok {
		handle(w, r)
	} else {
		w.Write([]byte("404 not found"))
	}
}

func (e *Engine) Run(port string) error {
	return http.ListenAndServe(port, e)
}
