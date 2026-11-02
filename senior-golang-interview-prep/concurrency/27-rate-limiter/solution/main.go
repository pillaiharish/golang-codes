package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const total = 20
	ticker := time.NewTicker(200 * time.Millisecond) // pace: five starts/second
	defer ticker.Stop()
	permits := make(chan struct{}, 3) // concurrency bound is separate
	var wg sync.WaitGroup

	for id := 0; id < total; id++ {
		<-ticker.C
		permits <- struct{}{}
		wg.Add(1)
		go func(task int) {
			defer wg.Done()
			defer func() { <-permits }()
			time.Sleep(300 * time.Millisecond)
			fmt.Println("completed:", task)
		}(id)
	}
	wg.Wait()
	fmt.Println("all tasks completed")
}
