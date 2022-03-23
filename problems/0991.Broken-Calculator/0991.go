package p0991

func brokenCalc(startValue int, target int) int {
	ans := 0
	for target > startValue {
		ans += 1
		if target%2 != 0 {
			target += 1
		} else {
			target /= 2
		}
	}
	return ans + startValue - target
}
