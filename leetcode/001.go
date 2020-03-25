package leetcode

func TwoSum(nums []int, target int) []int {
	numlen := len(nums)
	for i := 0; i < numlen-1; i++ {
		for j := i + 1; j < numlen; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
