package main

import (
	"gee"
	"net/http"
)

func Hello(c *gee.Context) {
	c.WriteString(http.StatusOK, "hello world")
}

func HelloPost(c *gee.Context) {
	p := &Person{Name: "Yova"}
	c.WriteJSON(http.StatusOK, p)
}

func HelloHTML(c *gee.Context) {
	html := "<h1>测试</h1>"
	id := c.PathParams.Get("id")
	c.HTML(http.StatusOK, html+id)
}

func HelloForm(c *gee.Context) {
	a := c.Form("a")
	b := c.QueryParams("b")
	c.WriteString(http.StatusOK, "form=%s,query=%s", a, b)
}

type Person struct {
	Name string `json:"name"`
}

func HelloJSON(c *gee.Context) {
	p := &Person{}
	err := c.JsonBind(p)
	if err != nil {
		c.WriteString(http.StatusOK, "有错误", err)
	}
	c.WriteJSON(http.StatusAccepted, p)
}

func main() {
	g := gee.New()
	// 添加路由
	g.Router.AddRoute("/", http.MethodGet, Hello)
	g.Post("/hello/index/", HelloPost)
	g.Get("/index/:id/", HelloHTML)
	g.Post("/form/", HelloForm)
	g.Post("/json/*path/", HelloJSON)
	g.Run(":9999")
}
