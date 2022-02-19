package p2178

func maximumEvenSplit(finalSum int64) []int64 {
	ans := []int64{}
	if finalSum%2 != 0 {
		return ans
	}
	var v int64 = 2
	for finalSum-v > v {
		ans = append(ans, v)
		finalSum -= v
		v += 2
	}
	if finalSum != 0 {
		ans = append(ans, finalSum)
	}
	return ans
}
