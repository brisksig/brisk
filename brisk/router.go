package brisk

import (
	"errors"
	"strings"
)

type Router struct {
	MethodTree map[string]*Trie
}

func NewRouter() *Router {
	return &Router{MethodTree: map[string]*Trie{}}
}

func (r *Router) Add(pattern string, method string, handler HandleFunc) {
	methodtree, ok := r.MethodTree[method]
	if !ok {
		r.MethodTree[method] = NewTrie()
		r.MethodTree[method].Insert(pattern, handler)
	} else {
		methodtree.Insert(pattern, handler)
	}
}

func (r *Router) Dispatch(method string, pattern string, c *Context) (HandleFunc, error) {
	methodtree, ok := r.MethodTree[method]
	if !ok {
		return nil, errors.New("404 NotFound")
	}
	handler, pathparams, err := methodtree.Search(pattern)
	c.PathParams = pathparams
	return handler, err
}

type Node struct {
	isTail    bool
	isDynamic bool
	handler   HandleFunc
	next      map[string]*Node
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{next: map[string]*Node{}},
	}
}

func parsePattern(pattern string) []string {
	pattern = strings.Trim(pattern, "/")
	parts := strings.Split(pattern, "/")
	for idx, val := range parts {
		parts[idx] = val + "/"
	}
	return parts
}

func validDynamic(part string) bool {
	return strings.HasPrefix(part, ":")
}

func (t *Trie) Insert(pattern string, handler HandleFunc) {
	parts := parsePattern(pattern)
	cur := t.root
	for _, part := range parts {
		if _, ok := cur.next[part]; !ok {
			newnode := &Node{next: map[string]*Node{}}
			if validDynamic(part) {
				for _, node := range cur.next {
					if node.isDynamic {
						err := errors.New("YOU CANNOT ADD THE SAME DYNAMIC RULE TO THE SAME ROUTE")
						panic(err)
					}
				}
				newnode.isDynamic = true
			} else {
				newnode.isDynamic = false
			}
			cur.next[part] = newnode
		}
		cur = cur.next[part]
	}
	cur.isTail = true
	cur.handler = handler
}

func (t *Trie) Search(pattern string) (HandleFunc, map[string]string, error) {
	// make pathparamsmap
	pathparams := make(map[string]string)
	parts := parsePattern(pattern)
	cur := t.root
	for _, part := range parts {
		if _, ok := cur.next[part]; !ok {
			found_dynamic := false
			for keypart, node := range cur.next {
				if node.isDynamic {
					found_dynamic = true
					partpath := strings.TrimLeft(keypart, ":")
					partpath = strings.TrimRight(partpath, "/")
					pathparams[partpath] = strings.TrimRight(part, "/")
					part = keypart
				}
			}
			if !found_dynamic {
				return nil, pathparams, errors.New("404 NotFound")
			}
		}
		cur = cur.next[part]
	}
	return cur.handler, pathparams, nil
}
