package p0997

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	n     int
	trust [][]int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{2, [][]int{{1, 2}}}, 2},
		{Input{3, [][]int{{1, 3}, {2, 3}}}, 3},
		{Input{3, [][]int{{1, 3}, {2, 3}, {3, 1}}}, -1},
	}
}

func TestFindJudge(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findJudge(testCase.Input.n, testCase.Input.trust)
		assert.Equal(t, testCase.Output, output)
	}
}
