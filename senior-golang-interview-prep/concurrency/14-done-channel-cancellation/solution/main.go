package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ticker := time.NewTicker(50 * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-done:
					fmt.Println("stopped:", id)
					return
				case <-ticker.C:
					fmt.Println("working:", id)
				}
			}
		}(i)
	}
	// A timer simulates an external stop request; WaitGroup handles completion.
	<-time.After(160 * time.Millisecond)
	close(done)
	wg.Wait()
	fmt.Println("all workers joined")
}
