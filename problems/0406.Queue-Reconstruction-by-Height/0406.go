package p0406

import (
	"container/list"
	"sort"
)

// 使用 linked list 的方式 insert
func reconstructQueue_1(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		// 同樣身高 k 值小的在前面
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		// 身高高的在前面
		return people[i][0] > people[j][0]
	})

	// 接下來只要按照 k 值將它插入到對應的位置就行了
	l := list.New()
	head := l.PushBack(0)
	cur := head
	curIndex := 0
	for i, p := range people {
		for p[1] != curIndex {
			cur = cur.Next()
			curIndex++
		}
		l.InsertAfter(i, cur)
		cur = head
		curIndex = 0
	}
	result := make([][]int, len(people))
	for i, cur := 0, head.Next(); cur != nil; i, cur = i+1, cur.Next() {
		result[i] = people[cur.Value.(int)]
	}
	return result
}

// 使用 append insert
func reconstructQueue_2(people [][]int) [][]int {
	// 先將身高高的人排在前面，同樣身高 k 值小的在前面
	sort.Slice(people, func(i, j int) bool {
		// 同樣身高 k 值小的在前面
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		// 身高高的在前面
		return people[i][0] > people[j][0]
	})

	// 接下來只要按照 k 值將它插入到對應的位置就行了
	result := [][]int{}
	for _, person := range people {
		idx := person[1]
		result = append(result[:idx], append([][]int{person}, result[idx:]...)...)
	}

	return result
}

// 使用 copy
func reconstructQueue_3(people [][]int) [][]int {
	// 先將身高高的人排在前面，同樣身高 k 值小的在前面
	sort.Slice(people, func(i, j int) bool {
		// 同樣身高 k 值小的在前面
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		// 身高高的在前面
		return people[i][0] > people[j][0]
	})

	// 接下來只要按照 k 值將它插入到對應的位置就行了
	for i, person := range people {
		idx := person[1]
		// 把 idx 到 i 範圍的值往後挪一格
		copy(people[idx+1:i+1], people[idx:i])
		people[idx] = person
	}

	return people
}
