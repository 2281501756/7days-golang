package main

import (
	"github.com/2281501756/7days-golang/web/gee"
)

func main() {
	engin := gee.New()
	engin.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>你好这里是首页</h1>")
	})
	engin.GET("/home", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"data": "这里是home",
			"code": 200,
		})
	})
	engin.Run(":9999")
}
