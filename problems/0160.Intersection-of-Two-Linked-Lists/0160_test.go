package p0160

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pincheng0101/leetcode/datastructures/linkedlist"
)

type Input struct {
	intersectVal int
	listA        []int
	listB        []int
	skipA        int
	skipB        int
}
type testCase struct {
	Input Input
	// Output: intersected *ListNode
}

func testCases() []testCase {
	return []testCase{
		{Input{
			intersectVal: 8,
			listA:        []int{4, 1, 8, 4, 5},
			listB:        []int{5, 6, 1, 8, 4, 5},
			skipA:        2,
			skipB:        3,
		}},
		{Input{
			intersectVal: 2,
			listA:        []int{1, 9, 1, 2, 4},
			listB:        []int{3, 2, 4},
			skipA:        3,
			skipB:        1,
		}},
		{Input{
			intersectVal: 0,
			listA:        []int{2, 6, 4},
			listB:        []int{1, 5},
			skipA:        3,
			skipB:        2,
		}},
	}
}

func TestGetIntersectionNode_1(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		listA := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.listA); i++ {
			listA.AddAtTail(testCase.Input.listA[i])
		}

		listB := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.listB); i++ {
			listB.AddAtTail(testCase.Input.listB[i])
		}

		listASkipNode := listA.GetNode(testCase.Input.skipA - 1)
		listBSkipNode := listB.GetNode(testCase.Input.skipB - 1)
		listBSkipNode.Next = listASkipNode.Next

		output := getIntersectionNode_1(listA.Head, listB.Head)
		assert.Equal(t, listASkipNode.Next, output)
	}
}

func TestGetIntersectionNode_2(t *testing.T) {
	testCases := testCases()
	for _, testCase := range testCases {
		listA := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.listA); i++ {
			listA.AddAtTail(testCase.Input.listA[i])
		}

		listB := NewSinglyLinkedList()
		for i := 0; i < len(testCase.Input.listB); i++ {
			listB.AddAtTail(testCase.Input.listB[i])
		}

		listASkipNode := listA.GetNode(testCase.Input.skipA - 1)
		listBSkipNode := listB.GetNode(testCase.Input.skipB - 1)
		listBSkipNode.Next = listASkipNode.Next

		output := getIntersectionNode_2(listA.Head, listB.Head)
		assert.Equal(t, listASkipNode.Next, output)
	}
}
