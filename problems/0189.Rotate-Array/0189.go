package p0189

// T: O(n), M: O(n)
func rotate_1(nums []int, k int) {
	n := len(nums)
	k = k % n
	copy(nums, append(nums[n-k:], nums[:n-k]...))
}

// 反轉三次, T: O(2n), M: O(1)
func rotate_2(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

// TODO: 環狀替換, T: O(n), M: O(1)
// func rotate_3(nums []int, k int)
