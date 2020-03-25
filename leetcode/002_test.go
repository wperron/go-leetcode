package leetcode

import (
	"testing"

	"github.com/wperron/go-leetcode/pkg/linkedlist"
)

func TestAddTwoNumbers(t *testing.T) {
	n1 := linkedlist.Node{2, nil}
	n2 := linkedlist.Node{4, nil}
	n3 := linkedlist.Node{3, nil}
	n1.Next = &n2
	n2.Next = &n3

	n4 := linkedlist.Node{5, nil}
	n5 := linkedlist.Node{6, nil}
	n6 := linkedlist.Node{4, nil}
	n4.Next = &n5
	n5.Next = &n6

	newhead := AddTwoNumbers(&n1, &n4)
	last := newhead
	for i, v := range [3]int{7, 0, 8} {
		if last.Val != v {
			t.Errorf("Error: expected %d at pos %d, got %d", v, i, last.Val)
		}
		last = last.Next
	}
}

func TestAddTwoNumbersWithZero(t *testing.T) {
	n1 := linkedlist.Node{1, nil}
	n2 := linkedlist.Node{8, nil}
	n1.Next = &n2

	n3 := linkedlist.Node{0, nil}

	newhead := AddTwoNumbers(&n1, &n3)
	last := newhead
	for i, v := range [2]int{1, 8} {
		if last.Val != v {
			t.Errorf("Error: expected %d at pos %d, got %d", v, i, last.Val)
		}
		last = last.Next
	}
}
