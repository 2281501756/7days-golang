package gee

import "net/http"

type handleFunc func(c *Context)

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
	c := newContext(w, r)
	e.router.handle(c)
}

func (e *Engine) Run(port string) error {
	return http.ListenAndServe(port, e)
}
