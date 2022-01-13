package p1010

func numPairsDivisibleBy60(time []int) int {
	m := make(map[int]int, 0)
	pairs := 0
	for _, duration := range time {
		// 條件是 1 <= time[i] <= 500，所以可以用 (600-duration)%60
		// 否則需要改為 (60-duration%60)%60，後面的 %60 是要處理 duration 為 60 的情況
		if v, ok := m[(600-duration)%60]; ok {
			pairs += v
		}

		// 可供配對的數量加 1
		m[duration%60]++
	}
	return pairs
}
