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
