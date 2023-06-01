package gee

type RouterGroup struct {
	prefix      string
	parent      *RouterGroup
	middlewares []HandleFunc
	engine      *Engine
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		engine: engine,
		parent: r,
		prefix: r.prefix + prefix,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (r *RouterGroup) Use(middleware ...HandleFunc) {
	r.middlewares = append(r.middlewares, middleware...)
}

func (r *RouterGroup) addRouter(method string, comp string, handle HandleFunc) {
	pattern := r.prefix + comp
	r.engine.router.addRouter(method, pattern, handle)
}

func (r *RouterGroup) GET(pattern string, handle HandleFunc) {
	r.addRouter("GET", pattern, handle)
}
func (r *RouterGroup) POST(pattern string, handle HandleFunc) {
	r.addRouter("POST", pattern, handle)
}
func (r *RouterGroup) PUT(pattern string, handle HandleFunc) {
	r.addRouter("PUT", pattern, handle)
}
func (r *RouterGroup) DELETE(pattern string, handle HandleFunc) {
	r.addRouter("DELETE", pattern, handle)
}
