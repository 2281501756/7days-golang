package gee

type RouterGroup struct {
	prefix      string
	parent      *RouterGroup
	middlewares []handleFunc
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

func (r *RouterGroup) addRouter(method string, comp string, handle handleFunc) {
	pattern := r.prefix + comp
	r.engine.router.addRouter(method, pattern, handle)
}

func (r *RouterGroup) GET(pattern string, handle handleFunc) {
	r.addRouter("GET", pattern, handle)
}
func (r *RouterGroup) POST(pattern string, handle handleFunc) {
	r.addRouter("POST", pattern, handle)
}
func (r *RouterGroup) PUT(pattern string, handle handleFunc) {
	r.addRouter("PUT", pattern, handle)
}
func (r *RouterGroup) DELETE(pattern string, handle handleFunc) {
	r.addRouter("DELETE", pattern, handle)
}
