package main

import (
	"fmt"
	"sync"
)

func main() {
	var value int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		value = 42
	}()
	wg.Wait() // Done synchronizes before Wait returns
	fmt.Println("after WaitGroup join:", value)

	var message string
	done := make(chan struct{})
	go func() {
		message = "visible after receive"
		close(done)
	}()
	<-done // close synchronizes before a receive observing closure
	fmt.Println("after channel signal:", message)
}
