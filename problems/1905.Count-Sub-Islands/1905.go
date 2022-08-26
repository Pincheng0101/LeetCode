package p1905

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	cnt := 0
	for row := 0; row < len(grid2); row++ {
		for column := 0; column < len(grid2[0]); column++ {
			if grid2[row][column] == 1 && dfs(grid1, grid2, row, column) {
				cnt++
			}
		}
	}
	return cnt
}

func dfs(grid1 [][]int, grid2 [][]int, row int, column int) bool {
	if row < 0 || column < 0 || row >= len(grid2) || column >= len(grid2[0]) || grid2[row][column] == 0 {
		return true
	}

	isSubIsland := true

	if grid1[row][column] == 0 {
		isSubIsland = false
	}

	grid2[row][column] = 0

	isSubIsland = dfs(grid1, grid2, row+1, column) && isSubIsland
	isSubIsland = dfs(grid1, grid2, row-1, column) && isSubIsland
	isSubIsland = dfs(grid1, grid2, row, column+1) && isSubIsland
	isSubIsland = dfs(grid1, grid2, row, column-1) && isSubIsland
	return isSubIsland
}
