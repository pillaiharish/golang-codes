package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	result := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(750 * time.Millisecond)
		result <- "finished"
		fmt.Println("worker continued after caller timeout")
	}()

	select {
	case v := <-result:
		fmt.Println("result:", v)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("operation timed out")
	}
	wg.Wait() // timeout did not cancel the worker
}
