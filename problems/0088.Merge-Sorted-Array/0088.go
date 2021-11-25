package p0088

import (
	"sort"
)

// 1. 先把 nums2 的值和 nums1 合併，再排序
// Runtime: 0 ms, Memory Usage: 2.4 MB
func merge1(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

// 2. 由於 nums1 的空間為 m + n，所以可以從後面往前比大小補齊
// Runtime: 0 ms, Memory Usage: 2.3 MB
func merge2(nums1 []int, m int, nums2 []int, n int) {
	for last := len(nums1) - 1; last >= 0; last-- {
		if m > 0 && n > 0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[last] = nums1[m-1]
				m--
			} else {
				nums1[last] = nums2[n-1]
				n--
			}
		} else if n > 0 {
			nums1[last] = nums2[n-1]
			n--
		} else {
			break
		}
	}
}
