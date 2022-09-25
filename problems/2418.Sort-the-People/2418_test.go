package p2418

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	names   []string
	heights []int
}

type testCase struct {
	Input  Input
	Output []string
}

func testCases() []testCase {
	return []testCase{
		{Input{[]string{"Mary", "John", "Emma"}, []int{180, 165, 170}}, []string{"Mary", "Emma", "John"}},
		{Input{[]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}}, []string{"Bob", "Alice", "Bob"}},
	}
}

func TestSortPeople(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := sortPeople(testCase.Input.names, testCase.Input.heights)
		assert.Equal(t, testCase.Output, output)
	}
}
