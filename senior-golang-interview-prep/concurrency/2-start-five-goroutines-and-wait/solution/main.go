package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Register all five tasks before starting any of them.
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// Pass the index explicitly so each invocation gets its own value.
		go func(index int) {
			defer wg.Done()
			fmt.Println("Goroutine:", index)
		}(i)
	}

	// This joins all workers; it does not order their print statements.
	wg.Wait()
	fmt.Println("All goroutines completed")
}
