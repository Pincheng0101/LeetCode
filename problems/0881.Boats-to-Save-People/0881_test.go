package p0881

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	people []int
	limit  int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2}, 3}, 1},
		{Input{[]int{3, 2, 2, 1}, 3}, 3},
		{Input{[]int{3, 5, 3, 4}, 5}, 4},
	}
}

func TestＮumRescueBoats(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := numRescueBoats(testCase.Input.people, testCase.Input.limit)
		assert.Equal(t, testCase.Output, output)
	}
}
