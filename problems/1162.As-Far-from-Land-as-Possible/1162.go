package p1162

import "container/list"

// dfs
func maxDistance_dfs(grid [][]int) int {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] == 1 {
				grid[row][column] = 0
				dfs(grid, row, column, 1)
			}
		}
	}

	maxDistance := -1
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] > 1 {
				maxDistance = max(maxDistance, grid[row][column]-1)
			}
		}
	}
	return maxDistance
}

func dfs(grid [][]int, row int, column int, dist int) {
	if row < 0 || column < 0 || row >= len(grid) || column >= len(grid[0]) || (grid[row][column] != 0 && grid[row][column] <= dist) {
		return
	}

	grid[row][column] = dist

	dfs(grid, row+1, column, dist+1)
	dfs(grid, row-1, column, dist+1)
	dfs(grid, row, column+1, dist+1)
	dfs(grid, row, column-1, dist+1)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// bfs
func maxDistance(grid [][]int) int {
	queue := list.New()

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] == 1 {
				queue.PushBack([]int{row + 1, column})
				queue.PushBack([]int{row - 1, column})
				queue.PushBack([]int{row, column + 1})
				queue.PushBack([]int{row, column - 1})
			}
		}
	}

	step := 0
	for queue.Len() != 0 {
		step++
		qLen := queue.Len()
		for i := 0; i < qLen; i++ {
			front := queue.Front()
			d := front.Value.([]int)
			row, column := d[0], d[1]
			queue.Remove(front)
			if row >= 0 && column >= 0 && row < len(grid) && column < len(grid[0]) && grid[row][column] == 0 {
				grid[row][column] = step
				queue.PushBack([]int{row + 1, column})
				queue.PushBack([]int{row - 1, column})
				queue.PushBack([]int{row, column + 1})
				queue.PushBack([]int{row, column - 1})
			}
		}
	}

	if step <= 1 {
		return -1
	}

	return step - 1
}
