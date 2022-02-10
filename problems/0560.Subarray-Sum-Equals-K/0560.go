package p0560

// Without Space
func subarraySum_1(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

// Hashmap
func subarraySum_2(nums []int, k int) int {
	// TODO: implement
	return 0
}
