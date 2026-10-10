package main

import (
	"fmt"
	"sync"
	"time"
)

// Using sync.Map:
// Go also provides a concurrent-safe map implementation out of the box
// in the sync package. It uses special internal optimizations for concurrent loads,
// stores, and deletes (best used when multiple goroutines read/write disjoint keys
// or write once and read many times)

func main() {

	var m sync.Map
	go func() {
		for i := range 1000 {
			m.Store("a", i)
		}
	}()
	go func() {
		for i := range 1000 {
			m.Store("b", i)
		}
	}()

	time.Sleep(1 * time.Second)

	val_a, ok := m.Load("a")
	fmt.Println("Value of key a:", val_a, ok)

	// But if we want to iterate through the entire map stored in sync.Map then
	// use m.Range. To iterate over all keys and values in a sync.Map,
	// we cannot use a standard for ... range loop (Go will throw a compilation
	// error if we try to range over a sync.Map). Instead, sync.Map provides a
	// built-in method called Range. It accepts a callback function that runs
	// for every key-value pair stored in the map.
	m.Range(func(key, value any) bool {
		fmt.Println("Key:", key, "value:", value)
		return true
	})
}

// Output
// harish $ go run -race main.go
// Value of key a: 999 true
// Key: b value: 999
// Key: a value: 999
