package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		panic("something went wrong") // Intentional process failure.
	}()
	wg.Wait()
}
