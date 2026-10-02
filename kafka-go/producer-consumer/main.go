package main

import (
	"fmt"
	"sync"
	"time"
)

// Producer sends data to the channel
func producer(ch chan<- int, id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this producer is done when the function exits

	for i := 1; i <= 5; i++ {
		item := id*100 + i
		fmt.Printf("Producer %d: producing %d\n", id, item)

		// Send item to the channel (blocks if channel buffer is full)
		ch <- item

		// Simulate work time
		time.Sleep(100 * time.Millisecond)
	}
}

// Consumer receives data from the channel
func consumer(ch <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this consumer is done

	// The 'range' loop automatically exits when the channel is closed
	for item := range ch {
		fmt.Printf("Consumer %d: consumed %d\n", id, item)

		// Simulate processing time
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	// 1. Create a buffered channel with a capacity of 10
	// Buffering allows producers to keep working even if consumers are temporarily slow
	taskChannel := make(chan int, 10)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// 2. Start Producers
	numProducers := 2
	for i := 1; i <= numProducers; i++ {
		producerWg.Add(1)
		go producer(taskChannel, i, &producerWg)
	}

	// 3. Start Consumers
	numConsumers := 3
	for i := 1; i <= numConsumers; i++ {
		consumerWg.Add(1)
		go consumer(taskChannel, i, &consumerWg)
	}

	// 4. Wait for all producers to finish, then CLOSE the channel
	// We do this in a separate goroutine so we don't block the main thread
	go func() {
		producerWg.Wait()
		fmt.Println("\n--- All producers finished. Closing channel. ---")
		close(taskChannel)
	}()

	// 5. Wait for all consumers to finish processing
	consumerWg.Wait()

	fmt.Println("All tasks completed. Exiting program.")
}

