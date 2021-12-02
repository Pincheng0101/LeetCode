package p0026

// 題目有提到不用管後面的值是什麼，只要把不重複的數字移到前面就行
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := 1
	v := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != v {
			nums[n] = nums[i]
			v = nums[i]
			n++
		}
	}
	return n
}
