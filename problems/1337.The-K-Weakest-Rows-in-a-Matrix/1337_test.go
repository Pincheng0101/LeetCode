package p1337

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	mat [][]int
	k   int
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{
			Input{[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			}, 3}, []int{2, 0, 3},
		},
		{
			Input{[][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			}, 2}, []int{0, 2},
		},
	}
}

func TestKWeakestRows(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := kWeakestRows(testCase.Input.mat, testCase.Input.k)
		assert.Equal(t, testCase.Output, output)
	}
}
