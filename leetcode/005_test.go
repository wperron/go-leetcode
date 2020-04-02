package leetcode

import (
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	var input, actual, expected string

	input = "babad"
	expected = "bab"
	actual = LongestPalindromicSubstring(input)
	if actual != expected {
		t.Errorf("Error: expected %s for input %s, got %s\n", expected, input, actual)
	}
}

func TestIsPalindrome(t *testing.T) {
	var input string
	var actual, expected bool

	input = "laval"
	expected = true
	actual = IsPalindrome(input)
	if actual != expected {
		t.Errorf("Error: expected %t for input %s, got %t\n", expected, input, actual)
	}

	input = "foobar"
	expected = false
	actual = IsPalindrome(input)
	if actual != expected {
		t.Errorf("Error: expected %t for input %s, got %t\n", expected, input, actual)
	}
}
