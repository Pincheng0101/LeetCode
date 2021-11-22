package p1346

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
		{[]int{10, 2, 5, 3}, true},
		{[]int{7, 1, 14, 11}, true},
		{[]int{3, 1, 7, 11}, false},
	}
}

func TestCheckIfExist(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := checkIfExist(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
