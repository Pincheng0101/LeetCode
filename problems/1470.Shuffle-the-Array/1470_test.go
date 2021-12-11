package p1470

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	Nums []int
	N    int
}

type testCase struct {
	Input  Input
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{Input{[]int{2, 5, 1, 3, 4, 7}, 3}, []int{2, 3, 5, 4, 1, 7}},
		{Input{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4}, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{Input{[]int{1, 1, 2, 2}, 2}, []int{1, 2, 1, 2}},
	}
}

func TestShuffle(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := shuffle(testCase.Input.Nums, testCase.Input.N)
		assert.Equal(t, testCase.Output, output)
	}
}
