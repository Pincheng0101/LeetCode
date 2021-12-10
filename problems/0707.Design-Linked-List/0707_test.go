package p0707

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyLinkedList1(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)
	assert.Equal(t, myLinkedList.Get(0), 1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2) // linked list becomes 1->2->3
	assert.Equal(t, myLinkedList.Get(0), 1)
	assert.Equal(t, myLinkedList.Get(1), 2)
	assert.Equal(t, myLinkedList.Get(2), 3)
	myLinkedList.DeleteAtIndex(1) // now the linked list is 1->3
	assert.Equal(t, myLinkedList.Get(1), 3)
}

func TestMyLinkedList2(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(7)
	assert.Equal(t, myLinkedList.Get(0), 7)
	myLinkedList.AddAtHead(2)
	assert.Equal(t, myLinkedList.Get(0), 2)
	myLinkedList.AddAtHead(1)
	assert.Equal(t, myLinkedList.Get(0), 1)
	myLinkedList.AddAtIndex(3, 0)
	assert.Equal(t, myLinkedList.Get(3), 0)
	myLinkedList.DeleteAtIndex(2)
	assert.Equal(t, myLinkedList.Get(2), 0)
	myLinkedList.AddAtHead(6)
	assert.Equal(t, myLinkedList.Get(0), 6)
	myLinkedList.AddAtTail(4)
	assert.Equal(t, myLinkedList.Get(0), 6)
	assert.Equal(t, myLinkedList.Get(1), 1)
	assert.Equal(t, myLinkedList.Get(2), 2)
	assert.Equal(t, myLinkedList.Get(3), 0)
	assert.Equal(t, myLinkedList.Get(4), 4)
	myLinkedList.AddAtHead(4)

	myLinkedList.AddAtIndex(5, 0)
	myLinkedList.AddAtHead(6)
}
