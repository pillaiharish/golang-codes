package main

import "fmt"

func main() {
	var x interface{} = "Hello" // This can also be replaced by "any"
	getX := x.(string)
	fmt.Println("Now we tell to get whatever is inside `x` and treat it as string")
	fmt.Println(getX)

	// The Dangerous Assertion (Panic!)
	getXint := x.(int)
	fmt.Println("Now we tell to get whatever is inside `x` and treat it as integer")
	fmt.Println(getXint)

}

// Output:
// Hello
// panic: interface conversion: interface {} is string, not int

// goroutine 1 [running]:
// main.main()
//         ~/golang-codes/go-syntax-traps/6-type-assertions-interface-1/main.go:12 +0xb4
// exit status 2
