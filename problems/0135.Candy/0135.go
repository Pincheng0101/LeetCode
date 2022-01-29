package p0135

func candy(ratings []int) int {
	size := len(ratings)
	if size < 2 {
		return size
	}

	candies := make([]int, size)
	for i := 0; i < size; i++ {
		candies[i] = 1
	}
	for i := 1; i < size; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	for i := size - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}
	count := 0
	for _, v := range candies {
		count += v
	}
	return count
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
