package gee

import (
	"fmt"
	"net/http"
	"strings"
)

type Gee struct {
	// router map
	// example "/" -> IndexHandleFunc
	Router     Router
	RouterTree *RouterTree
}

func New() *Gee {
	return &Gee{Router: make(Router)}
}

func (g *Gee) Post(pattern string, handler HandleFunc) {
	g.Router.Add(pattern, http.MethodPost, handler)
}

func (g *Gee) Get(pattern string, handler HandleFunc) {
	g.Router.Add(pattern, http.MethodGet, handler)
}

func (g *Gee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 路由分发
	g.Router.handle(w, req)
}

func (g *Gee) Run(addr string) (err error) {
	g.NewTree()
	fmt.Printf("server running····\n")
	if strings.HasPrefix(addr, ":") {
		fmt.Printf("http://localhost%s\n", addr)
	} else {
		fmt.Printf("http://%s\n", addr)
	}
	return http.ListenAndServe(addr, g)
}
