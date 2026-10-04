
package main

import "fmt"

// Two Crucial Rules of defer:
// 1) LIFO Order: If you have multiple defer statements, they execute in Last-In, First-Out order (like a stack).
// 2) Arguments are evaluated immediately: The values passed to a deferred function are locked in at the moment 
// defer is called, not when the function actually executes.

// defer keyword works in LIFO order like a stack when multiple defers are present
func main() {
	i := 1
	defer fmt.Println("1) first defer statement value of i:", i)
	i= i+1
	defer fmt.Println("2) second defer statement, value of i:", i)
	defer fmt.Println("3) third defer statement, value of i:", i)
	i=3
	
	fmt.Println("4) Print statement after all defer without any defer, value of i:", i)
}

// Output:
// 4) Print statement after all defer without any defer, value of i: 3
// 3) third defer statement, value of i: 2
// 2) second defer statement, value of i: 2
// 1) first defer statement value of i: 1
