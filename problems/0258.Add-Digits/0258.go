package p0258

// 0 0
// 1 1 10 1 91 1
// 2 2 11 2 92 2
// 3 3 12 3 93 3
// 4 4 13 4 94 4
// 5 5 14 5 95 5
// 6 6 15 6 96 6
// 7 7 16 7 97 7
// 8 8 17 8 98 8
// 9 9 18 9 99 9

// 可以看到數字從 1 開始會以 9 為一個循環，可以直接用 modulo 9 計算結果
func addDigits(num int) int {
	return 1 + (num-1)%9
}
