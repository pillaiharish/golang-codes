package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string {
	return "something failed"
}

func getError() error {
	var e *MyError = nil
	return e
}

func main() {
	err := getError()

	fmt.Println(err == nil) // It prints: false

	// Conceptually an interface contains:
	// (dynamic type, dynamic value)

	// A truly nil interface is:
	// (nil, nil)

	// But err contains:
	// (*MyError, nil)
	// Therefore the interface itself isn't nil.

	// 	err
	// +------------------------+
	// | dynamic type: *MyError |
	// | dynamic value: nil     |
	// +------------------------+

	// nil interface
	// +--------------------+
	// | type: nil          |
	// | value: nil         |
	// +--------------------+

	// So below creates a bug:
	// err != nil

	// Correct pattern:
	// 	func getError() error {
	//     // ...
	//     return nil
	// }
}

// Output:
