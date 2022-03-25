package main

import (
	"brisk"
	"net/http"
)

func Hello(c *brisk.Context) {
	c.WriteString(http.StatusOK, "hello world")
}

func HelloPost(c *brisk.Context) {
	p := &Person{Name: "Yova"}
	c.WriteJSON(http.StatusOK, p)
}

func HelloHTML(c *brisk.Context) {
	html := "<h1>测试</h1>"
	id := c.PathParams["id"]
	c.HTML(http.StatusOK, html+id)
}

func HelloForm(c *brisk.Context) {
	a := c.Form("a")
	b := c.QueryParams("b")
	c.WriteString(http.StatusOK, "form=%s,query=%s", a, b)
}

type Person struct {
	Name string `json:"name"`
}

func HelloJSON(c *brisk.Context) {
	p := &Person{}
	err := c.JsonBind(p)
	if err != nil {
		c.WriteString(http.StatusOK, "有错误", err)
	}
	c.WriteJSON(http.StatusAccepted, p)
}

func main() {
	g := brisk.New()
	// 添加路由
	g.Router.Add("/", http.MethodGet, Hello)
	g.Post("/hello/index/", HelloPost)
	g.Get("/index/:id/", HelloHTML)
	g.Post("/form/", HelloForm)
	g.Post("/json/path/", HelloJSON)
	g.Run(":9999")
}
