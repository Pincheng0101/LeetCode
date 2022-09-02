package p0542

import (
	"container/list"
	"math"
)

type Point struct {
	x int
	y int
}

func updateMatrix(mat [][]int) [][]int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	m, n := len(mat), len(mat[0])

	queue := list.New()

	for row := 0; row < m; row++ {
		for column := 0; column < n; column++ {
			if mat[row][column] == 0 {
				queue.PushBack(Point{x: row, y: column})
			} else {
				mat[row][column] = math.MaxInt64
			}
		}
	}

	for queue.Len() != 0 {
		qLen := queue.Len()
		for i := 0; i < qLen; i++ {
			front := queue.Front()
			d := front.Value.(Point)
			x, y := d.x, d.y
			queue.Remove(front)

			for _, direction := range directions {
				nx, ny := direction[0]+x, direction[1]+y
				if nx < 0 || nx >= m || ny < 0 || ny >= n || mat[nx][ny] <= mat[x][y] {
					continue
				}
				mat[nx][ny] = mat[x][y] + 1
				queue.PushBack(Point{x: nx, y: ny})
			}
		}
	}

	return mat
}
