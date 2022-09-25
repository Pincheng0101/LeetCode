package p2414

func longestContinuousSubstring(s string) int {
	max := 0
	slow, fast := 0, 0
	for fast+1 < len(s) {
		if byte(s[fast])+1 == byte(s[fast+1]) {
			fast++
		} else {
			size := fast - slow
			if size > max {
				max = size
			}
			slow = fast + 1
			fast = fast + 1
		}
	}
	if max == 0 {
		return fast - slow + 1
	}
	return max + 1
}
