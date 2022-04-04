package p0070

func climbStairs(n int) int {
	block := []int{1, 1}
	if n <= 1 {
		return block[n]
	}
	for i := 0; i < n-1; i++ {
		block[0], block[1] = block[1], block[0]+block[1]
	}
	return block[1]
}
