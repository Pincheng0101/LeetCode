package p0165

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	version1 string
	version2 string
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{"1.01", "1.001"}, 0},
		{Input{"1.0", "1.0.0"}, 0},
		{Input{"0.1", "1.1"}, -1},
	}
}

func TestCompareVersion(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := compareVersion(testCase.Input.version1, testCase.Input.version2)
		assert.Equal(t, testCase.Output, output)
	}
}
