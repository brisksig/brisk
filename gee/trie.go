package gee

import (
	"strings"
)

type MethodRootMap map[string]node

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isTail   bool    //是否叶子节点
	handler  HandleFunc
}

type RouterTree struct {
	MethodRootMap MethodRootMap
}

func createNode(current *node, part string) *node {
	if len(current.children) == 0 {
		new_child := &node{pattern: part}
		current.children = append(current.children, new_child)
		return new_child
	}
	for _, node := range current.children {
		if node.pattern == part {
			// 递归children
			return node
		}
	}
	new_child := &node{pattern: part}
	current.children = append(current.children, new_child)
	return new_child
}

func buildTreeByPath(path string, root *node, handler HandleFunc) {
	parts := strings.Split(path, "/")
	current := root
	var tail *node
	flag := false
	for _, part := range parts {
		part += "/"
		if part == "" {
			flag = true
			continue
		}
		tail = createNode(current, part)
	}
	if flag {
		return
	}
	tail.isTail = true
	tail.handler = handler
}

func (g *Gee) NewTree() {
	routermap := g.Router
	tree := &RouterTree{
		MethodRootMap: make(MethodRootMap),
	}
	for path, handler := range routermap {
		m_p := strings.Split(path, "-")
		method := m_p[1]
		path := m_p[0]
		methodRoot, ok := tree.MethodRootMap[method]
		if !ok {
			methodRoot = node{pattern: method}
			buildTreeByPath(path, &methodRoot, handler)
			tree.MethodRootMap[method] = methodRoot
		} else {
			buildTreeByPath(path, &methodRoot, handler)
			tree.MethodRootMap[method] = methodRoot
		}
	}
	g.RouterTree = tree
}

// func findNode(current *node, pattern string) (*node, error) {
// 	for _, node := range current.children {
// 		if node.pattern == pattern {
// 			if len(node.children) == 0 {
// 				// founded
// 				return node, nil
// 			} else {
// 				// 递归查找
// 				return findNode(node, pattern)
// 			}
// 		}
// 	}
// }

// // search handler
// func (g *Gee) SearchTree(path string, method string) (*HandleFunc, string) {
// 	patterns := strings.Split(path, "/")
// 	methodTree, ok := g.RouterTree.MethodRootMap[method]
// 	if !ok {
// 		return nil, fmt.Sprintf("不允许的请求方法%s", method)
// 	} else {
// 		current = methodTree.children
// 		pattern = patterns[0] + "/"
// 		for _,node := range(current) {
// 			if node.pattern == pattern {
// 				current =
// 			}
// 		}
// 		for _, pattern := range(patterns) {
// 			pattern += "/"
// 			if (current.pattern == pattern) {
// 				current = current.
// 			}
// 		}
// 	}
// }
