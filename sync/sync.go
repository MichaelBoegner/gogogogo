package sync

import "sync"

type Counter struct {
	mu     sync.Mutex
	amount int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.amount += 1
}

func (c *Counter) Value() int {
	return c.amount
}
