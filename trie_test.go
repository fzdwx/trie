package trie

import (
	"testing"
)

func TestBasicPutTest(t *testing.T) {
	trie := NewTrie[string]()
	trie.Put("h1", "world")
	trie.Put("h12", "world")
	assertEqual(t, trie.Get("h1"), "world")
	assertEqual(t, trie.Get("h12"), "world")
}

func TestTrieStructureCheck(t *testing.T) {
	trie := NewTrie[string]()
	trie.Put("test", "233")
	assertEqual(t, trie.Get("test"), "233")

	assertLen(t, trie.Root.Children, 1)
	assertLen(t, trie.Root.Children["t"].Children, 1)
	assertLen(t, trie.Root.Children["t"].Children["e"].Children, 1)
	assertLen(t, trie.Root.Children["t"].Children["e"].Children["s"].Children, 1)
	assertLen(t, trie.Root.Children["t"].Children["e"].Children["s"].Children["t"].Children, 0)
	assertTrue(t, trie.Root.Children["t"].Children["e"].Children["s"].Children["t"].IsValueNode)
}

func TestBasicPutGetTest(t *testing.T) {

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
