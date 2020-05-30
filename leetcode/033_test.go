package leetcode

import (
	"testing"
)

func TestSearch(t *testing.T) {
	list := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	pos := search(list, target)
	if pos != 4 {
		t.Errorf("Error: expected position to be 4 for input %v and target %d, got %d\n", list, target, pos)
	}
}

func BenchmarkSearch(b *testing.B) {
	list := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	for i := 0; i < b.N; i++ {
		search(list, target)
	}
}
