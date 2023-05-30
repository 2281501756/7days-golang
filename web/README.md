# web框架基本实现

## day1

想要实现web框架，我们得先利用net/http包帮助我们搭建服务器我们可以直接使用
```go
package main

import "net/http"

func main() {
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("你好这是home"))
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("你好这是root"))
	})
	http.ListenAndServe(":9999", nil)
}

```
以下代码可以实现简单服务器，通过HandleFunc给地址绑定handle然后通过ListenAndServe去开启服务

http.ListenAndServe中第二个参数可以传一个handle，接口规范如下
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```
因此实现 ServeHTTP方法就可以实现封装
```go
package main

import "net/http"

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte("这是root"))
	case "/home":
		w.Write([]byte("这里是home"))
	default:
		w.Write([]byte("404 NOT FOUND"))
	}
}

func main() {
	engine := new(Engine)
	http.ListenAndServe(":9999", engine)
}
```

然后我们根据这个就可以构建出gin框架的雏形，我们要实现new一个engine然后通过get post等方法添加路由，我们把自己的框架叫做gee然后新建一个文件夹编写

```go
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

```
main
```go
package main

import (
	"github.com/2281501756/7days-golang/web/gee"
	"net/http"
)

func main() {
	engin := gee.New()
	engin.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("这里是root"))
	})
	engin.GET("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("这里是home"))
	})
	engin.Run(":9999")
}

```
封装出gee框架的第一版，实现了框架的原型路由映射







