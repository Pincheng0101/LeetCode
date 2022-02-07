package p0389

func findTheDifference(s string, t string) byte {
	var i int
	var ans byte
	for i < len(s) {
		ans -= s[i]
		ans += t[i]
		i++
	}
	ans += t[i]
	return ans
}
