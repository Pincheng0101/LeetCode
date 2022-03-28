package p2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := map[int]struct{}{}
	m2 := map[int]struct{}{}
	for _, v := range nums1 {
		m1[v] = struct{}{}
	}
	for _, v := range nums2 {
		m2[v] = struct{}{}
	}
	ans := make([][]int, 2)
	ans[0], ans[1] = make([]int, 0), make([]int, 0)
	for key := range m2 {
		if _, ok := m1[key]; !ok {
			ans[1] = append(ans[1], key)
		}
	}
	for key := range m1 {
		if _, ok := m2[key]; !ok {
			ans[0] = append(ans[0], key)
		}
	}
	return ans
}
