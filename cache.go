package main

type Cache interface {
	read()
	write()
	new()
}
type CacheBase struct {
}
type CacheLRU struct {
}
type CacheLFU struct {
}
type CacheFIFO struct {
}

func main() {

}
