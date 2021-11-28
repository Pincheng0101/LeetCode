package p0053

import "math"

// Kadane's Algorithm
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	bestSum, currentSum := math.MinInt32, math.MinInt32
	for _, v := range nums {
		currentSum = max(v, currentSum+v)
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
