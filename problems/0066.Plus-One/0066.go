package p0066

func plusOne(digits []int) []int {
	carry := 0
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i] + carry
		carry = v / 10
		digits[i] = v % 10
	}
	if carry > 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}
