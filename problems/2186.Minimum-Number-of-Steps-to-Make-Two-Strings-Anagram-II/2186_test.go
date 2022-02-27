package p2186

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	s string
	t string
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{"leetcode", "coats"}, 7},
		{Input{"night", "thing"}, 0},
	}
}

func TestMinSteps(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := minSteps(testCase.Input.s, testCase.Input.t)
		assert.Equal(t, testCase.Output, output)
	}
}
