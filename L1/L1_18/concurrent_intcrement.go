package l118

import (
	"fmt"
	"sync"
)

type ConcurrentCounter struct {
	mu      sync.Mutex
	counter uint
}

func NewConcurrentCounter() *ConcurrentCounter {
	return &ConcurrentCounter{}
}

func (c *ConcurrentCounter) IncrementCounter() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *ConcurrentCounter) PrintCounter() {
	c.mu.Lock()
	fmt.Println(c.counter)
	c.mu.Unlock()
}
