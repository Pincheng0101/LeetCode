package p2164

// 可以選擇任何一種排序分別對奇數及偶數位置分別進行排序
func sortEvenOdd(nums []int) []int {
	// 偶數部分由小排到大
	bubbleSort(nums, 0, 2, false)
	// 奇數部分由大排到小
	bubbleSort(nums, 1, 2, true)
	return nums
}

func bubbleSort(nums []int, start int, steps int, reverse bool) {
	for i := start; i < len(nums)-steps; i += steps {
		for j := i + steps; j < len(nums); j += steps {
			if reverse {
				if nums[i] < nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
				}
			} else {
				if nums[i] > nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
				}
			}

		}
	}
}
