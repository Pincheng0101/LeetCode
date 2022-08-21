package p2381

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diff := make([]int, n+1)

	for _, shift := range shifts {
		start, end, d := shift[0], shift[1], shift[2]
		if d == 0 {
			diff[start]--
			diff[end+1]++
		} else {
			diff[start]++
			diff[end+1]--
		}
	}

	acc := 0
	bytes := []byte(s)
	for i := 0; i < n; i++ {
		acc = (acc + diff[i]) % 26
		bytes[i] = 'a' + (bytes[i]-'a'+byte(acc)+26)%26
	}

	return string(bytes)
}
