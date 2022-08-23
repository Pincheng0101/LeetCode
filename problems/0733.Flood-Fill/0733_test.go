package p0733

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	image [][]int
	sr    int
	sc    int
	color int
}

type testCase struct {
	Input  Input
	Output [][]int
}

func testCases() []testCase {
	return []testCase{
		{Input{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2}, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{Input{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0}, [][]int{{0, 0, 0}, {0, 0, 0}}},
	}
}

func TestFloodFill(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := floodFill(testCase.Input.image, testCase.Input.sr, testCase.Input.sc, testCase.Input.color)
		assert.Equal(t, testCase.Output, output)
	}
}
