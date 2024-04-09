package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int = 10000
var purchases [5]int = [5]int{3000, 273, 590, 499, 2030}
var wg sync.WaitGroup

func update_balance(c chan int) {
	defer wg.Done()

	purchase := <-c                                                           // receive value from channel
	balance = balance - purchase                                              // perform computation
	fmt.Printf("The balance after purchase of %v is %v\n", purchase, balance) // final balance

}

func main() {
	c := make(chan int) // channel creation
	wg.Add(5)
	for i := 0; i < 5; i++ {

		go update_balance(c)
		c <- purchases[i] // send value to channel
	}
	close(c) // close connection
	wg.Wait()
	fmt.Println("Final balance is ", balance)

	time.Sleep(1 * time.Second) // need to wait for goroutine to finish task
}
