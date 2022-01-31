package p2154

import "sort"

func findFinalValue_1(nums []int, original int) int {
	sort.Ints(nums)
	ans := original
	for i := 0; i < len(nums); i++ {
		if nums[i] == ans {
			ans = nums[i] * 2
		}
	}

	return ans
}

func findFinalValue_2(nums []int, original int) int {
	m := map[int]struct{}{}
	for _, v := range nums {
		m[v] = struct{}{}
	}

	ans := original
	for {
		if _, ok := m[ans]; ok {
			ans *= 2
		} else {
			break
		}
	}
	return ans
}
