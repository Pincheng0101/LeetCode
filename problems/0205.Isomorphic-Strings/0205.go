package p0205

func isIsomorphic(s string, t string) bool {
	m := map[byte]byte{}
	set := map[byte]struct{}{}
	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if v, ok := m[a]; ok {
			if b != v {
				return false
			}
		} else {
			if _, ok := set[b]; ok {
				return false
			}
			m[a] = b
			set[b] = struct{}{}
		}
	}
	return true
}
