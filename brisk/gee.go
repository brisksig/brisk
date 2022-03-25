package brisk

import (
	"fmt"
	"net/http"
	"strings"
)

type Brisk struct {
	Router *Router
}

func New() *Brisk {
	return &Brisk{Router: NewRouter()}
}

func (g *Brisk) Post(pattern string, handler HandleFunc) {
	g.Router.Add(pattern, http.MethodPost, handler)
}

func (g *Brisk) Get(pattern string, handler HandleFunc) {
	g.Router.Add(pattern, http.MethodGet, handler)
}

func (g *Brisk) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 1、上下文构建
	c := NewContext(w, req)
	// 2、路由分发
	method := c.Method
	pattern := c.Path
	handler, err := g.Router.Dispatch(method, pattern, c)
	if err != nil {
		c.WriteString(http.StatusNotFound, "404 Not Found")
	} else {
		handler(c)
	}
}

func (g *Brisk) Run(addr string) (err error) {
	fmt.Printf("server running····\n")
	if strings.HasPrefix(addr, ":") {
		fmt.Printf("listen on http://localhost%s\n", addr)
	} else {
		fmt.Printf("listen on http://%s\n", addr)
	}
	return http.ListenAndServe(addr, g)
}
