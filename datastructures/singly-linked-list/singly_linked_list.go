package singlylinkedlist

type ListNode struct {
	val  int
	next *ListNode
}

type SinglyLinkedList struct {
	head *ListNode
	size int
}

func NewListNode(val int) *ListNode {
	return &ListNode{val: val}
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: NewListNode(0),
		size: 0,
	}
}

func (this *SinglyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i < index+1; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *SinglyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *SinglyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *SinglyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	cur := this.head
	to_add := NewListNode(val)
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	to_add.next = cur.next
	cur.next = to_add
	this.size++
}

func (this *SinglyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	this.size--
}
