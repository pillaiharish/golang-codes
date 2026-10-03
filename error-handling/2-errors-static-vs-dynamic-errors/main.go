package main

import (
	"errors"
	"fmt"
)

// # Creating Errors
// Go provides two primary ways to create errors in the standard library:
// 1) errors.New("message"): Use this for simple, static error messages.
// 2) fmt.Errorf("message: %v", value): Use this when you need to include dynamic variables in your error message.

func main() {
	// Built-in error type from errors.New
	err1 := errors.New("something went wrong")
	fmt.Println("There was this error1: ", err1)

	// Print dynamic errors
	err2 := fmt.Errorf("failed to connect to port %d", 8080)
	fmt.Println("Second error: ", err2)

}

// Output:
// There was this error1:  something went wrong
// Second error:  failed to connect to port 8080
