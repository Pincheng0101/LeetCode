package p0007

func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if int(int32(result)) != result {
		return 0
	}

	return result
}
