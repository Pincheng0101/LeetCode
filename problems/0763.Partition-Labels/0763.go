package p0763

func partitionLabels(s string) []int {
	last := map[byte]int{}
	for i := 0; i < len(s); i++ {
		last[s[i]] = i
	}
	ans := make([]int, 0)
	anchor := 0
	for i, intervalEnd := 0, 0; i < len(s); i++ {
		intervalEnd = max(intervalEnd, last[s[i]])
		// 目前位置剛好等於這個位置上的值的最後位置，代表後面不會有一樣的值，可劃分一部分出來
		if i == intervalEnd {
			ans = append(ans, i-anchor+1)
			anchor = i + 1
		}
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
