package main

import (
	"context"
	"fmt"
	"sync"
)

func generate(ctx context.Context, wg *sync.WaitGroup, nums ...int) <-chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done(): return
			}
		}
	}()
	return out
}

func square(ctx context.Context, wg *sync.WaitGroup, in <-chan int) <-chan int {
	out := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)
		for {
			select {
			case <-ctx.Done(): return
			case n, ok := <-in:
				if !ok { return }
				select {
				case out <- n*n:
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
	var wg sync.WaitGroup
	out := square(ctx, &wg, generate(ctx, &wg, 1, 2, 3, 4, 5, 6))
	for i := 0; i < 3; i++ {
		fmt.Println("square:", <-out)
	}
	cancel() // printer stops early: signal every upstream stage
	wg.Wait()
	fmt.Println("pipeline stopped")
}
