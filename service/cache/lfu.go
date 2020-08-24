package cache

import (
	"container/heap"
	"fmt"
	"time"
)

type data struct {
	key       int
	value     int
	timestamp time.Time
}
type pqNode struct {
	p     data
	freq  int
	index int
}

type priorityQueue []*pqNode

func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].freq < pq[j].freq {
		return true
	}
	if pq[i].freq == pq[j].freq {
		fmt.Printf("Hello there")
		return pq[i].p.timestamp.Before(pq[j].p.timestamp)
	}
	return false
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqNode)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq *priorityQueue) update(item *pqNode, p data, freq int) {
	p.timestamp = time.Now()
	item.p = p
	item.freq = freq
	heap.Fix(pq, item.index)
}

// cacheLFU hold data structures that implements LFU
type cacheLFU struct {
	m  map[int]*pqNode
	pq priorityQueue
}

func (c *cacheLFU) Read(key int) int {
	if node, ok := c.m[key]; ok {
		c.pq.update(node, node.p, node.freq+1)
		return node.p.value
	}
	return -1
}

func (c *cacheLFU) Write(key, value int) {
	if node, ok := c.m[key]; ok {
		node.p.value = value
		c.pq.update(node, node.p, node.freq+1)
		return
	}
	if len(c.pq) >= capacity {
		node := heap.Pop(&c.pq).(*pqNode)
		delete(c.m, node.p.key)
	}
	node := &pqNode{data{key, value, time.Now()}, 1, 0}
	heap.Push(&c.pq, node)
	c.m[key] = node
}

func (c *cacheLFU) Display() {
	for i := 0; i < len(c.pq); i++ {
		fmt.Printf("%d ", c.pq[i].p.key)
	}
	fmt.Printf("\n")
	for i := 0; i < len(c.pq); i++ {
		fmt.Printf("%d ", c.pq[i].freq)
	}
	fmt.Printf("\n")
	for i := 0; i < len(c.pq); i++ {
		fmt.Printf("%v ", c.pq[i].p.timestamp)
	}
	fmt.Printf("\n")
	fmt.Printf("\n\n")
}
