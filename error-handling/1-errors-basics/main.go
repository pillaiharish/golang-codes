package main

import (
	"errors"
	"fmt"
)

// 1) the error Type (The Foundation)
// In Go, error is a built-in interface type defined in the standard library like this:
// type error interface {
// 	Error() string
// }

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero exception")
	}
	return a / b, nil
}
func main() {
	// Example of divide by zero exception handled
	x, y := 10, 0
	answer, err := divide(x, y)
	fmt.Println("Answer is: ", answer, " Error any: ", err)
	answer, err = divide(4, 2)
	fmt.Println("Answer is: ", answer, " Error any: ", err)

}

// Output:
// Answer is:  0  Error any:  Divide by zero exception
// Answer is:  2  Error any:  <nil>
