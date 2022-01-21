package p0344

func reverseString(s []byte) {
	for i, n := 0, len(s); i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}
