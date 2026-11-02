package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("case", "fixed", "fixed, send-no-receiver, receive-no-sender, range-no-close, early-exit")
	flag.Parse()

	switch *mode {
	case "fixed":
		ch := make(chan int)
		go func() {
			ch <- 42
			close(ch)
		}()
		for v := range ch {
			fmt.Println("received:", v)
		}
	case "send-no-receiver":
		ch := make(chan int)
		ch <- 1 // intentional deadlock
	case "receive-no-sender":
		ch := make(chan int)
		fmt.Println(<-ch) // intentional deadlock
	case "range-no-close":
		ch := make(chan int)
		go func() { ch <- 1 }() // producer does not close the channel
		for v := range ch {
			fmt.Println(v) // after 1, main waits forever
		}
	case "early-exit":
		go fmt.Println("worker may never print")
	default:
		fmt.Println("unknown case:", *mode)
	}
}
