package p1089

// 1: 先複製一份 arr 出來再去修改原先的 arr
func duplicateZeros_1(arr []int) {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	for arrIndex, arrCopyIndex := 0, 0; arrIndex < len(arr); {
		if arrCopy[arrCopyIndex] == 0 {
			arr[arrIndex] = 0
			arrIndex++
			if arrIndex < len(arr) {
				arr[arrIndex] = 0
			}
		} else {
			arr[arrIndex] = arrCopy[arrCopyIndex]
		}
		arrIndex++
		arrCopyIndex++
	}
}
