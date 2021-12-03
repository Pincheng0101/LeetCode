package p0152

// 概念和 0053 一樣，只是要注意正負號
func maxProduct(nums []int) int {
	minSub, maxSub, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		minSub, maxSub =
			min(nums[i], minSub*nums[i], maxSub*nums[i]),
			max(nums[i], minSub*nums[i], maxSub*nums[i])
		ans = max(ans, maxSub)
	}
	return ans
}

func min(v ...int) int {
	min := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < min {
			min = v[i]
		}
	}
	return min
}

func max(v ...int) int {
	max := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > max {
			max = v[i]
		}
	}
	return max
}
