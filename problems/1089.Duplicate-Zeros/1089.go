package p1089

// 1. 先複製一份 arr 出來再去修改原先的 arr
// Runtime: 4ms, Memory Usage: 3.5MB
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

// 2. 先計算 0 的數量，再從後面往前補
// Runtime: 4ms, Memory Usage: 3.3MB
func duplicateZeros_2(arr []int) {
	zeroCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeroCount++
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if i+zeroCount < len(arr) {
				arr[i+zeroCount] = 0
			}
			if i+zeroCount-1 < len(arr) {
				arr[i+zeroCount-1] = 0
			}
			zeroCount--
		} else {
			if i+zeroCount < len(arr) {
				arr[i+zeroCount] = arr[i]
			}
		}
	}
}

// 3. 同上，不過只計算複製 0 的數量，要額外判斷最後一個是 0 的情況，有可能不需要複製
// Runtime: 4ms, Memory Usage: 3.3MB
func duplicateZeros_3(arr []int) {
	zeroDupCount := 0
	arrLen := len(arr)
	for i := 0; i < arrLen-zeroDupCount; i++ {
		if arr[i] == 0 {
			// 事先處理好最後一個是 0 的情況
			// 後面只依靠 zeroDupCount 無法判斷最後一個是否需要複製 0
			if i == arrLen-zeroDupCount-1 {
				arrLen--
				arr[arrLen] = 0
				break
			}
			zeroDupCount++
		}
	}

	for i := arrLen - zeroDupCount - 1; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+zeroDupCount] = 0
			zeroDupCount--
			arr[i+zeroDupCount] = 0
		} else {
			arr[i+zeroDupCount] = arr[i]
		}
	}
}
