package main

import (
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]bool
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) setKey(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key] = true
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
