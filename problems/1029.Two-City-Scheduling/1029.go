package p1029

import "sort"

func twoCitySchedCost(costs [][]int) int {
	// 依兩城市相差距離最大的排列，優先分配差異大的
	sort.Slice(costs, func(i, j int) bool {
		return abs(costs[i][0]-costs[i][1]) > abs(costs[j][0]-costs[j][1])
	})
	ans := 0
	aCount := 0
	bCount := 0
	for _, cost := range costs {
		// 選擇 a city 的數量達到一半，直接選擇 b city
		if aCount == len(costs)/2 {
			ans += cost[1]
			bCount++
			// 選擇 b city 的數量達到一半，直接選擇 a city
		} else if bCount == len(costs)/2 {
			ans += cost[0]
			aCount++
			// 選擇 cost 較小的
		} else if cost[0] < cost[1] {
			ans += cost[0]
			aCount++
		} else {
			ans += cost[1]
			bCount++
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
