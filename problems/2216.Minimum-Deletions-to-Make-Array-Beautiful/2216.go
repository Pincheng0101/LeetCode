package p2216

func minDeletion(nums []int) int {
	reserve := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			reserve += 2
			i++
		}
	}
	return len(nums) - reserve
}
