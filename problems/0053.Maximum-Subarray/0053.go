package p0053

// Kadane's Algorithm
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	bestSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		bestSum = max(bestSum, currentSum)
	}
	return bestSum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
