package p0746

func minCostClimbingStairs(cost []int) int {
	dp := [2]int{}
	for i := 2; i <= len(cost); i++ {
		dp[0], dp[1] = dp[1], min(dp[0]+cost[i-2], dp[1]+cost[i-1])
	}
	return dp[1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
