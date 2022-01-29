package p0665

func checkPossibility(nums []int) bool {
	canModified := true
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if !canModified {
				return false
			}
			if i > 1 && nums[i] < nums[i-2] {
				nums[i] = nums[i-1]
			}
			canModified = false
		}
	}
	return true
}
