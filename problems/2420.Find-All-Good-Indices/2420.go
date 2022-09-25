package p2420

func goodIndices(nums []int, k int) []int {
	increasingDiff := make([]int, len(nums))
	decreasingDiff := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] >= 0 {
			increasingDiff[i] = increasingDiff[i-1] + 1
		}

		if nums[i]-nums[i-1] <= 0 {
			decreasingDiff[i] = decreasingDiff[i-1] + 1
		}
	}

	ans := []int{}
	for i := k; i < len(nums)-k; i++ {
		if decreasingDiff[i-1]-decreasingDiff[i-k] == k-1 && increasingDiff[i+k]-increasingDiff[i+1] == k-1 {
			ans = append(ans, i)
		}
	}
	return ans
}
