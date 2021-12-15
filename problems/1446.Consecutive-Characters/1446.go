package p1446

func maxPower(s string) int {
	if len(s) < 1 {
		return 0
	}
	max := 1
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		if count > max {
			max = count
		}
	}
	return max
}
