package p0881

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	light, heavy := 0, len(people)-1
	ans := 0
	for light <= heavy {
		if people[heavy]+people[light] <= limit {
			light++
		}
		heavy--
		ans++
	}
	return ans
}
