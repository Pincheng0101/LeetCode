package p2217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	queries   []int
	intLength int
}

type testCase struct {
	Input  Input
	Output []int64
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{1, 2, 3, 4, 5, 90}, 3}, []int64{101, 111, 121, 131, 141, 999}},
		{Input{[]int{2, 4, 6}, 4}, []int64{1111, 1331, 1551}},
	}
}

func TestKthPalindrome(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := kthPalindrome(testCase.Input.queries, testCase.Input.intLength)
		assert.Equal(t, testCase.Output, output)
	}
}
