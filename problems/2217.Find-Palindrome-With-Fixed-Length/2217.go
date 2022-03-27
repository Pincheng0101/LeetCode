package p2217

// Formula
func kthPalindrome(queries []int, intLength int) []int64 {
	ans := make([]int64, len(queries))
	for i, query := range queries {
		var base int64 = halfBase(intLength)
		if int64(query) > base*9 {
			ans[i] = -1
			continue
		}
		var half int64 = base + int64(query) - 1
		var reverseHalfNum int64

		if isEven(intLength) {
			reverseHalfNum = reverseNum(half)
		} else {
			reverseHalfNum = reverseNum(half / 10)
		}
		for i := 0; i < intLength/2; i++ {
			half *= 10
		}
		ans[i] = half + reverseHalfNum
	}
	return ans
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func halfBase(n int) int64 {
	if n == 1 {
		return 1
	}
	var base int64 = 1
	if float64(n)/2 > float64(n/2) {
		base *= 10
	}
	for i := 0; i < n/2-1; i++ {
		base *= 10
	}
	return base
}

func reverseNum(num int64) int64 {
	var v int64 = 0
	for num > 0 {
		res := num % 10
		num /= 10
		v = v*10 + res
	}
	return v
}

// Time Exceeded
// func kthPalindrome(queries []int, intLength int) []int64 {
// 	ans := make([]int64, 0)
// 	m := map[int]int64{}

// 	var num int64 = 1
// 	for intLength > 1 {
// 		num *= 10
// 		intLength--
// 	}
// 	var max int64 = num * 10

// 	cnt := 0
// 	for i := 0; i < len(queries); i++ {
// 		if m[queries[i]] > 0 {
// 			ans = append(ans, m[queries[i]])
// 			continue
// 		}
// 		flag := true
// 		for num < max {
// 			if isPalindrome(num) {
// 				cnt++
// 				m[cnt] = num
// 				if queries[i] == cnt {
// 					ans = append(ans, num)
// 					num++
// 					flag = false
// 					break
// 				}
// 			}
// 			num++
// 		}
// 		if flag && num >= max {
// 			ans = append(ans, -1)
// 		}

// 	}

// 	return ans
// }

// func isPalindrome(x int64) bool {
// 	if x < 0 {
// 		return false
// 	}
// 	if x < 10 {
// 		return true
// 	}
// 	var xReversed int64 = reverse(x)

// 	return xReversed == x
// }

// func reverse(x int64) int64 {
// 	sign := "positive"
// 	if x >= 0 {
// 		sign = "positive"
// 	} else {
// 		sign = "negative"
// 	}

// 	x = int64(math.Abs(float64(x)))

// 	var reversedDigit int64

// 	for x > 0 {
// 		var lastDigit int64 = x % 10
// 		reversedDigit = reversedDigit*10 + lastDigit

// 		x = x / 10
// 	}

// 	if sign == "negative" {
// 		reversedDigit = reversedDigit * -1
// 	}

// 	return reversedDigit
// }
