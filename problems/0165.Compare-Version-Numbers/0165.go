package p0165

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	n1, n2 := len(version1), len(version2)
	for i < n1 || j < n2 {
		v1, v2 := 0, 0
		for i < n1 && version1[i] != '.' {
			v1 = v1*10 + int(version1[i]-'0')
			i++
		}
		for j < n2 && version2[j] != '.' {
			v2 = v2*10 + int(version2[j]-'0')
			j++
		}
		if v1 > v2 {
			return 1
		}
		if v2 > v1 {
			return -1
		}
		i, j = i+1, j+1
	}
	return 0
}
