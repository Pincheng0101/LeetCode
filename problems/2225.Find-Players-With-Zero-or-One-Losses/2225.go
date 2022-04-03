package p2225

import "sort"

func findWinners(matches [][]int) [][]int {
	winner := map[int]bool{}
	loserCnt := map[int]int{}
	for _, match := range matches {
		if _, ok := winner[match[0]]; !ok {
			winner[match[0]] = true
		}
		winner[match[1]] = false
		loserCnt[match[1]]++
	}

	ans := make([][]int, 2)
	ans[0] = make([]int, 0)
	ans[1] = make([]int, 0)
	for key, v := range winner {
		if v == true {
			ans[0] = append(ans[0], key)
		}
	}
	for key, v := range loserCnt {
		if v == 1 {
			ans[1] = append(ans[1], key)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	return ans
}
