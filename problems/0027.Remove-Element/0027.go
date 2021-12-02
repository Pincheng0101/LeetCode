package p0027

// 題目有提到不用管後面的值是什麼，只要把和 val 不一樣的數字移到前面就行
func removeElement(nums []int, val int) int {
	k := 0
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}
	return k
}
