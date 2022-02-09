package p0532

func findPairs(nums []int, k int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	ans := 0
	if k == 0 {
		for _, cnt := range m {
			if cnt > 1 {
				ans++
			}
		}
	} else {
		for key, _ := range m {
			if _, ok := m[key+k]; ok {
				ans++
			}
		}
	}
	return ans
}
