package p0488

import (
	"sort"
)

// 1. 建一個新的 Slice，排序後把遺漏的放進去
// Time Complexity: O(n), Space Complexity: O(n)
func findDisappearedNumbers_1(nums []int) []int {
	sort.Ints(nums)
	result := make([]int, 0)
	v := 1
	for i := 0; i < len(nums); {
		if nums[i] == v {
			i++
			v++
		} else if nums[i] < v {
			i++
		} else {
			result = append(result, v)
			v++
		}
	}
	for i := v; i <= len(nums); i++ {
		result = append(result, i)
	}
	return result
}

// 數字範圍最大值和 nums 的長度一樣，使用 nums 充當 hash table
// 把存在的數值資訊放到對應 key 的位置中
// Time Complexity: O(n), Space Complexity: O(1)
func findDisappearedNumbers_2(nums []int) []int {
	result := make([]int, 0)
	for _, v := range nums {
		// 將存在的數值對應的 key 裡面的數值改成負的
		v = abs(v) - 1
		if nums[v] > 0 {
			nums[v] *= -1
		}
		// 也可以用 % n 的方式（因為最大值等於長度）
		// n = len(nums)
		// v = (v - 1) % n
		// nums[v] += n
	}
	for i, v := range nums {
		// 正的代表不存在
		if v > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
