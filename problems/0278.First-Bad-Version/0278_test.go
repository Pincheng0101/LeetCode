package p0278

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	n   int
	bad int
}

type testCase struct {
	Input  Input
	Output int
}

func testCases() []testCase {
	return []testCase{
		{Input{5, 4}, 4},
		{Input{1, 1}, 1},
	}
}

func TestSearch(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		oldValue := FirstBadVersion
		defer func() {
			FirstBadVersion = oldValue
		}()
		FirstBadVersion = testCase.Input.bad

		output := firstBadVersion(testCase.Input.n)
		assert.Equal(t, testCase.Output, output)
	}
}
