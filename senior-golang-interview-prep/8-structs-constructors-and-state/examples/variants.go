package main

import "sync"

// LazyCounter illustrates a type whose zero value is usable for writes.
// It is not safe for concurrent use.
type LazyCounter struct {
	counts map[string]int
}

func (c *LazyCounter) Increment(key string) {
	if c.counts == nil {
		c.counts = make(map[string]int)
	}
	c.counts[key]++
}

func (c *LazyCounter) GetKey(key string) int {
	return c.counts[key]
}

// Gauge illustrates the difference between changing a receiver copy
// and changing the original value through a pointer.
type Gauge struct {
	value int
}

func (g Gauge) TrySetByValue(n int) {
	g.value = n
}

func (g *Gauge) SetByPointer(n int) {
	g.value = n
}

// SafeCounter makes the stateful counter usable concurrently.
// Do not copy this struct after its mutex has been used.
type SafeCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{counts: make(map[string]int)}
}

func (c *SafeCounter) Increment(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++
}

func (c *SafeCounter) GetKey(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[key]
}
