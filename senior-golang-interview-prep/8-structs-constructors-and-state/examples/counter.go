package main

// Counter stores independent per-key counts in an initialized map.
// Its zero value can be read but cannot be incremented until initialized.
type Counter struct {
	counts map[string]int
}

// NewCounter is an ordinary function, not a special Go constructor.
// Returning a value works here; callers can invoke pointer-receiver methods
// on their addressable Counter variables.
func NewCounter() Counter {
	return Counter{counts: make(map[string]int)}
}

func (c *Counter) Increment(key string) {
	c.counts[key]++
}

func (c *Counter) GetKey(key string) int {
	return c.counts[key]
}
