package p0009

import "strconv"

func isPalindrome_1(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome_2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	v := 0
	for x > v {
		v = v*10 + x%10
		x = x / 10
	}
	return v == x || v/10 == x
}
