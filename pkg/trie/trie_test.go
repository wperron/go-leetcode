package trie

import (
	"testing"
)

func TestInsert(t *testing.T) {
	trie := InitTrie()
	trie.Insert("leetcode")
	trie.Insert("leetmus")
}
