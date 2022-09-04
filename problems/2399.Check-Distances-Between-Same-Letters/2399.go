package p2399

func checkDistances(s string, distance []int) bool {
	m := map[byte]int{}
	for nowIndex, char := range s {
		if beforeIndex, ok := m[byte(char)]; ok {
			if nowIndex-beforeIndex-1 != distance[char-'a'] {
				return false
			}
		}
		m[byte(char)] = nowIndex
	}
	return true
}
