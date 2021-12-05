package p1217

// 只要位置不要改變奇偶就不會有消耗，所以只要先把奇偶位置的數量計算出來，再把數量少的移過去就行
func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, v := range position {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd < even {
		return odd
	}
	return even
}
