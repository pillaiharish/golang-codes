package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("buffered before third send:", len(ch), "of", cap(ch))

	thirdSent := make(chan struct{})
	go func() {
		ch <- 3 // waits until main receives a value
		close(thirdSent)
	}()

	fmt.Println("first receive:", <-ch)
	<-thirdSent
	close(ch)
	for v := range ch {
		fmt.Println("remaining:", v)
	}
}
