package p0125

func isPalindrome(s string) bool {
	for left, right := 0, len(s)-1; left < right; {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	if (c >= byte('a') && c <= byte('z')) ||
		(c >= byte('A') && c <= byte('Z')) ||
		(c >= byte('0') && c <= byte('9')) {
		return true
	}
	return false
}

func toLower(c byte) byte {
	if c >= byte('A') && c <= byte('Z') {
		return c - byte('A') + byte('a')
	}
	return c
}
