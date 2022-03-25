package p0062

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	m int
	n int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{3, 7}, 28},
		{Input{3, 2}, 3},
	}
}

func TestUniquePaths(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := uniquePaths(testCase.Input.m, testCase.Input.n)
		assert.Equal(t, testCase.Output, output)
	}
}
