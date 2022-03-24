package gee

import (
	"net/http"
	"strings"
)

type HandleFunc func(c *Context)

type router struct {
	roots    map[string]*node
	handlers map[string]HandleFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func NewRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandleFunc),
	}
}

func (r router) handle(w http.ResponseWriter, req *http.Request) {
	// build Context
	c := NewContext(w, req)
	n, params := r.GetRoute(c.Method, c.Path)
	if n != nil {
		c.PathParams = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.WriteString(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

func (r *router) AddRoute(pattern string, method string, handler HandleFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) GetRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
