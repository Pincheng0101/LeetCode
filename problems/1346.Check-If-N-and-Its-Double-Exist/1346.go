package p1346

func checkIfExist(arr []int) bool {
	vMap := make(map[int]struct{})
	for _, v := range arr {
		if _, ok := vMap[v*2]; ok {
			return true
		}
		if _, ok := vMap[v/2]; ok && v%2 == 0 {
			return true
		}
		vMap[v] = struct{}{}
	}
	return false
}
