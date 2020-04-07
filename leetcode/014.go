package leetcode

import "github.com/wperron/go-leetcode/pkg/trie"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	trie := trie.InitTrie()
	for _, s := range strs {
		if len(s) == 0 {
			return ""
		}
		trie.Insert(s)
	}

	var max int
	node := trie.Root
	for len(node.Links) == 1 && !node.End {
		node = node.Links[rune(strs[0][max])]
		max++
	}

	return strs[0][0:max]
}
