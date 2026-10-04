package main

import "fmt"

func divide(a,b int) (result int) {
	defer func(){
		if r:=recover(); r!=nil {
			fmt.Println("Recover from panic due to value:", r)
			result = 0
		}
	}()

	if b==0{
		panic("Divide by zero panic")
	}
	return a/b
}

func main() {
	fmt.Println("--- Test 1: Normal Execution ---")
	answer := divide(10,2)
	fmt.Println("Divide 10 by 2", answer)

	fmt.Println("\n--- Test 2: Panic and Recover ---")
	answer = divide(10,0)
	fmt.Println("Divide 10 by 0", answer)
}

// Output:
// --- Test 1: Normal Execution ---
// Divide 10 by 2 5
//
// --- Test 2: Panic and Recover ---
// Recover from panic due to value: Divide by zero panic
// Divide 10 by 0 0
