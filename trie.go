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

	for n, char := range key {
		nodePrefix := string(char)
		node, ok := curr.Children[nodePrefix]
		if ok == false {
			return t
		}

		if islast(key, n) {
			if len(node.Children) == 0 {
				delete(curr.Children, nodePrefix)
			} else {
				curr.Children[nodePrefix] = NewNode(node.Children)
			}
		}
		curr = node
	}

	return NewTrieWithRoot[T](newRoot)
}

func (t *Trie[T]) Put(key string, value T) *Trie[T] {
	if key == "" {
		t.Root = NodeWithValueAndChildren(t.Root.Children, value)
	}

	curr := t.Root
	newRoot := curr
	for n, char := range key {
		var next *Node[T]
		nodePrefix := string(char)
		if curr == nil {
			next = NewNode[T](map[string]*Node[T]{})
			m := map[string]*Node[T]{
				nodePrefix: next,
			}
			curr = NewNode(m)
			newRoot = curr
		} else {
			preNode, ok := curr.Children[nodePrefix]
			if islast(key, n) {
				if ok {
					curr.Children[nodePrefix] = NodeWithValueAndChildren[T](preNode.Children, value)
				} else {
					curr.Children[nodePrefix] = NodeWithValue[T](value)
				}
				break
			}

			if ok == false {
				next = NewNode[T](map[string]*Node[T]{})
			} else {
				next = NodeWithValueAndChildren[T](preNode.Children, preNode.Value)
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
