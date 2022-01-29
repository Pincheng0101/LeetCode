package p0058

func lengthOfLastWord(s string) int {
	i, lastWordLen := len(s)-1, 0
	for i >= 0 && s[i] == ' ' {
		i--
	}

	for ; i >= 0; i-- {
		if s[i] != ' ' {
			lastWordLen++
		} else {
			break
		}
	}

	return lastWordLen
}
