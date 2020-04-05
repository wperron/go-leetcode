package leetcode

import (
	"testing"
)

func TestFib(t *testing.T) {
	var input, expected, actual int

	input = 4
	expected = 3
	actual = Fib(input)
	if expected != actual {
		t.Errorf("Error: expected %d for input %d, got %d\n", expected, input, actual)
	}
}
