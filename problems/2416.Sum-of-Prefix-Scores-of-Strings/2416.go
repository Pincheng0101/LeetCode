package p2416

import "strings"

func sumPrefixScores(words []string) []int {
	m := map[string]int{}
	for _, word := range words {
		var s strings.Builder
		s.Grow(len(word))
		for i := 0; i < len(word); i++ {
			s.WriteString(string(word[i]))
			m[s.String()]++
		}
	}

	ans := []int{}
	for _, word := range words {
		sum := 0
		var s strings.Builder
		s.Grow(len(word))
		for i := 0; i < len(word); i++ {
			s.WriteString(string(word[i]))
			sum += m[s.String()]
		}
		ans = append(ans, sum)
	}
	return ans
}

// TODO: Trie Tree Solution
