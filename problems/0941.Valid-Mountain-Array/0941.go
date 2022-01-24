package p0941

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	n := len(arr)
	i := 0
	// walk up
	for i < n-1 && arr[i] < arr[i+1] {
		i++
	}

	if i == 0 || i == n-1 {
		return false
	}

	// walk down
	for i < n-1 && arr[i] > arr[i+1] {
		i++
	}

	return i == n-1
}
