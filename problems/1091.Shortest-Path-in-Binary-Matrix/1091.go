package p1091

import "container/list"

type Point struct {
	x int
	y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	if grid[0][0] != 0 || grid[m-1][n-1] != 0 {
		return -1
	}

	directions := [][]int{{1, 1}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}, {-1, -1}, {1, -1}, {-1, 1}}

	queue := list.New()
	queue.PushBack(Point{x: 0, y: 0})

	grid[0][0] = 1

	for queue.Len() != 0 {
		front := queue.Front()
		d := front.Value.(Point)
		x, y := d.x, d.y
		queue.Remove(front)

		if x == m-1 && y == n-1 {
			return grid[x][y]
		}

		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]
			if nx >= 0 && ny >= 0 && nx < m && ny < n && grid[nx][ny] == 0 {
				queue.PushBack(Point{x: nx, y: ny})
				grid[nx][ny] = grid[x][y] + 1
			}
		}
	}

	return -1
}
