package p2220

import "math/bits"

func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}
