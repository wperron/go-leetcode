package leetcode

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	var input []int
	var expected, actual int

	input = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected = 49
	actual = MaxArea(input)
	if actual != expected {
		t.Errorf("Error: expected %d for input %v, got %d", expected, input, actual)
	}
}
