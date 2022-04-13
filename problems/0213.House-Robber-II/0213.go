package p0213

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// compare if nums[0] be robbery with if nums[1] be robbery
	return max(robI(nums[:len(nums)-1]), robI(nums[1:]))
}

func robI(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := [3]int{}
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[2] + dp[0]
	for i := 3; i < len(nums); i++ {
		dp[0], dp[1], dp[2] = dp[1], dp[2], max(dp[0]+nums[i], dp[1]+nums[i])
	}
	return max(dp[1], dp[2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
