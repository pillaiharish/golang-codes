package main

import (
	"context"
	"fmt"
)

func bridge(ctx context.Context, streams <-chan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			var stream <-chan int
			select {
			case <-ctx.Done(): return
			case next, ok := <-streams:
				if !ok { return }
				stream = next
			}
			for stream != nil {
				select {
				case <-ctx.Done(): return
				case v, ok := <-stream:
					if !ok { stream = nil; continue }
					select {
					case out <- v:
					case <-ctx.Done(): return
					}
				}
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	streams := make(chan (<-chan int))
	go func() {
		defer close(streams)
		for _, group := range [][]int{{1, 2, 3}, {4, 5}, {6}} {
			inner := make(chan int)
			streams <- inner
			for _, v := range group { inner <- v }
			close(inner)
		}
	}()
	for v := range bridge(ctx, streams) {
		fmt.Println(v)
	}
}
