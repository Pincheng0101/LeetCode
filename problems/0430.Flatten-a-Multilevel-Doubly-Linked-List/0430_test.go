package p0430

import (
	"testing"
)

type testCase struct {
	Input  []interface{}
	Output []int
}

func testCases() []testCase {
	return []testCase{
		{[]interface{}{1, 2, 3, 4, 5, 6, nil, nil, nil, 7, 8, 9, 10, nil, nil, 11, 12}, []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6}},
		{[]interface{}{1, 2, nil, 3}, []int{1, 3, 2}},
		{[]interface{}{}, []int{}},
	}
}

// TODO: TestFlatten
func TestFlatten(t *testing.T) {

}
