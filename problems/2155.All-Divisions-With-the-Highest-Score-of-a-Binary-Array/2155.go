package p1001

func maxScoreIndices(nums []int) []int {
	m := map[int]int{}
	numsLeftScore := 0
	numsRightScore := 0
	for _, v := range nums {
		if v == 1 {
			numsRightScore++
		}
	}
	m[0] = numsRightScore
	maxValue := m[0]
	for i := 1; i < len(nums)+1; i++ {
		if nums[i-1] == 1 {
			numsRightScore--
		} else {
			numsLeftScore++
		}
		m[i] = numsLeftScore + numsRightScore
		if m[i] > maxValue {
			maxValue = m[i]
		}
	}
	ans := []int{}
	for i, v := range m {
		if v == maxValue {
			ans = append(ans, i)
		}
	}
	return ans
}
