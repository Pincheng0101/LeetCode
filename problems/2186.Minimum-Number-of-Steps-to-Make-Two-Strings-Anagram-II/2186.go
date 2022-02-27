package p2186

func minSteps(s string, t string) int {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[byte(s[i])]++
	}
	matchCount := 0
	for i := 0; i < len(t); i++ {
		if m[byte(t[i])] > 0 {
			m[byte(t[i])]--
			matchCount++
		}
	}
	return len(s) + len(t) - 2*matchCount
}
