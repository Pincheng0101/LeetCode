package p1337

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	rowCount := make([][]int, len(mat))
	for i, m := range mat {
		cnt := 0
		for _, v := range m {
			cnt += v
		}
		rowCount[i] = []int{cnt, i}
	}
	sort.Slice(rowCount, func(i, j int) bool {
		if rowCount[i][0] == rowCount[j][0] {
			return rowCount[i][1] < rowCount[j][1]
		}
		return rowCount[i][0] < rowCount[j][0]
	})
	ans := make([]int, k)
	i := 0
	for k > 0 {
		ans[i] = rowCount[i][1]
		i++
		k--
	}
	return ans
}
