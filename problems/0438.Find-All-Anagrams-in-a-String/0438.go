package p0438

// Sliding Window
func findAnagrams(s string, p string) []int {
	chars, flag := [26]int{}, [26]bool{}
	for i := 0; i < len(p); i++ {
		flag[val(p[i])] = true
		chars[val(p[i])]++
	}
	cnt, left, ans := 0, 0, []int{}
	for right := 0; right < len(s); {
		if flag[val(s[right])] && chars[val(s[right])] > 0 {
			if cnt == 0 {
				left = right
			}
			chars[val(s[right])]--
			cnt++
			if cnt == len(p) && right-left+1 == len(p) {
				ans = append(ans, left)
				chars[val(s[left])]++
				left++
				cnt--
			}
			right++
		} else if flag[val(s[left])] {
			chars[val(s[left])]++
			left++
			cnt--
		} else {
			right++
		}
	}
	return ans
}

func val(b byte) int {
	return int(b - byte('a'))
}
