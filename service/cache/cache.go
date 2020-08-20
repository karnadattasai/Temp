//Package cache Provides LRU, LFU, FIFU Cache
package cache

//capacity is the maximum size of the Cache
var capacity = 10

// Cache interface contains methods read, write and new methods
type Cache interface {
	Read(key int) (value int) // reads and returns a value at given key
	Write(key int, value int) // writes the given value at given key
}

// NewLRUCache return new Cache with LRU as replacement policy
func NewLRUCache() Cache {
}

// NewLFUCache return new Cache with LFU as replacement policy
func NewLFUCache() Cache {
}

// NewFIFOCache return new Cache with FIFO as replacement policy
func NewFIFOCache() Cache {
}
