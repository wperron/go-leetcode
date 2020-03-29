package leetcode

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	var input string
	var actual, expected int
	input = "abcabcbb"
	expected = 3
	actual = LongestSubstring(input)
	if actual != expected {
		t.Errorf("Error: expected %d for input '%s', got %d", expected, input, actual)
	}
}

func TestLongestSubstringEnd(t *testing.T) {
	var input string
	var actual, expected int
	input = "pwwkew"
	expected = 3
	actual = LongestSubstring(input)
	if actual != expected {
		t.Errorf("Error: expected %d for input '%s', got %d", expected, input, actual)
	}
}
