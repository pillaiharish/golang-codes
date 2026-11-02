package main

import (
	"context"
	"fmt"
	"sync"
)

func tee(ctx context.Context, input <-chan int) (<-chan int, <-chan int) {
	outA, outB := make(chan int), make(chan int)
	go func() {
		defer close(outA)
		defer close(outB)
		for {
			var v int
			select {
			case <-ctx.Done(): return
			case value, ok := <-input:
				if !ok { return }
				v = value
			}
			a, b := outA, outB
			for i := 0; i < 2; i++ {
				select {
				case a <- v: a = nil // disable an already-served branch
				case b <- v: b = nil
				case <-ctx.Done(): return
				}
			}
		}
	}()
	return outA, outB
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	in := make(chan int)
	a, b := tee(ctx, in)
	var wg sync.WaitGroup
	for id, output := range []<-chan int{a, b} {
		wg.Add(1)
		go func(name int, ch <-chan int) {
			defer wg.Done()
			for v := range ch { fmt.Printf("subscriber %d: %d\n", name, v) }
		}(id, output)
	}
	for i := 1; i <= 5; i++ { in <- i }
	close(in)
	wg.Wait()
}
