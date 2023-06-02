package main

import (
	"fmt"
	"github.com/2281501756/7days-golang/web/gee"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>你好这里是首页</h1>")
	})
	r.GET("/home", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"data": "这里是home",
			"code": 200,
		})
	})
	userMiddleware := func() gee.HandleFunc {
		return func(c *gee.Context) {
			fmt.Println("请求了user")
			c.Next()
		}
	}
	user := r.Group("/user")
	user.Static("/static", "./img")
	user.Use(userMiddleware())
	user.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1 style='color: red'>用户分组</h1>")
	})

	user.GET("/:id", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"id":   c.Param("id"),
			"user": "user",
		})
	})

	r.Run(":9999")
}
