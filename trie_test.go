package trie

import (
	"testing"
)

func TestBasicPutTest(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("h1", "world")
	trie = trie.Put("h12", "world")
	assertEqual(t, trie.Get("h1"), "world")
	assertEqual(t, trie.Get("h12"), "world")
}

func TestTrieStructureCheck(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("test", "233")
	assertEqual(t, trie.Get("test"), "233")

	assertLen(t, trie.Root.Children, 1)
	assertLen(t, trie.Root.Children["t"].Children, 1)
	assertLen(t, trie.Root.Children["t"].Children["e"].Children, 1)
	assertLen(t, trie.Root.Children["t"].Children["e"].Children["s"].Children, 1)
	assertLen(t, trie.Root.Children["t"].Children["e"].Children["s"].Children["t"].Children, 0)
	assertTrue(t, trie.Root.Children["t"].Children["e"].Children["s"].Children["t"].IsValueNode)
}

func TestBasicPutGetTest(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("test", "233")
	assertEqual(t, trie.Get("test"), "233")

	trie = trie.Put("test", "23333333")
	assertEqual(t, trie.Get("test"), "23333333")

	assertEqual(t, trie.Get("test-2333"), "")

	trie = trie.Put("", "empty-key")
	assertEqual(t, trie.Get(""), "empty-key")
}

func TestPutGetOnePath(t *testing.T) {
	trie := NewTrie[string]()
	trie = trie.Put("111", "111")
	trie = trie.Put("11", "11")
	trie = trie.Put("1111", "1111")
	trie = trie.Put("11", "22")

	assertEqual(t, trie.Get("11"), "22")
	assertEqual(t, trie.Get("111"), "111")
	assertEqual(t, trie.Get("1111"), "1111")
}

func assertTrue(t *testing.T, node bool) {
	if node != true {
		t.Fatalf("expected %t, got %t", true, node)
	}
}

func size(m map[string]*Node[string]) int {
	return len(m)
}

func assertEqual(t *testing.T, s2, s1 string) {
	if s1 != s2 {
		t.Fatalf("expected %s, got %s", s1, s2)
	}
}

func assertLen(t *testing.T, s2 map[string]*Node[string], s1 int) {
	if s1 != size(s2) {
		t.Fatalf("expected %d, got %d", s1, size(s2))
	}
}
