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

func (b *Brisk) Post(pattern string, handler HandleFunc) {
	b.Router.Add(pattern, http.MethodPost, handler)
}

func (b *Brisk) Get(pattern string, handler HandleFunc) {
	b.Router.Add(pattern, http.MethodGet, handler)
}

func (b *Brisk) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 1、上下文构建
	c := NewContext(w, req)
	// 2、路由分发
	method := c.Method
	pattern := c.Path
	handler, err := b.Router.Dispatch(method, pattern, c)
	if err != nil {
		c.WriteString(http.StatusNotFound, "404 Not Found")
	} else {
		// 3、前置中间件
		for _, middleware := range b.Router.Middlewares {
			middleware.process_request(c)
		}
		handler(c)
		// 4、后置中间件
		for _, middleware := range b.Router.Middlewares {
			middleware.process_response(c)
		}
	}

}

func (b *Brisk) Run(addr string) (err error) {
	fmt.Printf("server running····\n")
	if strings.HasPrefix(addr, ":") {
		fmt.Printf("listen on http://localhost%s\n", addr)
	} else {
		fmt.Printf("listen on http://%s\n", addr)
	}
	return http.ListenAndServe(addr, b)
}
