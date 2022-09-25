package p2419

func longestSubarray(nums []int) int {
	maxLength, maxValue := 1, 0
	length, value := 1, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= value && value&nums[i] >= nums[i] {
			length++
			value &= nums[i]
		} else {
			length = 1
			value = nums[i]
		}

		if value > maxValue || (value >= maxValue && length > maxLength) {
			maxValue = value
			maxLength = length
		}
	}

	return maxLength
}
