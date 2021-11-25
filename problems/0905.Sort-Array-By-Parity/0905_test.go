package p0905

import (
	"reflect"
	"testing"
)

type testCase struct {
	Input          []int
	AcceptedOutput [][]int
}

func testCases() []testCase {
	return []testCase{
		{
			Input:          []int{3, 1, 2, 4},
			AcceptedOutput: [][]int{{2, 4, 3, 1}, {4, 2, 3, 1}, {2, 4, 1, 3}, {4, 2, 1, 3}},
		},
		{
			Input:          []int{0},
			AcceptedOutput: [][]int{{0}},
		},
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func TestSortArrayByParity(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		output := sortArrayByParity(testCase.Input)
		failed := true
		for _, v := range testCase.AcceptedOutput {
			if reflect.DeepEqual(v, output) {
				t.Log("success")
				failed = false
				break
			}
		}
		if failed {
			t.Log("fail")
		}
	}
}
