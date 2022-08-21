package p0848

// difference algorithm
func shiftingLetters(s string, shifts []int) string {
	acc := 0
	bytes := []byte(s)
	for i := len(s) - 1; i >= 0; i-- {
		acc = (acc + shifts[i]) % 26
		bytes[i] = (bytes[i]-'a'+byte(acc))%26 + 'a'
	}

	return string(bytes)
}
