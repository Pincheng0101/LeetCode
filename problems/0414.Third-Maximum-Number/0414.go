package p0414

import "sort"

// 1. 先由大排到小，再來找第三最大的數值
// Runtime: 4 ms, Memory Usage: 3 MB
func thirdMax_1(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	if len(nums) < 3 {
		return nums[0]
	}
	distinctMaximumCount := 0
	diffValue := 0
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] != diffValue {
			diffValue = nums[i]
			distinctMaximumCount++
			if distinctMaximumCount == 3 {
				return nums[i]
			}
		}
	}
	return nums[0]
}
