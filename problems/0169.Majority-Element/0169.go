package p0169

func majorityElement(nums []int) int {
	target, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == target {
			count++
		} else {
			if count-1 == 0 {
				target = nums[i]
			} else {
				count--
			}
		}
	}
	return target
}
