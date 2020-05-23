package leetcode

import (
	"reflect"
	"testing"
)

type Case struct {
	Input    []int
	Actual   [][]int
	Expected [][]int
}

func errorAndFail(c Case, t *testing.T) {
	t.Errorf("Error: expected %v for input %v, got %v. \n", c.Expected, c.Input, c.Actual)
	t.FailNow()
}

func TestThreeSums(t *testing.T) {
	c := Case{
		Input: []int{-1, 0, 1, 2, -1, -4},
		Expected: [][]int{
			{-1, -1, 2}, {-1, 0, 1},
		},
	}

	c.Actual = threeSums(c.Input)
	if !reflect.DeepEqual(c.Actual, c.Expected) {
		errorAndFail(c, t)
	}
}
