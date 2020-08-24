package main

import (
	"fmt"

	"github.com/karnadattasai/Cache-Go/service/cache"
)

func main() {
	var c cache.Cache
	c = cache.NewLFUCache()
	fmt.Println(c.Read(1))
	c.Write(7, 7)
	c.Display()
	c.Write(0, 0)
	c.Display()
	c.Write(1, 1)
	c.Display()
	c.Write(2, 2)
	c.Display()
	c.Write(0, 0)
	c.Display()
	// fmt.Println(c.Read(0))
	c.Write(3, 3)
	c.Display()
	c.Write(0, 0)
	c.Display()
	c.Write(4, 4)
	c.Display()
	c.Write(2, 2)
	c.Display()
	c.Write(3, 3)
	c.Display()
	c.Write(0, 0)
	c.Display()
	c.Write(3, 3)
	c.Display()
	c.Write(2, 2)
	c.Display()
	c.Write(1, 1)
	c.Display()
	c.Write(2, 2)
	c.Display()
	c.Write(0, 0)

	fmt.Println(c.Read(0))

}
