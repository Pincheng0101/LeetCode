package p0074

func searchMatrix(matrix [][]int, target int) bool {
    row, col := len(matrix), len(matrix[0])
    left, right := 0, row*col-1
    for left <= right {
        mid := left + (right - left)/2
        v := matrix[mid/col][mid%col]
        if v == target {
            return true
        }
        if v < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}