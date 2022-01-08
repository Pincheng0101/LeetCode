package p0003

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int, 0)
	k, max := -1, 0
	for i := 0; i < len(s); i++ {
		// 檢查 map 裡面是否有位置大於 k 的相同字元
		if cIndex, ok := m[s[i]]; ok && cIndex > k {
			// 將 k 移到前一個相同字元的位置
			k = cIndex
			// 原先字元的位置用新的位置取代
			m[s[i]] = i
		} else {
			// 加入新的字元及位置
			m[s[i]] = i
			if i-k > max {
				max = i - k
			}
		}
	}
	return max
}
