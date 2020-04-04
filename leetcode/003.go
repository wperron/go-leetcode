package leetcode

func LongestSubstring(input string) (longest int) {
	seen := make(map[byte]int)
	for i, j := 0, 0; i < len(input); i++ {
		_, found := seen[input[i]]
		if found {
			j = intmax(j, seen[input[i]]+1)
		}
		seen[input[i]] = i
		longest = intmax(longest, i-j+1)
	}
	return
}

func intmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
