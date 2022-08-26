package p1020

func numEnclaves(grid [][]int) int {
	rowSize := len(grid)
	columnSize := len(grid[0])

	for row := 0; row < rowSize; row++ {
		fill(grid, row, 0)
		fill(grid, row, columnSize-1)
	}

	for column := 0; column < columnSize; column++ {
		fill(grid, 0, column)
		fill(grid, rowSize-1, column)
	}

	cnt := 0
	for row := 1; row < rowSize-1; row++ {
		for column := 1; column < columnSize-1; column++ {
			if grid[row][column] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func fill(grid [][]int, row int, column int) {
	if row < 0 || column < 0 || row >= len(grid) || column >= len(grid[0]) || grid[row][column] == 0 {
		return
	}

	grid[row][column] = 0

	fill(grid, row+1, column)
	fill(grid, row-1, column)
	fill(grid, row, column+1)
	fill(grid, row, column-1)
}
