package p2185

func prefixCount(words []string, pref string) int {
	count := 0
	for _, s := range words {
		if len(s) >= len(pref) {
			if pref == s[:len(pref)] {
				count++
			}
		}
	}
	return count
}
