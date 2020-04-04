package leetcode

func isPalindrome(x int) bool {
	reversed := 0
	xcopy := x
	for xcopy > 0 {
		tmp := xcopy % 10
		reversed = reversed*10 + tmp
		xcopy /= 10
	}

	return x == reversed
}
