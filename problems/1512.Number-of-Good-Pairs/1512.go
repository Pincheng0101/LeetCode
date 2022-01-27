package p1512

func numIdenticalPairs(nums []int) int {
	count := 0
	m := map[int]int{}
	for _, v := range nums {
		count += m[v]
		m[v]++
	}

	return count
}
