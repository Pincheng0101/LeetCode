package p0438

// Sliding Window
func findAnagrams(s string, p string) []int {
	flag, chars := [26]bool{}, [26]int{}
	for i := 0; i < len(p); i++ {
		flag[val(p[i])] = true
		chars[val(p[i])]++
	}
	left, right, cnt, ans := 0, 0, 0, []int{}
	for right < len(s) {
		// check right and resources are available
		if flag[val(s[right])] && chars[val(s[right])] > 0 {
			if cnt == 0 {
				left = right
			}
			chars[val(s[right])]--
			cnt++
			if cnt == len(p) && right-left+1 == len(p) {
				ans = append(ans, left)
				// left recycle
				chars[val(s[left])]++
				left++
				cnt--
			}
			right++
			// check left and recycle, then let right check again (do not right++)
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
