package p2162

import "math"

// 要特別注意 6039 的情況，必須要把 100:39 轉為 99:99 處理
func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	min := targetSeconds / 60
	sec := targetSeconds % 60

	// 把有可能的情況先存起來，之後再ㄧ併計算 cost
	possible := [][]int{}

	// 第一種情況： 把一分種補到秒數不會超過 99
	if targetSeconds > 60 && sec+60 < 100 {
		m := min - 1
		s := sec + 60
		p := []int{}
		for s > 0 {
			p = append(p, s%10)
			s = s / 10
		}
		for m > 0 {
			p = append(p, m%10)
			m = m / 10
		}
		possible = append(possible, p)
	}

	// 第二種情況： 需要考慮秒數需要補 0 的情況，及分鐘不能超過 99
	if min < 100 {
		p := []int{}
		paddingZeroCount := 0
		if min > 0 {
			if sec == 0 {
				paddingZeroCount = 2
			} else if sec < 10 {
				paddingZeroCount = 1
			}
		}
		for sec > 0 {
			p = append(p, sec%10)
			sec = sec / 10
		}
		for paddingZeroCount > 0 {
			p = append(p, 0)
			paddingZeroCount--
		}
		for min > 0 {
			p = append(p, min%10)
			min = min / 10
		}
		possible = append(possible, p)
	}

	// 計算出兩種情況中 cost 比較小的
	minCost := math.MaxInt32
	for _, pp := range possible {
		now := startAt
		score := 0
		for i := len(pp) - 1; i >= 0; i-- {
			if now == pp[i] {
				score += pushCost
			} else {
				score += pushCost
				score += moveCost
			}
			now = pp[i]
		}
		if score < minCost {
			minCost = score
		}
	}

	return minCost
}
