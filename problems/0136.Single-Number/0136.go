package p0136

func singleNumber(nums []int) int {
	var result int = 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}
