package p0695

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] == 1 {
				area := areaSize(grid, row, column)
				if area > maxArea {
					maxArea = area
				}
			}

		}
	}
	return maxArea
}

func areaSize(grid [][]int, row int, column int) int {
	if row < 0 || column < 0 || row >= len(grid) || column >= len(grid[0]) || grid[row][column] == 0 {
		return 0
	}

	grid[row][column] = 0

	size := 1
	size += areaSize(grid, row+1, column)
	size += areaSize(grid, row-1, column)
	size += areaSize(grid, row, column+1)
	size += areaSize(grid, row, column-1)
	return size
}
