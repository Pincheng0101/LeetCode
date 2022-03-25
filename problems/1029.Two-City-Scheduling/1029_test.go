package p1029

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  [][]int
	Output int
}

func testCases() []testCase {
	return []testCase{
		{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, 110},
		{[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}, 1859},
		{[][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}}, 3086},
	}
}

func TestTwoCitySchedCost(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := twoCitySchedCost(testCase.Input)
		assert.Equal(t, testCase.Output, output)
	}
}
