package p0217

func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = struct{}{}
		}
	}
	return false
}
