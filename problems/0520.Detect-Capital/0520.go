package p0520

import "unicode"

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	if unicode.IsLower(rune(word[0])) && unicode.IsUpper(rune(word[1])) {
		return false
	}

	for i := 1; i < len(word)-1; i++ {
		if unicode.IsLower(rune(word[i])) != unicode.IsLower(rune(word[i+1])) {
			return false
		}
	}

	return true
}
