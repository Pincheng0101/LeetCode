package p2161

func pivotArray(nums []int, pivot int) []int {
	ans := []int{}
	for _, v := range nums {
		if v < pivot {
			ans = append(ans, v)
		}
	}
	for _, v := range nums {
		if v == pivot {
			ans = append(ans, v)
		}
	}
	for _, v := range nums {
		if v > pivot {
			ans = append(ans, v)
		}
	}
	return ans
}
