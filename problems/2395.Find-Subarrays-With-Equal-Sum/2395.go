package p2395

func findSubarrays(nums []int) bool {
	m := map[int]struct{}{}
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i] + nums[i+1]
		if _, ok := m[sum]; ok {
			return true
		}
		m[sum] = struct{}{}
	}
	return false
}
