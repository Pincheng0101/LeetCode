package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2) // linked list becomes 1->2->3
	assert.Equal(t, list.size, 3)
	assert.Equal(t, list.Get(0), 1)
	assert.Equal(t, list.Get(1), 2)
	assert.Equal(t, list.Get(2), 3)
	list.DeleteAtIndex(1) // now the linked list is 1->3
	assert.Equal(t, list.size, 2)
	assert.Equal(t, list.Get(1), 3)
}
