package leetcode

func LongestPalindromicSubstring(s string) string {
	// Handle the easy cases first. If string is only 1 character
	// long, it's clearly a palindrome. If the whole string is a
	// a palindrome, then it's obviously the longest one.
	if len(s) == 1 {
		return s
	}
	if IsPalindrome(s) {
		return s
	}

	// Looping through every letter. We'll initialize a `max` variable
	// to hold the length of the longest palindrome found so far and a
	// `start` variable to hold the starting index of the max palindrome.
	i, max, start := 0, 0, 0
	for i < len(s) {
		// We'll initialize two cursor at the current index; one will
		// move left, the other move right. The first inner loop will go
		// over duplicate characters so that they start at the 'edge'
		// so to speak.
		// example for "baaabad" if `i` is 1, we'll move the `right`
		// cursor at position 3. We'll also move the `i` to `right` + 1
		// again to skip series of continuous letters because they're
		// covered by that first inner loop.
		left, right := i, i
		for right < len(s)-1 && s[right] == s[right+1] {
			right++
		}
		i = right + 1
		
		// Look at the next character over for each cursor, if they're
		// the same, it means we're still looking at a palindrome, so
		// we keep going.
		for right < len(s)-1 && left > 0 && s[right+1] == s[left-1] {
			right++
			left--
		}
		
		// calculate the length of the current palindrome and save the
		// `max` and `start` if we've found a longer palindrome.
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
