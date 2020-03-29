package leetcode

import (
	"testing"
)

func TestArrayMedian(t *testing.T) {
	var inputs [2][]int
	var actual, expected float64
	inputs = [2][]int{
		[]int{1, 3},
		[]int{2},
	}
	expected = 2.0
	actual = ArrayMedian(inputs[0], inputs[1])
	if actual != expected {
		t.Errorf("Error: expected %f for inputs %v and %v, got %f", expected, inputs[0], inputs[1], actual)
	}
}
