package p0349

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersection(nums2, nums1)
	}
	m := map[int]struct{}{}
	for _, v := range nums1 {
		m[v] = struct{}{}
	}

	intersection := []int{}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			intersection = append(intersection, v)
			delete(m, v)
		}
	}
	return intersection
}
