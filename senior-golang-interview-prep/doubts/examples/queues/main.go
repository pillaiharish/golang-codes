package main

import (
	"context"
	"fmt"
)

// Callers must prevent sends racing with channel closure.
func trySend(q chan<- int, v int) bool {
	select {
	case q <- v:
		return true
	default:
		return false
	}
}

func send(ctx context.Context, q chan<- int, v int) error {
	select {
	case q <- v:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	q := make(chan int, 1)
	fmt.Println(trySend(q, 1), trySend(q, 2)) // true false
	fmt.Println(<-q)                          // 1
	unbuffered := make(chan int)
	fmt.Println(trySend(unbuffered, 3)) // false: no receiver
	q <- 4
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	fmt.Println(send(ctx, q, 5)) // context canceled: q remains full
	close(q)                     // no other sender exists
	fmt.Println(<-q)             // drain 4
	value, open := <-q
	fmt.Println(value, open) // 0 false
}
