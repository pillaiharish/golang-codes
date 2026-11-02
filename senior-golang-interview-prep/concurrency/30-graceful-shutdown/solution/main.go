package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	jobs := make(chan int, 4)
	producerDone := make(chan struct{})
	go func() {
		defer close(producerDone)
		defer close(jobs) // sole producer is the only closer
		ticker := time.NewTicker(40 * time.Millisecond)
		defer ticker.Stop()
		for id := 1; ; id++ {
			select {
			case <-signalCtx.Done(): return // stop admission
			case <-ticker.C:
				select {
				case jobs <- id:
				case <-signalCtx.Done(): return
				}
			}
		}
	}()

	var workers sync.WaitGroup
	for i := 0; i < 3; i++ {
		workers.Add(1)
		go func(workerID int) {
			defer workers.Done()
			for job := range jobs { // drains accepted work after cancellation
				time.Sleep(75 * time.Millisecond)
				fmt.Println("worker", workerID, "finished job", job)
			}
		}(i)
	}
	// Automatically send a stop request for the demo; Ctrl-C works too.
	go func() { time.Sleep(350 * time.Millisecond); stop() }()
	<-signalCtx.Done()
	<-producerDone // wait until the producer closes admission/jobs

	joined := make(chan struct{})
	go func() { workers.Wait(); close(joined) }()
	select {
	case <-joined:
		fmt.Println("shutdown complete")
	case <-time.After(2 * time.Second):
		fmt.Println("drain deadline exceeded; some work may remain")
	}
}
