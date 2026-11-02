package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Config struct{ Name string }

func main() {
	var once sync.Once
	var calls atomic.Int64
	var instance *Config
	getInstance := func() *Config {
		once.Do(func() {
			calls.Add(1)
			instance = &Config{Name: "loaded once"}
		})
		return instance
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if getInstance().Name != "loaded once" {
				panic("unexpected configuration")
			}
		}()
	}
	wg.Wait()
	fmt.Println("initializer calls:", calls.Load())
}
