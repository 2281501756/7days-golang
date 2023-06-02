package gee

import (
	"net/http"
	"path"
)

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

func (r *RouterGroup) Static(relativePath string, root string) {
	handle := r.createStaticHandle(relativePath, http.Dir(root))
	urlPath := path.Join(relativePath, "/*filepath")
	r.GET(urlPath, handle)
}
func (r *RouterGroup) createStaticHandle(pattern string, handle http.FileSystem) HandleFunc {
	absolutePath := path.Join(r.prefix, pattern)
	serveFileHandle := http.StripPrefix(absolutePath, http.FileServer(handle))
	return func(c *Context) {
		// 获取传入参数中文件的路径
		file := c.Param("filepath")
		// 查看文件是否能打开，如果不能则返回404
		if _, err := handle.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		// 有文件传入context
		serveFileHandle.ServeHTTP(c.Writer, c.Req)
	}
}
