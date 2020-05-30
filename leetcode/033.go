package leetcode

func search(nums []int, target int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
