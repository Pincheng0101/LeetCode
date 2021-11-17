package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	var testCases = []struct {
		Input  []int
		Output int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 0, 1, 1, 0, 1}, 2},
	}

	for _, testCase := range testCases {
		output := findMaxConsecutiveOnes(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
