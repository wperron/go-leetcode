package leetcode

func ArrayMedian(l1, l2 []int) (median float64) {
	merged := []int{}
	for len(l1) > 0 && len(l2) > 0 {
		if l1[0] < l2[0] {
			merged = append(merged, l1[0])
			l1 = l1[1:]
		} else {
			merged = append(merged, l2[0])
			l2 = l2[1:]
		}
	}

	if len(l1) > 0 {
		merged = append(merged, l1...)
	} else if len(l2) > 0 {
		merged = append(merged, l2...)
	}

	half := len(merged) / 2
	if len(merged) % 2 == 0 {
		median = (float64(merged[half-1]) + float64(merged[half])) / 2
	} else {
		median = float64(merged[half])
	}
	return
}
