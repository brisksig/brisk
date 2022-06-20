package brisk

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTrie(t *testing.T) {
	// build trie
	trie := NewTrie()
	// assert node map
	assert.Equal(t, &Node{next: map[string]*Node{}, supportMethods: make([]string, 0)}, trie.root)
}

func TestNewRouter(t *testing.T) {
	// build router instance
	router := NewRouter()
	// assert trie
	assert.Equal(t, NewTrie(), router.tree)
}

func TestInsertWithHandler(t *testing.T) {
	// build a trie
	trie := NewTrie()
	handler := func(c *Context) {}
	trie.Insert("/", http.MethodGet, handler)
	// assert.Equal(``)
	// assert root != tail
	assert.Equal(t, false, trie.root.isTail)
	expected_next := make(map[string]*Node)
	expected_supportMethods := []string{http.MethodGet}
	expected_node := &Node{
		isTail:         true,
		isDynamic:      false,
		handler:        handler,
		next:           expected_next,
		supportMethods: expected_supportMethods,
	}
	// assert root.next's length == 1
	assert.Equal(t, 1, len(trie.root.next))
	actual_node := trie.root.next["/"]
	// assert root.next["/"] == Node{""}
	// assert root.next["/"].isTail is true
	// assert root.next["/"].isDynamic is false
	assert.Equal(t, expected_node.isTail, actual_node.isTail)
	assert.Equal(t, expected_node.isDynamic, actual_node.isDynamic)
	assert.Equal(t, expected_node.next, actual_node.next)
}
