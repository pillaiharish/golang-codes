package main

import "fmt"

func main() {
	ch := make(chan int) // zero capacity: sender and receiver rendezvous
	go func() { ch <- 42 }()
	fmt.Println("received:", <-ch)
}
