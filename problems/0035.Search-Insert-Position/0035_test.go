package p0035

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	Nums   []int
	Target int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{Nums: []int{1, 3, 5, 6}, Target: 5}, 2},
		{Input{Nums: []int{1, 3, 5, 6}, Target: 2}, 1},
		{Input{Nums: []int{1, 3, 5, 6}, Target: 7}, 4},
		{Input{Nums: []int{1, 3, 5, 6}, Target: 0}, 0},
		{Input{Nums: []int{1, 3, 5, 6}, Target: 1}, 0},
	}
}

func TestSearchInsert_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := searchInsert_1(testCase.Input.Nums, testCase.Input.Target)
		assert.Equal(t, testCase.Output, output)
	}
}

func TestSearchInsert_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := searchInsert_2(testCase.Input.Nums, testCase.Input.Target)
		assert.Equal(t, testCase.Output, output)
	}
}
