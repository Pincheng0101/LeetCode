package p2156

// Reverse Rolling Hash (Rabin-Karp algorithm)
func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	pp, hash, res := 1, 0, 0
	size := len(s)
	for i := size - 1; i >= 0; i-- {
		hash = (hash*power + val(s[i])) % modulo
		if i+k >= size {
			pp = pp * power % modulo
		} else {
			// hash - pp*val(s[i+k] 有可能是負的，所以加一個 modulo 讓它不會是負的
			hash = (modulo + hash - pp*val(s[i+k])%modulo) % modulo
		}
		if hash == hashValue {
			res = i
		}
	}
	return s[res : res+k]
}

func val(b byte) int {
	return int(b - byte('`'))
}
