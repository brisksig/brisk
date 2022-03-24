package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type Gee struct {
	// router map
	// example "/" -> IndexHandleFunc
	Router *router
}

func New() *Gee {
	return &Gee{Router: NewRouter()}
}

func (g *Gee) Post(pattern string, handler HandleFunc) {
	g.Router.AddRoute(pattern, http.MethodPost, handler)
}

func (g *Gee) Get(pattern string, handler HandleFunc) {
	g.Router.AddRoute(pattern, http.MethodGet, handler)
}

func (g *Gee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 路由分发
	g.Router.handle(w, req)
}

func (g *Gee) Run(addr string) (err error) {
	fmt.Printf("server running····\n")
	if strings.HasPrefix(addr, ":") {
		fmt.Printf("listen on http://localhost%s\n", addr)
	} else {
		fmt.Printf("listen on http://%s\n", addr)
	}
	return http.ListenAndServe(addr, g)
}
