package p2220

func minBitFlips(start int, goal int) int {
	ans := 0
	xorNumber := start ^ goal
	for xorNumber > 0 {
		ans += xorNumber & 1
		xorNumber >>= 1
	}
	return ans
}
