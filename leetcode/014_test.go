package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var input []string
	var expected, actual string

	input = []string{"flower", "flow", "flight"}
	expected = "fl"
	actual = LongestCommonPrefix(input)
	if actual != expected {
		t.Errorf("Error: expected %s for input %v, got %s\n", expected, input, actual)
	}
}

func TestWithEmptyString(t *testing.T) {
	var input []string
	var expected, actual string

	input = []string{"", "b"}
	expected = ""
	actual = LongestCommonPrefix(input)
	if actual != expected {
		t.Errorf("Error: expected %s for input %v, got %s\n", expected, input, actual)
	}
}

func TestWithRepeatingCharacters(t *testing.T) {
	var input []string
	var expected, actual string

	input = []string{"aa", "a"}
	expected = "a"
	actual = LongestCommonPrefix(input)
	if actual != expected {
		t.Errorf("Error: expected %s for input %v, got %s\n", expected, input, actual)
	}
}
