package trie

type Trie[T any] struct {
	Root *Node[T]
}

func (t *Trie[T]) Put(key string, value T) {
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
			if n == len(key)-1 {
				root.Children[nodePrefix] = NodeWithValue[T](value)
				t.Root = newRoot
				return
			}
			preNode, ok := root.Children[nodePrefix]
			if ok == false {
				next = NewNode[T](map[string]*Node[T]{})
			} else {
				next = NodeWithValueAndChildren[T](preNode.Children, preNode.Value)
			}
			root.Children[nodePrefix] = next
		}
		root = next
	}
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
