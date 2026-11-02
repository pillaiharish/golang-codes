package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sem := make(chan struct{}, 3)
	var wg sync.WaitGroup
	var mu sync.Mutex
	active, maxActive := 0, 0

	for task := 0; task < 12; task++ {
		sem <- struct{}{} // bound admission BEFORE creating the goroutine
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			active++
			if active > maxActive { maxActive = active }
			mu.Unlock()
			defer func() {
				mu.Lock()
				active--
				mu.Unlock()
				<-sem
			}()
			time.Sleep(30 * time.Millisecond)
			fmt.Println("completed task:", id)
		}(task)
	}
	wg.Wait()
	fmt.Println("maximum active:", maxActive)
}
