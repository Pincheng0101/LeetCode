package p0200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  [][]byte
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'}}, 1},
		{[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'}}, 3},
	}
}

func TestNumIslands(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := numIslands(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
