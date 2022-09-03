package p2396

// Proper Approach
func isStrictlyPalindromic_1(n int) bool {
	for i := n - 2; i >= 2; i-- {
		tmp := n
		nums := []int{}
		for tmp != 0 {
			nums = append(nums, tmp%i)
			tmp /= i
		}
		for j := 0; j < len(nums)/2; j++ {
			if nums[j] != nums[len(nums)-j-1] {
				return false
			}
		}
	}
	return true
}

// 從 n = 4 和 5 都不是嚴格的回文，更多就更不可能符合
func isStrictlyPalindromic(n int) bool {
	return false
}
