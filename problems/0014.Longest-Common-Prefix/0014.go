package p0014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	commonPrefix := ""
out:
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 || strs[j][i] != c {
				break out
			}
		}
		commonPrefix += string(c)
	}
	return commonPrefix
}
