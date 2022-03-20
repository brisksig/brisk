package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Router map[string]HandleFunc

func (r Router) Add(path string, method string, handler HandleFunc) {
	key := path + "-" + method
	r[key] = handler
}

func (r Router) New() Router {
	return make(Router)
}

func (r Router) handle(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path + "-" + req.Method
	if handler, ok := r[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 Not Found %s", req.URL.Path)
	}
}
