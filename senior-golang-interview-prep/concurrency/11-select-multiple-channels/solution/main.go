package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	first := make(chan string, 1)
	second := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		first <- "A"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		second <- "B"
	}()

	select {
	case v := <-first:
		fmt.Println("first selected:", v)
	case v := <-second:
		fmt.Println("first selected:", v)
	}
	wg.Wait() // the unselected sender can still finish because channels are buffered
}
