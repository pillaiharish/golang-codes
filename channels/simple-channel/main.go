package main

import (
	"fmt"
	"time"
)

var balance int = 500

func update_balance(c chan int) {
	purchase := <-c                                       // receive value from channel
	balance = balance - purchase                          // perform computation
	fmt.Println("The balance after purchase is", balance) // final balance

}

func main() {
	c := make(chan int) // channel creation

	go update_balance(c)
	c <- 100 // send value to channel
	close(c) // close connection

	time.Sleep(1 * time.Second) // need to wait for goroutine to finish task
}
