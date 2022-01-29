package p0055

func canJump(nums []int) bool {
	size := len(nums)
	farest := 0
	for i := 0; i < size; i++ {
		if i+nums[i] > farest {
			farest = i + nums[i]
		}
		if farest >= size-1 {
			return true
		}
		if nums[i] == 0 && farest == i {
			return false
		}
	}
	return false
}
