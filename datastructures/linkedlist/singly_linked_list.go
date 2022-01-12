package linkedlist

type SinglyLinkedList struct {
	Head *ListNode
	Size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		Head: NewListNode(0),
		Size: 0,
	}
}

func (this *SinglyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	cur := this.Head
	for i := 0; i < index+1; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *SinglyLinkedList) GetNode(index int) *ListNode {
	if index < 0 || index >= this.Size {
		return nil
	}
	cur := this.Head
	for i := 0; i < index+1; i++ {
		cur = cur.Next
	}
	return cur
}

func (this *SinglyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *SinglyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *SinglyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}
	if index < 0 {
		index = 0
	}
	cur := this.Head
	to_add := NewListNode(val)
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	to_add.Next = cur.Next
	cur.Next = to_add
	this.Size++
}

func (this *SinglyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.Size--
}

func (this *SinglyLinkedList) AddValuesWithSlices(values []int) {
	cur := this.Head
	for _, v := range values {
		cur.Next = NewListNode(v)
		cur = cur.Next
	}
	this.Size += len(values)
}
