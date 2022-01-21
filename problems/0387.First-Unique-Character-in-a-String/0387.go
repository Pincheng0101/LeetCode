package p0387

func firstUniqChar(s string) int {
	m := map[byte]int{}
	for _, c := range s {
		m[byte(c)]++
	}
	for i, c := range s {
		if m[byte(c)] == 1 {
			return i
		}
	}
	return -1
}
