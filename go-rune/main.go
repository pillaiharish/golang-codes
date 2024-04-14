/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*
* A rune in Go represents a Unicode code point and is an alias for int32.
* This type of slice is often used for processing Unicode characters of strings.
*
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package main

import "fmt"

func main() {
	// initialize empty rune slice
	stack := []rune{}
	s := "hello"
	for _, i := range s {
		stack = append(stack, i)
	}

	// pop rune from the stack
	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		top := stack[lastIdx]
		stack = stack[:lastIdx]
		fmt.Printf("%c \n", top) // This will print the string in reverse
	}

	// declaring and initializing a Unicode character
	emoji := '😀'

	// %T represents the type of the variable num
	fmt.Printf("Data type of %c is %T and the rune value is %U \n", emoji, emoji, emoji)

	var str string = "ABCD"
	r_array := []rune(str)
	fmt.Printf("Array of rune values for A, B, C and D: %U\n", r_array)

}
