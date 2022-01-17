package p1365

// 暴力解, T: O(n^2), M: O(1)
func smallerNumbersThanCurrent_1(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				result[i]++
			} else if nums[i] < nums[j] {
				result[j]++
			}
		}
	}
	return result
}

// 排序, T: O(n+k), M: O(k), k 為數值範圍大小 0 <= nums[i] <= 100
func smallerNumbersThanCurrent_2(nums []int) []int {
	cnt := make([]int, 101)
	for _, v := range nums {
		cnt[v]++
	}
	for i := 1; i < 101; i++ {
		cnt[i] += cnt[i-1]
	}
	ans := make([]int, len(nums))
	for i, v := range nums {
		if v > 0 {
			ans[i] = cnt[v-1]
		}
	}
	return ans
}
