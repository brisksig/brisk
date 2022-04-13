// Copyright 2022 DomineCore.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package brisk

import (
	"errors"
	"fmt"
	"strings"
)

type Router struct {
	tree        *Trie
	Middlewares []MiddleWare
}

func NewRouter() *Router {
	return &Router{tree: NewTrie()}
}

func (r *Router) Use(middleware MiddleWare) {
	r.Middlewares = append(r.Middlewares, middleware)
}

func (r *Router) Add(pattern string, method string, handler HandleFunc) {
	r.tree.Insert(pattern, method, handler)
}

func (r *Router) Include(pattern string, child *Router) {
	pattern = strings.TrimLeft(pattern, "/")
	r.tree.root.next[pattern] = child.tree.root
}

func (r *Router) Dispatch(method string, pattern string, c *Context) (HandleFunc, error) {
	handler, pathparams, err := r.tree.Search(pattern, c.Method)
	c.PathParams = pathparams
	return handler, err
}

type Node struct {
	isTail         bool
	isDynamic      bool
	handler        HandleFunc
	next           map[string]*Node
	supportMethods []string
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{next: map[string]*Node{}, supportMethods: make([]string, 0)},
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

func (t *Trie) Insert(pattern string, method string, handler HandleFunc) {
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
	cur.supportMethods = append(cur.supportMethods, method)
}

func (t *Trie) InsertChild(pattern string, child *Router) {
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
	cur.next = child.tree.root.next
}

func (t *Trie) searchNode(pattern string) (*Node, map[string]string, error) {
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
	return cur, pathparams, nil
}

func (t *Trie) Search(pattern string, method string) (HandleFunc, map[string]string, error) {
	node, pathparams, err := t.searchNode(pattern)
	if node != nil {
		// 请求方法合法性校验
		for _, support_method := range node.supportMethods {
			if support_method == method {
				return node.handler, pathparams, err
			}
		}
		method_err := fmt.Errorf("不受支持的方法%s", method)
		return nil, pathparams, method_err
	} else {
		return nil, pathparams, err
	}
}
