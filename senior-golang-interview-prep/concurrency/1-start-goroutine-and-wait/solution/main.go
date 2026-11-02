package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Register the task before the new goroutine can run.
	wg.Add(1)
	go func() {
		// Signal completion even if the function returns early.
		defer wg.Done()
		fmt.Println("Hello from goroutine")
	}()

	// main must not exit until the registered task finishes.
	wg.Wait()
	fmt.Println("Main finished")
}
