package p0028

import "strings"

// Runtime: 4 ms, Memory Usage: 2.5 MB
func strStr_1(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// Runtime: 0 ms, Memory Usage: 2.5 MB
func strStr_2(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
