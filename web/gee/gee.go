package gee

import (
	"net/http"
	"strings"
)

type HandleFunc func(c *Context)

type Engine struct {
	router *Router
	groups []*RouterGroup // 保存所有的RouterGroup

	*RouterGroup // 继承routerGroup让engine作为最上层的router
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = append(engine.groups, engine.RouterGroup)
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middleware []HandleFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middleware = append(middleware, group.middlewares...)
		}
	}
	c := newContext(w, r)
	c.handlers = middleware
	e.router.handle(c)
}

func (e *Engine) Run(port string) error {
	return http.ListenAndServe(port, e)
}
