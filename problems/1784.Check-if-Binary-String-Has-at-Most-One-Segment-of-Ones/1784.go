package p1784

import "strings"

func checkOnesSegment(s string) bool {
	return strings.Index(s, "01") == -1
}
