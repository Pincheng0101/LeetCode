package p0413

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sum := 0
	dp := 0
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i-2] == nums[i]-nums[i-1] {
			dp++
			sum += dp
		} else {
			dp = 0
		}
	}
	return sum
}
