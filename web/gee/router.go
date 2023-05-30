package gee

type handleFunc func(c *Context)

type Router struct {
	handles map[string]handleFunc
}

func (r *Router) addRouter(method string, pattern string, handle handleFunc) {
	key := method + "-" + pattern
	r.handles[key] = handle
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handle, ok := r.handles[key]; ok {
		handle(c)
	} else {
		c.Writer.Write([]byte("404 not found"))
	}
}
