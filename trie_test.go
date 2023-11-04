package trie

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicPutTest(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("h1", "world")
	trie = trie.Put("h12", "world")
	Equal(t, trie.Get("h1"), "world")
	Equal(t, trie.Get("h12"), "world")
}

func TestTrieStructureCheck(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("test", "233")
	Equal(t, trie.Get("test"), "233")

	Len(t, trie.Root.Children, 1)
	Len(t, trie.Root.Children["t"].Children, 1)
	Len(t, trie.Root.Children["t"].Children["e"].Children, 1)
	Len(t, trie.Root.Children["t"].Children["e"].Children["s"].Children, 1)
	Len(t, trie.Root.Children["t"].Children["e"].Children["s"].Children["t"].Children, 0)
	True(t, trie.Root.Children["t"].Children["e"].Children["s"].Children["t"].IsValueNode)
}

func TestBasicPutGetTest(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("test", "233")
	Equal(t, trie.Get("test"), "233")

	trie = trie.Put("test", "23333333")
	Equal(t, trie.Get("test"), "23333333")

	Equal(t, trie.Get("test-2333"), "")

	trie = trie.Put("", "empty-key")
	Equal(t, trie.Get(""), "empty-key")
}

func TestPutGetOnePath(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("111", "111")
	trie = trie.Put("11", "11")
	trie = trie.Put("1111", "1111")
	trie = trie.Put("11", "22")

	Equal(t, trie.Get("11"), "22")
	Equal(t, trie.Get("111"), "111")
	Equal(t, trie.Get("1111"), "1111")
}

func TestBasicRemoveTest1(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("test", "2333")
	Equal(t, trie.Get("test"), "2333")

	trie = trie.Put("te", "23")
	Equal(t, trie.Get("te"), "23")

	trie = trie.Put("tes", "233")
	Equal(t, trie.Get("tes"), "233")

	trie = trie.Remove("test")
	trie = trie.Remove("tes")
	trie = trie.Remove("te")

	Equal(t, trie.Get("te"), "")
	Equal(t, trie.Get("tes"), "")
	Equal(t, trie.Get("test"), "")
}

func TestBasicRemoveTest2(t *testing.T) {
	trie := NewTrie[string]()
	// Put something
	trie = trie.Put("test", "2333")
	Equal(t, trie.Get("test"), "2333")
	trie = trie.Put("te", "23")
	Equal(t, trie.Get("te"), "23")
	trie = trie.Put("tes", "233")
	Equal(t, trie.Get("tes"), "233")
	trie = trie.Put("", "123")
	Equal(t, trie.Get(""), "123")
	// Delete something
	trie = trie.Remove("")
	trie = trie.Remove("te")
	trie = trie.Remove("tes")
	trie = trie.Remove("test")

	Equal(t, trie.Get(""), "")
	Equal(t, trie.Get("te"), "")
	Equal(t, trie.Get("tes"), "")
	Equal(t, trie.Get("test"), "")
}

func TestRemoveFreeTest(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("test", "2333")
	trie = trie.Put("te", "23")
	trie = trie.Put("tes", "233")
	trie = trie.Remove("tes")
	trie = trie.Remove("test")
	Len(t, trie.Root.Children["t"].Children["e"].Children, 0)
	trie = trie.Remove("te")

	Nil(t, trie.Root)
}

func TestCopyOnWriteTest1(t *testing.T) {
	trie := NewTrie[string]()
	// Put something
	trie1 := trie.Put("test", "2333")
	trie2 := trie1.Put("te", "23")
	trie3 := trie2.Put("tes", "233")

	// Delete something
	trie4 := trie3.Remove("te")
	trie5 := trie3.Remove("tes")
	trie6 := trie3.Remove("test")

	// Check each snapshot
	Equal(t, trie3.Get("te"), "23")
	Equal(t, trie3.Get("tes"), "233")
	Equal(t, trie3.Get("test"), "2333")

	Equal(t, trie4.Get("te"), "")
	Equal(t, trie4.Get("tes"), "233")
	Equal(t, trie4.Get("test"), "2333")

	Equal(t, trie5.Get("te"), "23")
	Equal(t, trie5.Get("tes"), "")
	Equal(t, trie5.Get("test"), "2333")

	Equal(t, trie6.Get("te"), "23")
	Equal(t, trie6.Get("tes"), "233")
	Equal(t, trie6.Get("test"), "")
}
