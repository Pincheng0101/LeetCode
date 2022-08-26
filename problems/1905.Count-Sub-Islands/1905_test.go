package p1905

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	grid1 [][]int
	grid2 [][]int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{
			[][]int{
				{1, 1, 1, 0, 0},
				{0, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 1, 1}},
			[][]int{
				{1, 1, 1, 0, 0},
				{0, 0, 1, 1, 1},
				{0, 1, 0, 0, 0},
				{1, 0, 1, 1, 0},
				{0, 1, 0, 1, 0}}}, 3},
		{Input{
			[][]int{
				{1, 0, 1, 0, 1},
				{1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1},
				{1, 0, 1, 0, 1}},
			[][]int{
				{0, 0, 0, 0, 0},
				{1, 1, 1, 1, 1},
				{0, 1, 0, 1, 0},
				{0, 1, 0, 1, 0},
				{1, 0, 0, 0, 1}}}, 2},
	}
}

func TestCountSubIslands(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := countSubIslands(testCase.Input.grid1, testCase.Input.grid2)
		assert.Equal(t, testCase.Output, output)
	}
}
