package p0997

// 法官會受到 n - 1 人的信任
func findJudge(n int, trust [][]int) int {
	degrees := make([]int, n+1)
	for _, t := range trust {
		degrees[t[0]]--
		degrees[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if degrees[i] == n-1 {
			return i
		}
	}
	return -1
}
