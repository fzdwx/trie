package trie

type Trie[T any] struct {
	Root *Node[T]
}

func (t *Trie[T]) Remove(key string) *Trie[T] {
	if t.Root == nil {
		return t
	}

	curr := t.Root
	newRoot := curr

	if len(key) == 0 {
		return NewTrieWithRoot(NewNode(t.Root.Children))
	}

	var nodes []*Node[T]
	for n, char := range key {
		nodePrefix := string(char)
		node, ok := curr.Children[nodePrefix]
		if ok == false {
			return t
		}

		if islast(key, n) {
			if len(node.Children) == 0 {
				delete(curr.Children, nodePrefix)
				if len(curr.Children) == 0 {
					curr.IsValueNode = false
				}
			} else {

				curr.Children[nodePrefix] = NewNode(node.Children)
			}
		}
		nodes = append(nodes, curr)
		curr = node
	}

	for i := len(nodes) - 1; i >= 0; i-- {
		//for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		for k, child := range node.Children {
			if child.IsValueNode == false && len(child.Children) == 0 {
				delete(node.Children, k)
			}
		}
	}

	if newRoot.IsValueNode == false && len(newRoot.Children) == 0 {
		newRoot = nil
	}

	return NewTrieWithRoot[T](newRoot)
}

func (t *Trie[T]) Put(key string, value T) *Trie[T] {
	if key == "" {
		t.Root = NodeWithValueAndChildren(t.Root.Children, value)
	}

	var (
		curr    = t.Root
		newRoot = curr
	)

	for n, char := range key {
		var (
			next       *Node[T]
			nodePrefix = string(char)
		)

		if curr == nil {
			next = NewNode[T](map[string]*Node[T]{})
			curr = NewNode(map[string]*Node[T]{nodePrefix: next})
			newRoot = curr
		} else {
			preNode, ok := curr.Children[nodePrefix]
			if islast(key, n) {
				if ok == false {
					next = NodeWithValue[T](value)
				} else {
					next = NodeWithValueAndChildren[T](preNode.Children, value)
				}
			} else {
				if ok == false {
					next = NewNode[T](map[string]*Node[T]{})
				} else {
					next = NodeWithValueAndChildren[T](preNode.Children, preNode.Value)
				}
			}
			curr.Children[nodePrefix] = next
		}
		curr = next
	}

	return NewTrieWithRoot[T](newRoot)
}

func (t *Trie[T]) Get(key string) T {
	var value T
	root := t.Root
	if root == nil {
		return value
	}

	for _, char := range key {
		node, ok := root.Children[string(char)]
		if ok == false {
			return value
		}
		root = node
	}

	return root.Value
}

type Node[T any] struct {
	IsValueNode bool
	Children    map[string]*Node[T]
	Value       T
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		Root: nil,
	}
}

func NewTrieWithRoot[T any](root *Node[T]) *Trie[T] {
	return &Trie[T]{
		Root: root,
	}
}

func NewNode[T any](children map[string]*Node[T]) *Node[T] {
	return &Node[T]{
		IsValueNode: false,
		Children:    children,
	}
}

func NodeWithValue[T any](val T) *Node[T] {
	return &Node[T]{
		IsValueNode: true,
		Value:       val,
		Children:    map[string]*Node[T]{},
	}
}

func NodeWithValueAndChildren[T any](children map[string]*Node[T], val T) *Node[T] {
	return &Node[T]{
		IsValueNode: true,
		Value:       val,
		Children:    children,
	}
}

func islast(key string, n int) bool {
	return n == len(key)-1
}
