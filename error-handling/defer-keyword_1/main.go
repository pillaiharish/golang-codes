package main

import "fmt"

// 1. defer: The Cleanup Crew
// The defer keyword schedules a function call to be executed right before 
// the surrounding function returns, no matter how it returns (whether it 
// finishes normally, returns early, or panics).
// Common uses: Closing files, unlocking mutexes, closing database connections, 
// or printing final logs.

func main() {

	fmt.Println("Print first line")
	
	// since we use defer it will be printed last just before the function exit.
	defer fmt.Println("Print last line")

	fmt.Println("Print second line")
}

// Output:
// Print first line
// Print second line
// Print last line
