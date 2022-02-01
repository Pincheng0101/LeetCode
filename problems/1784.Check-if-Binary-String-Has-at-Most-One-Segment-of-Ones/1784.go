package p1784

import "strings"

func checkOnesSegment_1(s string) bool {
	return strings.Index(s, "01") == -1
}

func checkOnesSegment_2(s string) bool {
	i := 1
	for ; i < len(s) && s[i] == '1'; i++ {
	}
	for ; i < len(s) && s[i] == '0'; i++ {
	}
	return i == len(s)
}
