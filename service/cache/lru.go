package cache

import (
	"github.com/karnadattasai/Cache-Go/service/list"
)

// cacheLRU hold data structures that implements LRU
type cacheLRU struct {
	l list.List
	m map[int]*list.Node
}

func (c *cacheLRU) Display() {
	c.l.Display()
}

func (c *cacheLRU) Read(key int) int {
	if node, ok := c.m[key]; ok {
		c.l.MoveBack(node)
		return node.P.Value
	}
	return -1
}

func (c *cacheLRU) Write(key int, value int) {
	if node, ok := c.m[key]; ok {
		c.l.MoveBack(node)
		node.P.Value = value
		return
	}
	// fmt.Println("hi ", c.l.Len(), key)
	if c.l.Len() >= capacity {
		delete(c.m, c.l.Front().P.Key)
		c.l.Remove(c.l.Front())
		// fmt.Println("hi ", c.l.Len(), key)
	}
	c.m[key] = c.l.Push(list.Pair{Key: key, Value: value})

}
