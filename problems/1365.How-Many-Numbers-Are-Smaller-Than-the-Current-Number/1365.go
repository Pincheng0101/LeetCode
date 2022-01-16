package p1365

func smallerNumbersThanCurrent(nums []int) []int {
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
