package main

import (
	"github.com/2281501756/7days-golang/web/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>你好这里是首页</h1>")
	})
	r.GET("/home", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"data": "这里是home",
			"code": 200,
		})
	})
	user := r.Group("/user")
	user.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1 style='color: red'>用户分组</h1>")
	})

	user.GET("/:id", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"id":   c.Param("id"),
			"user": "user",
		})
	})
	user.GET("/:id/:name", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"params": c.Params,
		})
	})
	r.GET("/static/*filename", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"params": c.Params,
		})
	})
	r.Run(":9999")
}
