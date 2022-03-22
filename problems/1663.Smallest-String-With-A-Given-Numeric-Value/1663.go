package p1663

func getSmallestString(n int, k int) string {
	ans := make([]byte, n)
	k -= n
	for i := n - 1; i >= 0; i-- {
		addValue := min(25, k)
		ans[i] = 'a' + byte(addValue)
		k -= addValue
	}
	return string(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
