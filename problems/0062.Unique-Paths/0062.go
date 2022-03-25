package p0062

// 由上而下
func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}

	var dfs func(int, int) int
	dfs = func(m, n int) int {
		ans := 0
		if m == 0 && n == 0 {
			return 1
		}
		if m > 0 {
			if cache[m-1][n] == 0 {
				cache[m-1][n] = dfs(m-1, n)
			}
			ans += cache[m-1][n]
		}
		if n > 0 {
			if cache[m][n-1] == 0 {
				cache[m][n-1] = dfs(m, n-1)
			}
			ans += cache[m][n-1]
		}
		return ans
	}

	return dfs(m-1, n-1)
}
