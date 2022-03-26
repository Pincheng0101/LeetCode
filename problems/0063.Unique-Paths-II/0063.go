package p0063

// 由上而下
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	mLen, nLen := len(obstacleGrid), len(obstacleGrid[0])
	cache := make([][]int, mLen)
	for i := 0; i < mLen; i++ {
		cache[i] = make([]int, nLen)
	}
	var dp func(int, int) int
	dp = func(m, n int) int {
		if obstacleGrid[mLen-m-1][nLen-n-1] == 1 {
			return 0
		}
		if m == 0 && n == 0 {
			return 1
		}
		cnt := 0
		if m > 0 {
			if cache[m-1][n] == 0 {
				cache[m-1][n] = dp(m-1, n)
			}
			cnt += cache[m-1][n]
		}
		if n > 0 {
			if cache[m][n-1] == 0 {
				cache[m][n-1] = dp(m, n-1)
			}
			cnt += cache[m][n-1]
		}
		return cnt
	}

	return dp(mLen-1, nLen-1)
}
