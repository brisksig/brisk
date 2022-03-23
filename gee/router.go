package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(c *Context)

type Router map[string]HandleFunc

func (r Router) Add(pattern string, method string, handler HandleFunc) {
	key := pattern + "-" + method
	r[key] = handler
}

func (r Router) New() Router {
	return make(Router)
}

func (r Router) handle(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path + "-" + req.Method
	// build Context
	c := NewContext(w, req)
	if handler, ok := r[key]; ok {
		handler(c)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 Not Found %s", req.URL.Path)
	}
}
