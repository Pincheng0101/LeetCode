package p2226

import "sort"

func maximumCandies(candies []int, k int64) int {
	sort.Ints(candies)
	ans := 0
	left, right := 0, candies[len(candies)-1]
	for left <= right {
		mid := left + (right-left)/2
		if check(candies, mid+1, k) {
			ans = max(ans, mid+1)
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func check(candies []int, target int, k int64) bool {
	sum := 0
	for _, candyCnt := range candies {
		sum += candyCnt / target
	}
	return int64(sum) >= k
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
