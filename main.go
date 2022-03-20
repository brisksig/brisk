package main

import (
	"fmt"
	"gee"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func HelloPost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "这是post请求的hello world")
}

func main() {
	g := gee.New()
	// 添加路由
	g.Router.Add("/", http.MethodGet, Hello)
	g.Post("/hello/", HelloPost)
	g.Run(":9999")
}
