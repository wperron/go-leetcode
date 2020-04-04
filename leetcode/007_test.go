package leetcode

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var input, expected, actual int

	input = 123
	expected = 321
	actual = Reverse(input)
	if expected != actual {
		t.Errorf("Error: expected %d for input %d, got %d", expected, input, actual)
	}

	input = 1534236469
	expected = 0
	actual = Reverse(input)
	if expected != actual {
		t.Errorf("Error: expected %d for input %d, got %d", expected, input, actual)
	}
}
