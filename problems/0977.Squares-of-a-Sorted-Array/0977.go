package p0977

import "sort"

// 1. 用最簡潔的方法實現，不過輸入是有序的，可以針對有序的方式去想
func sortedSquares_1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}

// 2. 從左右兩邊判斷大小依序填入
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func sortedSquares_2(nums []int) []int {
	results := make([]int, len(nums))
	left_index, right_index := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[left_index]) >= abs(nums[right_index]) {
			results[i] = nums[left_index] * nums[left_index]
			left_index++
		} else {
			results[i] = nums[right_index] * nums[right_index]
			right_index--
		}
	}
	return results
}
