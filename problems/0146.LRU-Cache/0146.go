package p0146

import "container/list"

type item struct {
	key   int
	value int
}

type LRUCache struct {
	capacity  int
	itemsList *list.List
	itemsMap  map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		itemsList: list.New(),
		itemsMap:  map[int]*list.Element{},
	}
}

func (this *LRUCache) Get(key int) int {
	if element, ok := this.itemsMap[key]; ok {
		this.itemsList.MoveToFront(element)
		return element.Value.(item).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if element, ok := this.itemsMap[key]; ok {
		element.Value = item{
			key: key, value: value,
		}
		this.itemsList.MoveToFront(element)
		return
	}
	newElement := this.itemsList.PushFront(item{
		key: key, value: value,
	})
	this.itemsMap[key] = newElement
	if this.itemsList.Len() > this.capacity {
		removeElement := this.itemsList.Back()
		this.itemsList.Remove(removeElement)
		delete(this.itemsMap, removeElement.Value.(item).key)
	}
}

/**
 * Your LRUitemsMap object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
