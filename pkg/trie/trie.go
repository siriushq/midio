// Package trie implements a simple trie tree for minio server/tools borrows
// idea from - https://godoc.org/golang.org/x/text/internal/triegen.
package trie

// Node trie tree node container carries value and children.
type Node struct {
	exists bool
	value  string
	child  map[rune]*Node // runes as child.
}

// newNode create a new trie node.
func newNode() *Node {
	return &Node{
		exists: false,
		value:  "",
		child:  make(map[rune]*Node),
	}
}

// Trie is a trie container.
type Trie struct {
	root *Node
	size int
}

// Root returns root node.
func (t *Trie) Root() *Node {
	return t.root
}

// Insert insert a key.
func (t *Trie) Insert(key string) {
	curNode := t.root
	for _, v := range key {
		if curNode.child[v] == nil {
			curNode.child[v] = newNode()
		}
		curNode = curNode.child[v]
	}

	if !curNode.exists {
		// increment when new rune child is added.
		t.size++
		curNode.exists = true
	}
	// value is stored for retrieval in future.
	curNode.value = key
}

// PrefixMatch - prefix match.
func (t *Trie) PrefixMatch(key string) []string {
	node, _ := t.findNode(key)
	if node == nil {
		return nil
	}
	return t.Walk(node)
}

// Walk the tree.
func (t *Trie) Walk(node *Node) (ret []string) {
	if node.exists {
		ret = append(ret, node.value)
	}
	for _, v := range node.child {
		ret = append(ret, t.Walk(v)...)
	}
	return
}

// find nodes corresponding to key.
func (t *Trie) findNode(key string) (node *Node, index int) {
	curNode := t.root
	f := false
	for k, v := range key {
		if f {
			index = k
			f = false
		}
		if curNode.child[v] == nil {
			return nil, index
		}
		curNode = curNode.child[v]
		if curNode.exists {
			f = true
		}
	}

	if curNode.exists {
		index = len(key)
	}

	return curNode, index
}

// NewTrie create a new trie.
func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
		size: 0,
	}
}
