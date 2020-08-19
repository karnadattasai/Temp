package cache

// Cache interface contains methods read, write and new methods
type Cache interface {
	Read(address int) (value int) // reads and returns a value at given address
	Write(address int, value int) // writes the given value at given address
	Initialize(size int)          // initializes cache
}

// base is a generic type cache that holds data and size
type base struct {
}

// LRU is on top of base with additional data structures that implements LRU
type LRU struct {
}

// LFU is on top of base with additional data structures that implements LRU
type LFU struct {
}

// FIFO is on top of base with additional data structures that implements LRU
type FIFO struct {
}
