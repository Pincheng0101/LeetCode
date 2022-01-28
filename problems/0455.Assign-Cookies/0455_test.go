package p0455

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	g []int
	s []int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 3}, []int{1, 1}}, 1},
		{Input{[]int{1, 2}, []int{1, 2, 3}}, 2},
	}
}

func TestFindContentChildren(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := findContentChildren(testCase.Input.g, testCase.Input.s)
		assert.Equal(t, testCase.Output, output)
	}
}
