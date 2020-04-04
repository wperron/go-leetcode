package leetcode

import (
	"testing"
)

func TestIntIsPalindrome(t *testing.T) {
	var input int
	var actual, expected bool

	input = 121
	expected = true
	actual = isPalindrome(input)
	if expected != actual {
		t.Errorf("Error: expected %t for input %d, got %t\n", expected, input, actual)
	}
}
