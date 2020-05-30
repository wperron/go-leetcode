package leetcode

import (
	"reflect"
	"testing"
)

type Case027 struct {
	Input    Input
	Actual   int
	Expected int
}

type Input struct {
	Nums []int
	Val  int
}

func TestRemove(t *testing.T) {
	c := Case027{
		Input:    Input{Nums: []int{3, 2, 2, 3}, Val: 3},
		Expected: 2,
	}

	c.Actual = removeElement(c.Input.Nums, c.Input.Val)
	if !reflect.DeepEqual(c.Input.Nums[:c.Actual], []int{2, 2}) {
		t.Errorf("Error: expected something else to happen")
	}
}
