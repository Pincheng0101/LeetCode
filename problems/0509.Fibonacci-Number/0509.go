package p0509

func fib_1(n int) int {
	if n <= 1 {
		return n
	}
	return fib_1(n-1) + fib_1(n-2)
}

func fib_2(n int) int {
	cache := map[int]int{}
	var fibWichCache func(int) int
	fibWichCache = func(nn int) int {
		if nn <= 1 {
			return nn
		}
		sum := 0
		if cache[nn-1] > 0 {
			sum += cache[nn-1]
		} else {
			v := fibWichCache(nn - 1)
			cache[nn-1] = v
			sum += v
		}
		if cache[nn-2] > 0 {
			sum += cache[nn-2]
		} else {
			v := fibWichCache(nn - 2)
			cache[nn-2] = v
			sum += v
		}
		return sum
	}
	return fibWichCache(n)
}
