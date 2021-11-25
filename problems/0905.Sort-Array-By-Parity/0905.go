package p0905

func sortArrayByParity(nums []int) []int {
	for left, right := 0, len(nums)-1; left <= right; {
		if nums[left]%2 == 0 {
			left++
		} else if nums[right]%2 == 1 {
			right--
		} else {
			nums[right], nums[left] = nums[left], nums[right]
		}
	}
	return nums
}
