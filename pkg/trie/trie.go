package trie

type TrieNode struct {
	End   bool
	Links map[rune]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func InitTrieNode() *TrieNode {
	return &TrieNode{
		End:   false,
		Links: make(map[rune]*TrieNode),
	}
}

func InitTrie() *Trie {
	return &Trie{
		Root: InitTrieNode(),
	}
}

func (t *Trie) Insert(s string) {
	root := t.Root
	node := root
	for _, char := range s {
		if node.Links[char] == nil {
			node.Links[char] = InitTrieNode()
		}
		node = node.Links[char]
	}
	node.End = true
}
