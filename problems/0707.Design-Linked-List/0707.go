package p0707

type ListNode struct {
	val  int
	next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{size: 0, head: &ListNode{}}
}

func NewListNode(val int) ListNode {
	return ListNode{val: val}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i < index+1; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// if index is greater than the length,
	// the node will not be inserted.
	if index > this.size {
		return
	}
	// the node will be inserted at the head of the list.
	if index < 0 {
		index = 0
	}
	this.size++
	prev := this.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	to_add := NewListNode(val)
	to_add.next = prev.next
	prev.next = &to_add
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	this.size--
	prev := this.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
