package p2180

func countEven(num int) int {
	v := 0
	ans := num / 2
	n := num
	for n != 0 {
		v += n % 10
		n /= 10
	}

	if v%2 == 1 && num%2 != 1 {
		ans -= 1
	}

	return ans
}
