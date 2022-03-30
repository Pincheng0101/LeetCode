package p0074

func searchMatrix(matrix [][]int, target int) bool {
	top, bottom := 0, len(matrix)-1
	for top <= bottom {
		mid := top + (bottom-top)/2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] < target {
			top = mid + 1
		} else {
			bottom = mid - 1
		}
	}
	if top > 0 {
		top--
	}
	left, right := 0, len(matrix[top])-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[top][mid] == target {
			return true
		}
		if matrix[top][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
