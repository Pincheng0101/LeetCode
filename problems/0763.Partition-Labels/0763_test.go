package p0763

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"eccbbbbdec", []int{10}},
	}
}

func TestPartitionLabels(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := partitionLabels(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
