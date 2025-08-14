package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <- chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range(jobs) {
		results <- n*n
	}
}

func main() {

	jobs    := make(chan int)      // fan-out
	results := make(chan int)      // fan-in
	var wg sync.WaitGroup
	numWorkers := 3
	
	// Start workers
	wg.Add(numWorkers)
	for i:=0;i<numWorkers;i++ {
		go worker(i, jobs, results, &wg)
	}

	// closing results when all workers are done
	go func(){
		wg.Wait()
		close(results)
	}()

	// feed jobs then close jobs
	// square of first 20 numbers
	go func() {
		for j:=1;j<=20;j++ {
			jobs <- j
		}
		close(jobs)
	}()
	for r := range results {
			fmt.Println("square: ", r)
	}
}


















/*
jobs ───▶  workers  ───▶ results
              │            ▲
              └─ wg.Done() │
*/
