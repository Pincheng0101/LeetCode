package p0567

// sliding window
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	chars1, chars2 := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		chars1[s1[i]-'a']++
		chars2[s2[i]-'a']++
	}
	for i := len(s1); i < len(s2) && chars1 != chars2; i++ {
		chars2[s2[i-len(s1)]-'a']--
		chars2[s2[i]-'a']++
	}
	return chars1 == chars2
}
