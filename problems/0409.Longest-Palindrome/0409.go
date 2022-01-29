package p0409

func longestPalindrome(s string) int {
	cnt := map[byte]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}
	ans := 0
	for _, v := range cnt {
		if ans%2 == 1 {
			ans += v / 2 * 2
		} else {
			ans += v
		}

	}
	return ans
}
