package p0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var table [26]int
	for _, ch := range s {
		table[ch-'a']++
	}
	for _, ch := range t {
		table[ch-'a']--
		if table[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
