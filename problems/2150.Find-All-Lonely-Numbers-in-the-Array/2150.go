package p2150

import "sort"

// Sort, sort time + O(n) time, O(1) space
func findLonely_1(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}
	sort.Ints(nums)
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			// 檢查右邊的不能和中間值相同, 不能為中間值+1
			if nums[i+1] != nums[i] && nums[i+1] != nums[i]+1 {
				result = append(result, nums[i])
			}
		} else if i == n-1 {
			// 檢查左邊的不能和中間值相同, 不能為中間值-1
			if nums[i-1] != nums[i] && nums[i-1] != nums[i]-1 {
				result = append(result, nums[i])
			}
		} else {
			// 檢查左邊的不能和中間值相同及中間值-1
			// 檢查右邊的不能和中間值相同及中間值+1
			if nums[i-1] != nums[i] && nums[i-1] != nums[i]-1 &&
				nums[i+1] != nums[i] && nums[i+1] != nums[i]+1 {
				result = append(result, nums[i])
			}
		}
	}
	return result
}

// Hashmap, O(2n) time, O(n) space
func findLonely_2(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}

	// val: count
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}

	result := make([]int, 0)
	for _, v := range nums {
		if m[v] == 1 && m[v-1] == 0 && m[v+1] == 0 {
			result = append(result, v)
		}
	}
	return result
}
