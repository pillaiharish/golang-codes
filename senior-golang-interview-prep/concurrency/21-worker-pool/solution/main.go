package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)
	results := make(chan int)
	var workers sync.WaitGroup
	for id := 0; id < 3; id++ {
		workers.Add(1)
		go func(workerID int) {
			defer workers.Done()
			for job := range jobs {
				results <- job * 2
			}
		}(id)
	}
	go func() {
		for i := 1; i <= 10; i++ { jobs <- i }
		close(jobs)
	}()
	go func() {
		workers.Wait()
		close(results) // close only after all worker sends finish
	}()

	count := 0
	for result := range results {
		fmt.Println("result:", result)
		count++
	}
	fmt.Println("results collected:", count)
}
