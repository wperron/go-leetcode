package leetcode

import (
	"github.com/wperron/go-leetcode/pkg/linkedlist"
)

func SwapPairs(head *linkedlist.Node) *linkedlist.Node {
	newhead := linkedlist.Node{Val: -1, Next: nil}
	curr := &newhead
	i := head
	for i != nil {
		if i.Next == nil {
			// do something
			curr.Next = i
			i = i.Next
		} else {
			// do something else
			a, b := &linkedlist.Node{Val: i.Val, Next: nil}, &linkedlist.Node{Val: i.Next.Val, Next: nil}

			curr.Next = b
			curr.Next.Next = a

			i = i.Next.Next
			curr = curr.Next.Next
		}
	}
	return newhead.Next
}
