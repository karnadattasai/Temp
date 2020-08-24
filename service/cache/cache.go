//Package cache Provides LRU, LFU, FIFU Cache
package cache

import (
	"container/heap"
	"fmt"

	"github.com/karnadattasai/Cache-Go/service/list"
)

//capacity is the maximum size of the Cache
var capacity = 3

// Cache interface contains methods read, write and new methods
type Cache interface {
	Read(key int) (value int) // reads and returns a value at given key
	Write(key int, value int) // writes the given value at given key
	Display()
}

// NewLRUCache return new Cache with LRU as replacement policy
func NewLRUCache() Cache {
	c := cacheLRU{}
	c.m = make(map[int]*list.Node)
	return &c
}

// NewLFUCache return new Cache with LFU as replacement policy
func NewLFUCache() Cache {
	pq := make(priorityQueue, 0)
	fmt.Printf("Hi")
	// pq[0].freq = 0
	// fmt.Printf("Hi")
	// pq[0].index = 0
	// pq[0].p = pair{0, 0}
	// pq[1].freq = 0
	// pq[1].index = 1
	// pq[1].p = pair{0, 0}

	fmt.Printf("Hi")

	heap.Init(&pq)
	m := make(map[int]*pqNode)
	return &cacheLFU{m, pq}
}

// // NewFIFOCache return new Cache with FIFO as replacement policy
// func NewFIFOCache() Cache {
// }
