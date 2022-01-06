package p2114

func mostWordsFound(sentences []string) int {
	maxSpacesCount := 0
	for _, sentence := range sentences {
		count := 0
		for _, c := range sentence {
			if c == ' ' {
				count++
			}
		}
		if count > maxSpacesCount {
			maxSpacesCount = count
		}
	}
	return maxSpacesCount + 1
}
