package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *safeCounter) increment(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *safeCounter) value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := safeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.increment("hello")
	}

	time.Sleep(time.Second)
	fmt.Println(c.value("hello"))
}
