package p0013

func romanToInt(s string) int {
	ans := 0
	before := '0'
	for _, c := range s {
		switch c {
		case 'I':
			ans += 1
		case 'V':
			ans += 5
			if before == 'I' {
				ans -= 2
			}
		case 'X':
			ans += 10
			if before == 'I' {
				ans -= 2
			}
		case 'L':
			ans += 50
			if before == 'X' {
				ans -= 20
			}
		case 'C':
			ans += 100
			if before == 'X' {
				ans -= 20
			}
		case 'D':
			ans += 500
			if before == 'C' {
				ans -= 200
			}
		case 'M':
			ans += 1000
			if before == 'C' {
				ans -= 200
			}
		}
		before = c
	}
	return ans
}
