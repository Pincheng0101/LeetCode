package p0740

func deleteAndEarn(nums []int) int {
	points := map[int]int{}
	maxNumber := 0
	for _, v := range nums {
		points[v] += v
		maxNumber = max(maxNumber, v)
	}

	maxPoints := make([]int, maxNumber+1)
	maxPoints[1] = points[1]

	for i := 2; i < len(maxPoints); i++ {
		maxPoints[i] = max(maxPoints[i-1], maxPoints[i-2]+points[i])
	}
	return maxPoints[maxNumber]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
