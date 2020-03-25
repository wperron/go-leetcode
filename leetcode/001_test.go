package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	in := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	actual := TwoSum(in, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error: expected [ %v ] got [ %v ]\n", expected, actual)
	}
}
