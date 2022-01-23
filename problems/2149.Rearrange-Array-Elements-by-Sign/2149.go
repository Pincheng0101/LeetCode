package p2149

func rearrangeArray_1(nums []int) []int {
	result := make([]int, len(nums))
	for i, j := 0, 0; i < len(nums); i += 2 {
		for ; nums[j] < 0; j++ {
		}
		result[i] = nums[j]
		j++
	}
	for i, j := 1, 0; i < len(nums); i += 2 {
		for ; nums[j] > 0; j++ {
		}
		result[i] = nums[j]
		j++
	}
	return result
}

func rearrangeArray_2(nums []int) []int {
	result := make([]int, len(nums))
	for i, pos, neg := 0, 0, 0; i < len(nums); i, pos, neg = i+2, pos+1, neg+1 {
		for ; nums[pos] < 0; pos++ {
		}
		for ; nums[neg] > 0; neg++ {
		}
		result[i], result[i+1] = nums[pos], nums[neg]
	}
	return result
}
