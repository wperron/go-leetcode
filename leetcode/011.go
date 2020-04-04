package leetcode

func MaxArea(height []int) int {
	var max int
	i, j := 0, len(height)-1
	for i < j {
		max = intmax(max, intmin(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
