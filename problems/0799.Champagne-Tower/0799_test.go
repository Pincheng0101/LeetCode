package p0799

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	poured      int
	query_row   int
	query_glass int
}

type testCase struct {
	Input  Input
	Output float64
}

func testCases() []testCase {
	return []testCase{
		{Input{1, 1, 1}, 0.00000},
		{Input{2, 1, 1}, 0.50000},
		{Input{100000009, 33, 17}, 1.00000},
	}
}

func TestChampagneTower(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := champagneTower(testCase.Input.poured, testCase.Input.query_row, testCase.Input.query_glass)
		assert.Equal(t, testCase.Output, output)
	}
}
