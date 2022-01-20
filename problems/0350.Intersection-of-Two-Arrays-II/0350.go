package p0350

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, v := range nums1 {
		m[v]++
	}

	intersect := []int{}
	for _, v := range nums2 {
		if m[v] > 0 {
			intersect = append(intersect, v)
			m[v]--
		}
	}
	return intersect
}
