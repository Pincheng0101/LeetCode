package p0055

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  []int
	Output bool
}

func testCases() []testCase {
	return []testCase{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}
}

func TestCanJump(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := canJump(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
