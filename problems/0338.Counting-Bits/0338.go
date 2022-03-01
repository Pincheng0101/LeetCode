package p0338

func countBits(n int) []int {
	ans := make([]int, n+1)
	pow := 1
	for i := 1; i <= n; i++ {
		if i == pow*2 {
			pow *= 2
		}
		ans[i] = 1 + ans[i-pow]
	}
	return ans
}
