package p1137

// recursive with cache
func tribonacci_1(n int) int {
	cache := map[int]int{}

	var tribonacciWithCache func(int) int
	tribonacciWithCache = func(nn int) int {
		if nn <= 1 {
			return nn
		}
		if nn == 2 {
			return 1
		}
		sum := 0
		if cache[nn-3] > 0 {
			sum += cache[nn-3]
		} else {
			v := tribonacciWithCache(nn - 3)
			cache[nn-3] = v
			sum += v
		}
		if cache[nn-2] > 0 {
			sum += cache[nn-2]
		} else {
			v := tribonacciWithCache(nn - 2)
			cache[nn-2] = v
			sum += v
		}
		if cache[nn-1] > 0 {
			sum += cache[nn-1]
		} else {
			v := tribonacciWithCache(nn - 1)
			cache[nn-1] = v
			sum += v
		}
		return sum
	}
	return tribonacciWithCache(n)
}

// from bottom to top
func tribonacci_2(n int) int {
	block := []int{0, 1, 1}
	if n <= 2 {
		return block[n]
	}
	for n > 2 {
		block[0], block[1], block[2] =
			block[1], block[2], block[0]+block[1]+block[2]
		n--
	}
	return block[2]
}
