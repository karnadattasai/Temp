package cache

import (
	"github.com/karnadattasai/Cache-Go/service/list"
)

// cacheLRU hold data structures that implements LRU
type cacheLRU struct {
	cacheList         list.List // in the order from lru to most recently used
	keyNodePointerMap map[int]*list.Node
}

func (c *cacheLRU) Read(key int) int {
	if node, ok := c.keyNodePointerMap[key]; ok { // check if key is present
		c.cacheList.MoveBack(node)
		return node.P.Value
	}
	return -1
}

func (c *cacheLRU) Write(key int, value int) {
	// if given key is already presesnt update the value and move the node to back of the list as it is the most recently used node
	if node, ok := c.keyNodePointerMap[key]; ok {
		c.cacheList.MoveBack(node)
		node.P.Value = value
		return
	}
	// if key not present, first check if the length of list is less than capacity of cache else remove the LRU node
	if c.cacheList.Len() >= capacity {
		delete(c.keyNodePointerMap, c.cacheList.Front().P.Key)
		c.cacheList.Remove(c.cacheList.Front())
	}
	// Add the node to list and insert the key in the map
	c.keyNodePointerMap[key] = c.cacheList.Push(list.Pair{Key: key, Value: value})
}
