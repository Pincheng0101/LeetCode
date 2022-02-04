package p0525

// PrefixSum + HashMap
func findMaxLength(nums []int) int {
	m := map[int]int{}
	sum, maxLength := 0, 0
	for i, v := range nums {
		if v == 1 {
			sum++
		} else {
			sum--
		}
		if sum == 0 {
			maxLength = i + 1
		} else if prefixIndex, ok := m[sum]; ok {
			maxLength = max(maxLength, i-prefixIndex)
		} else {
			m[sum] = i
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
