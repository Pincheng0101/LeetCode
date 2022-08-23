package p0066

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			break
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
