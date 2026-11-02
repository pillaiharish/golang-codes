package main

import (
	"context"
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	broken := flag.Bool("leak", false, "intentionally wait on a worker blocked forever")
	flag.Parse()
	baseline := runtime.NumGoroutine()
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	if *broken {
		go func() { defer wg.Done(); ch <- 1 }()
		wg.Wait() // intentional deadlock: nobody receives
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		select {
		case ch <- 1:
		case <-ctx.Done():
			fmt.Println("blocked sender canceled")
		}
	}()
	<-time.After(100 * time.Millisecond)
	cancel()
	wg.Wait()
	fmt.Printf("goroutines before: %d, after joining: %d\n", baseline, runtime.NumGoroutine())
}
