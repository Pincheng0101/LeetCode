package p1051

import "sort"

func heightChecker(heights []int) int {
	dups := make([]int, len(heights))
	copy(dups, heights)
	sort.Ints(dups)

	diff := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != dups[i] {
			diff++
		}
	}
	return diff
}
