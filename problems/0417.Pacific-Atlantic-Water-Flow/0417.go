package p0417

// 分別從鄰近 Pacific 和 Atlantic 的 land 做 dfs
func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])

	pacific, atlantic := makeMatrix(m, n), makeMatrix(m, n)

	for row := 0; row < m; row++ {
		dfs(heights, pacific, 0, row, 0)
		dfs(heights, atlantic, 0, row, n-1)
	}

	for column := 0; column < n; column++ {
		dfs(heights, pacific, 0, 0, column)
		dfs(heights, atlantic, 0, m-1, column)
	}

	ans := [][]int{}
	for row := 0; row < m; row++ {
		for column := 0; column < n; column++ {
			if pacific[row][column] && atlantic[row][column] {
				ans = append(ans, []int{row, column})
			}
		}
	}
	return ans
}

func dfs(heights [][]int, visited [][]bool, height int, row int, column int) {
	m, n := len(heights), len(heights[0])
	if row < 0 || column < 0 || row >= m || column >= n || visited[row][column] || heights[row][column] < height {
		return
	}

	visited[row][column] = true

	dfs(heights, visited, heights[row][column], row+1, column)
	dfs(heights, visited, heights[row][column], row-1, column)
	dfs(heights, visited, heights[row][column], row, column+1)
	dfs(heights, visited, heights[row][column], row, column-1)
}

func makeMatrix(m, n int) [][]bool {
	arr := make([][]bool, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]bool, n)
	}
	return arr
}
