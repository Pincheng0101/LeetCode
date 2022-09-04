package p2400

func numberOfWays(startPos int, endPos int, k int) int {
	type key struct {
		distance int
		k        int
	}

	cache := map[key]int{}

	var dfs func(int, int) int
	dfs = func(distance int, k int) int {
		if distance == k {
			return 1
		}

		if distance > k {
			return 0
		}

		if v, ok := cache[key{distance: distance, k: k}]; ok {
			return v
		}

		direction := -1
		if distance == 0 {
			direction = 1
		}
		left := dfs(distance+1, k-1)
		right := dfs(distance+direction, k-1)
		sum := (left + right) % 1000000007

		cache[key{distance: distance, k: k}] = sum

		return sum
	}

	return dfs(abs(startPos-endPos), k)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
