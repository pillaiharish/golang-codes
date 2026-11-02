package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("case", "drain", "drain, send-closed, close-twice, close-nil, send-nil, receive-nil")
	flag.Parse()

	switch *mode {
	case "drain":
		ch := make(chan int, 1)
		ch <- 7
		close(ch)
		v, ok := <-ch
		fmt.Println("buffered value:", v, ok)
		v, ok = <-ch
		fmt.Println("closed and drained:", v, ok)
	case "send-closed":
		ch := make(chan int)
		close(ch)
		ch <- 1 // intentionally panics
	case "close-twice":
		ch := make(chan int)
		close(ch)
		close(ch) // intentionally panics
	case "close-nil":
		var ch chan int
		close(ch) // intentionally panics
	case "send-nil":
		var ch chan int
		ch <- 1 // intentionally deadlocks
	case "receive-nil":
		var ch chan int
		<-ch // intentionally deadlocks
	default:
		fmt.Println("unknown case:", *mode)
	}
}
