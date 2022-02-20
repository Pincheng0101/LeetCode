package p1288

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	ans := 1
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] > right {
			right = intervals[i][1]
			ans++
		}
	}
	return ans
}
