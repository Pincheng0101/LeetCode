package p0014

// 兩兩比較
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs) && prefix != ""; i++ {
		prefix = lcp(prefix, strs[i])
	}
	return prefix
}

func lcp(s1, s2 string) string {
	n := min(len(s1), len(s2))
	i := 0
	for i < n && s1[i] == s2[i] {
		i++
	}
	return s1[:i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
