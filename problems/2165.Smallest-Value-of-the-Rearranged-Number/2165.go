package p2165

import "sort"

func smallestNumber(num int64) int64 {
	// 需要分別處理正數和負數的部分
	isPositive := true
	if num < 0 {
		isPositive = false
	}

	nums := []int64{}
	for num != 0 {
		nums = append(nums, num%10)
		num = num / 10
	}

	// 由大排到小
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	var ans int64 = 0
	if isPositive {
		findFirstNonZero := true
		for i := 0; i < len(nums); i++ {
			// 找到第一個不是 0 的值
			if nums[i] != 0 && findFirstNonZero {
				ans = nums[i]
				for zeroCount := i; zeroCount > 0; zeroCount-- {
					ans *= 10
				}
				findFirstNonZero = false
			} else {
				ans = ans*10 + nums[i]
			}
		}
	} else {
		for _, v := range nums {
			ans = ans*10 + abs(v)
		}
		ans = -ans
	}

	return ans
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
