package p0485

func findMaxConsecutiveOnes(nums []int) int {
	counter := 0
	max_count := 0
	for _, num := range nums {
		if num == 1 {
			counter += 1
			if counter > max_count {
				max_count = counter
			}
		} else {
			counter = 0
		}
	}
	return max_count
}
