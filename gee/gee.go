package gee

import (
	"net/http"
)

type Gee struct {
	// router map
	// example "/" -> IndexHandleFunc
	Router Router
}

func New() *Gee {
	return &Gee{Router: make(Router)}
}

func (g *Gee) Post(path string, handler HandleFunc) {
	g.Router.Add(path, http.MethodPost, handler)
}

func (g *Gee) Get(path string, handler HandleFunc) {
	g.Router.Add(path, http.MethodGet, handler)
}

func (g *Gee) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 路由分发
	g.Router.handle(w, req)
}

func (g *Gee) Run(addr string) (err error) {
	return http.ListenAndServe(addr, g)
}
