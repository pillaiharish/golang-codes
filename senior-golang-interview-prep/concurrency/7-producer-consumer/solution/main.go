package main

import "fmt"

func main() {
	values := make(chan int)
	go func() {
		defer close(values) // only the producer closes the channel
		for i := 1; i <= 10; i++ {
			values <- i
		}
	}()

	for value := range values {
		fmt.Println(value)
	}
}
