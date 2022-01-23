package p2148

func countElements(nums []int) int {
	count := 0
	min, max := findMinAndMax(nums)
	for _, v := range nums {
		if v > min && v < max {
			count++
		}
	}
	return count
}

func findMinAndMax(nums []int) (min int, max int) {
	min, max = nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}
