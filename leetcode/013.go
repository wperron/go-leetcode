package leetcode

import (
	"reflect"
	"sort"
)

func threeSums(nums []int) [][]int {
	var results [][]int
	cp := make([]int, len(nums))
	copy(cp, nums)
	sort.Ints(cp)

	l := len(cp)
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				if cp[i]+cp[j]+cp[k] == 0 {
					res := []int{cp[i], cp[j], cp[k]}
					sort.Ints(res)

					exists := false
					for _, v := range results {
						if reflect.DeepEqual(v, res) {
							exists = true
							break
						}
					}

					if !exists {
						results = append(results, res)
					}
				}
			}
		}
	}

	return results
}
