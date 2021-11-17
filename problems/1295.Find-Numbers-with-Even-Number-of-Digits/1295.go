package leetcode

// 1: 每次除 10 算次數，到 0 停止
func findNumbers_1(nums []int) int {
	count := 0
	for _, num := range nums {
		digits := 0
		for num > 0 {
			num /= 10
			digits += 1
		}
		if digits%2 == 0 {
			count += 1
		}
	}
	return count
}

// 2: 轉成字串後用字串長度判斷
func findNumbers_2(nums []int) int {
	count := 0
	for _, num := range nums {
		s := strconv.Itoa(num)
		if len(s)%2 == 0 {
			count += 1
		}
	}
	return count
}
