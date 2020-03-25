package leetcode

import (
	"github.com/wperron/go-leetcode/pkg/linkedlist"
)

func AddTwoNumbers(l1 *linkedlist.Node, l2 *linkedlist.Node) *linkedlist.Node {
	res := linkedlist.Node{0, nil}
	lastres := &res
	lastl1 := l1
	lastl2 := l2
	var carry, this int

	// Add numbers of each list with the other at the same position
	for lastl1 != nil && lastl2 != nil {
		this, carry = add(lastl1.Val, lastl2.Val+carry)
		newres := linkedlist.Node{this, nil}
		lastres.Next = &newres
		lastres = &newres

		lastl1 = lastl1.Next
		lastl2 = lastl2.Next
	}

	// if one of the list still has elements, add them to the result,
	// taking into account remaining carries
	if lastl1 != nil {
		for lastl1 != nil {
			this, carry = add(lastl1.Val, carry)
			newres := linkedlist.Node{this, nil}
			lastres.Next = &newres
			lastres = &newres
			lastl1 = lastl1.Next
		}
	} else if lastl2 != nil {
		for lastl2 != nil {
			this, carry = add(lastl2.Val, carry)
			newres := linkedlist.Node{this, nil}
			lastres.Next = &newres
			lastres = &newres
			lastl2 = lastl2.Next
		}
	}

	// append any remaining carry to the list
	if carry != 0 {
		last := linkedlist.Node{carry, nil}
		lastres.Next = &last
	}

	return res.Next
}

// adds two single digit integer and returns a single digit integer and the carry
func add(a, b int) (answer, carry int) {
	answer = a + b
	if answer > 9 {
		carry = answer / 10
		answer = answer % 10
	}
	return
}
