package p0283

// 先把非 0 的值複製到前面，再把後面的部分改成 0
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
	}

	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}
