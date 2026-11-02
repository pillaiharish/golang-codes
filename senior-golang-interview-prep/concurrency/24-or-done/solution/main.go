package main

import (
	"context"
	"fmt"
	"sync"
)

func orDone(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done(): return
			case value, ok := <-in:
				if !ok { return }
				select {
				case out <- value:
				case <-ctx.Done(): return
				}
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	source := make(chan int)
	var producer sync.WaitGroup
	producer.Add(1)
	go func() {
		defer producer.Done()
		defer close(source)
		for i := 1; i <= 100; i++ {
			select {
			case source <- i:
			case <-ctx.Done(): return
			}
		}
	}()

	output := orDone(ctx, source)
	for v := range output {
		fmt.Println(v)
		if v == 3 { cancel(); break }
	}
	for range output {} // confirm the relay actually closed
	producer.Wait()
	fmt.Println("relay and producer stopped")
}
