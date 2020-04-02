package leetcode

func LongestPalindromicSubstring(s string) string {
	if len(s) == 1 {
		return s
	}
	if IsPalindrome(s) {
		return s
	}

	i, max, start := 0, 0, 0
	for i < len(s) {
		left, right := i, i
		for right < len(s)-1 && s[right] == s[right+1] {
			right++
		}
		i = right + 1
		for right < len(s)-1 && left > 0 && s[right+1] == s[left-1] {
			right++
			left--
		}
		curr := right - left + 1
		if curr > max {
			start = left
			max = curr
		}
	}
	return s[start : start+max]
}

func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
