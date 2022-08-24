package p0200

func numIslands(grid [][]byte) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				fill(grid, i, j)
				cnt++
			}
		}
	}
	return cnt
}

func fill(grid [][]byte, row int, column int) {
	if row < 0 || column < 0 || row >= len(grid) || column >= len(grid[0]) || grid[row][column] == '0' {
		return
	}

	grid[row][column] = '0'

	fill(grid, row+1, column)
	fill(grid, row-1, column)
	fill(grid, row, column+1)
	fill(grid, row, column-1)
}
