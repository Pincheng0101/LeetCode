package p0146

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)                   // cache is {1=1}
	lruCache.Put(2, 2)                   // cache is {1=1, 2=2}
	assert.Equal(t, 1, lruCache.Get(1))  // return 1
	lruCache.Put(3, 3)                   // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	assert.Equal(t, -1, lruCache.Get(2)) // returns -1 (not found)
	lruCache.Put(4, 4)                   // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	assert.Equal(t, -1, lruCache.Get(1)) // return -1 (not found)
	assert.Equal(t, 3, lruCache.Get(3))  // return 3
	assert.Equal(t, 4, lruCache.Get(4))  // return 4

}
