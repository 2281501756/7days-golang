package gee

import (
	"strings"
)

type Router struct {
	handleMap map[string]handleFunc // 路由与handle的map
	trieMap   map[string]*node      // 前缀树的map key是 GET POST这些value是路由的前缀树
}

func newRouter() *Router {
	return &Router{
		handleMap: map[string]handleFunc{},
		trieMap:   map[string]*node{},
	}
}

func parsePattern(pattern string) []string {
	res := make([]string, 0)
	arr := strings.Split(pattern, "/")
	for _, i := range arr {
		if i != "" {
			res = append(res, i)
			if i[0] == '*' {
				break
			}
		}
	}
	return res
}

func (r *Router) addRouter(method string, pattern string, handle handleFunc) {
	key := method + "-" + pattern
	parseArray := parsePattern(pattern)
	if r.trieMap[method] == nil {
		r.trieMap[method] = &node{}
	}
	r.trieMap[method].insert(pattern, parseArray, 0)
	r.handleMap[key] = handle
}

func (r *Router) getRouter(method, path string) (*node, map[string]string) {
	parseArray := parsePattern(path)
	if r.trieMap[method] == nil {
		return nil, nil
	}
	n := r.trieMap[method].search(path, parseArray, 0)
	if n == nil {
		return nil, nil
	}
	params := map[string]string{}
	parts := parsePattern(n.pattern)

	for i, item := range parts {
		if item[0] == ':' {
			params[item[1:]] = parseArray[i]
		}
		if item[0] == '*' && len(parts) > 1 {
			params[item[1:]] = strings.Join(parseArray[i:], "/")
			break
		}
	}
	return n, params

}

func (r *Router) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		r.handleMap[key](c)
	} else {
		_, err := c.Writer.Write([]byte("404 not found"))
		if err != nil {
			panic("404 not found")
		}
	}
}
