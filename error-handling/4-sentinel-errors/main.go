package main

import (
	"errors"
	"fmt"
)

// A "sentinel error" is a pre-declared, global error variable
// used to signify a very specific, recognizable error condition.
// By convention, they start with Err.

var ErrNotAuthorized = errors.New("Not Authorized: ")
var ErrLoginRequired = errors.New("Please login to access")

func checkAccess(val bool) (int, error) {
	if val == false {
		return 0, ErrNotAuthorized
	}
	return 1, nil
}
func main() {
	_, err := checkAccess(false)
	fmt.Println(err, ErrLoginRequired)
}

// Output:
// Not Authorized:  Please login to access
