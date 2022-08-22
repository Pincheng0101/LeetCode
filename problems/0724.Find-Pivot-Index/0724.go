package p0724

func pivotIndex(nums []int) int {
	leftSum, rightSum := 0, sum(nums)
	for i, num := range nums {
		rightSum -= num
		if leftSum == rightSum {
			return i
		}
		leftSum += num
	}
	return -1
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
