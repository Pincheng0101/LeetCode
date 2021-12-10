package p0001

func twoSum(nums []int, target int) []int {
	// value: index
	m := make(map[int]int)
	m[target-nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return []int{m[nums[i]], i}
		}
		m[target-nums[i]] = i
	}
	return nil
}
