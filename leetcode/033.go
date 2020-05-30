package leetcode

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := len(nums) / 2
	result := -1
	for low < mid {
		if target == nums[mid] {
			result = mid
			low, high = high, low
		} else if target >= nums[low] && target < nums[mid] {
			mid, high = low+((mid-low)/2), mid
		} else {
			low, mid = mid, mid+((high-mid)/2)
		}
	}
	return result
}
