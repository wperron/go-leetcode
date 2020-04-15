package leetcode

import (
	"reflect"
	"testing"

	"github.com/wperron/go-leetcode/pkg/linkedlist"
)

func TestSwapPairs(t *testing.T) {
	// test list with even number of elements
	n1 := linkedlist.Node{Val: 1, Next: nil}
	n2 := linkedlist.Node{Val: 2, Next: nil}
	n3 := linkedlist.Node{Val: 3, Next: nil}
	n4 := linkedlist.Node{Val: 4, Next: nil}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4

	res := SwapPairs(&n1)
	last := res
	var actual []int
	expected := []int{2, 1, 4, 3}
	for last != nil {
		actual = append(actual, last.Val)
		last = last.Next
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error: expected %v for input %v, got %v", expected, n1, actual)
	}

	// test list with odd number of elements
	n5 := linkedlist.Node{Val: 5, Next: nil}
	n4.Next = &n5
	actual = []int{}
	expected = []int{2, 1, 4, 3, 5}
	res = SwapPairs(&n1)
	last = res
	for last != nil {
		actual = append(actual, last.Val)
		last = last.Next
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error: expected %v for input %v, got %v", expected, n1, actual)
	}
}
