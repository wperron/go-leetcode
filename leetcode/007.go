package leetcode

import (
	"math"
)

func Reverse(x int) int {
	// simply return single digit inputs
	if x < 9 && x > -9 {
		return x
	}

	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	var answer int
	for x > 0 {
		tmp := x % 10
		answer = answer*10 + tmp
		x /= 10
	}

	answer = sign * answer
	if answer > math.MaxInt32 || answer < math.MinInt32 {
		return 0
	}
	return answer
}
