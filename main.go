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
	b := brisk.New()
	// 添加路由
	b.Router.Add("/", http.MethodGet, Hello)
	b.Post("/hello/index/", HelloPost)
	b.Get("/index/:id/", HelloHTML)
	b.Post("/form/", HelloForm)
	b.Post("/json/path/", HelloJSON)
	// 子路由
	r := brisk.NewRouter()
	r.Add("/v1/", http.MethodGet, Hello)
	r.Add("/v2/", http.MethodPost, HelloJSON)
	b.Router.Include("/api/", r)
	// 中间件
	b.Router.Use(&brisk.LoggingMiddleware{})
	b.Run("0.0.0.0:8001")
}
