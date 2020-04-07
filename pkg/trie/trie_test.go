package trie

import (
	"testing"
)

func TestInsert(t *testing.T) {
	trie := InitTrie()
	trie.Insert("leetcode")
	trie.Insert("leetmus")
}

func TestCommonRoot(t *testing.T) {
	trie := InitTrie()
	trie.Insert("apple")
	trie.Insert("alice")

	if len(trie.Root.Links) != 1 {
		t.Errorf("Error: expected 'apple' and 'alice' to share the same root\n")
	}

	keys := make([]rune, 0, len(trie.Root.Links))
	for k := range trie.Root.Links {
		keys = append(keys, k)
	}

	if string(keys[0]) != "a" {
		t.Errorf("Error: expected first element of the 'apple' trie to be 'a', got %s", string(keys[0]))
	}
}
