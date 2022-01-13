package p0452

import "sort"

// 將 points 依 xEnd 由小排到大，如果下一個點的 xStart 大於前一個 xEnd，射擊數量就加 1
func findMinArrowShots_1(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	shoots := 1
	prevEnd := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > prevEnd {
			prevEnd = points[i][1]
			shoots++
		}
	}
	return shoots
}

// 將 points 依 xStart 由大排到小，如果下一個點的 xEnd 小於前一個 xStart，射擊數量就加 1
func findMinArrowShots_2(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] > points[j][0]
	})

	shoots := 1
	prevStart := points[0][0]
	for i := 1; i < len(points); i++ {
		if points[i][1] < prevStart {
			prevStart = points[i][0]
			shoots++
		}
	}
	return shoots
}
