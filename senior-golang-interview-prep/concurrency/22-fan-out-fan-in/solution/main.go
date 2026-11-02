package main

import (
	"fmt"
	"sync"
)

func merge(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(inputs))
	for _, input := range inputs {
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch { out <- v }
		}(input)
	}
	go func() { wg.Wait(); close(out) }()
	return out
}

func main() {
	jobs := make(chan int)
	outputs := make([]<-chan int, 0, 3)
	for i := 0; i < 3; i++ {
		out := make(chan int)
		outputs = append(outputs, out)
		go func(workerID int, output chan<- int) {
			defer close(output)
			for job := range jobs { output <- job * 2 }
		}(i, out)
	}
	go func() {
		for i := 1; i <= 10; i++ { jobs <- i }
		close(jobs)
	}()

	count := 0
	for v := range merge(outputs...) {
		fmt.Println(v)
		count++
	}
	fmt.Println("merged results:", count)
}
