package p0035

// O(n)
func searchInsert_1(nums []int, target int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
		if target < nums[i] {
			break
		}
		index = i + 1
	}
	return index
}

// O(log n)
func searchInsert_2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left
}
