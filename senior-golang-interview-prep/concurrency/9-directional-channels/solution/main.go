package main

import "fmt"

func produce(out chan<- int) {
	defer close(out)
	for i := 1; i <= 3; i++ {
		out <- i
	}
}

func consume(in <-chan int) {
	for v := range in {
		fmt.Println("received:", v)
	}
}

func main() {
	ch := make(chan int)
	go produce(ch)
	consume(ch)
}
