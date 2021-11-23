package p0941

// 先遞增，後遞減
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	up := arr[1] > arr[0]
	if !up {
		return false
	}
	for i := 2; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
		if up {
			if arr[i] < arr[i-1] {
				up = false
			}
		} else {
			if arr[i] > arr[i-1] {
				return false
			}
		}
	}
	return !up
}
