package main

import (
	"fmt"
	"time"
)

// Deliberately unreliable approach: sleeping is not a completion barrier.
func main() {
	go func() {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Hello from goroutine")
	}()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Main finished (may have abandoned worker)")
}
