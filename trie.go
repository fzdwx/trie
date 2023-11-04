package trie

type Trie[T any] struct {
	Root *Node[T]
}

func (t *Trie[T]) Put(key string, value T) *Trie[T] {
	if key == "" {
		t.Root = NodeWithValueAndChildren(t.Root.Children, value)
	}

	root := t.Root
	newRoot := root
	for n, char := range key {
		var next *Node[T]
		nodePrefix := string(char)
		if root == nil {
			next = NewNode[T](map[string]*Node[T]{})
			m := map[string]*Node[T]{
				nodePrefix: next,
			}
			root = NewNode(m)
			newRoot = root
		} else {
			preNode, ok := root.Children[nodePrefix]
			if n == len(key)-1 {
				if ok {
					root.Children[nodePrefix] = NodeWithValueAndChildren[T](preNode.Children, value)
				} else {
					root.Children[nodePrefix] = NodeWithValue[T](value)
				}
				break
			}
			if ok == false {
				next = NewNode[T](map[string]*Node[T]{})
			} else {
				next = NodeWithValueAndChildren[T](preNode.Children, preNode.Value)
			}
			root.Children[nodePrefix] = next
		}
		root = next
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
