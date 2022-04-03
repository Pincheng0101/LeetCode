package p2221

func triangularSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	newNums := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		newNums[i] = (nums[i] + nums[i+1]) % 10
	}
	return triangularSum(newNums)
}
