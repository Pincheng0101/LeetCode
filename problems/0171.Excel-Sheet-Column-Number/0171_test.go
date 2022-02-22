package p0171

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  string
	Output int
}

func testCases() []testCase {
	return []testCase{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}
}

func TestTitleToNumber(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := titleToNumber(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
