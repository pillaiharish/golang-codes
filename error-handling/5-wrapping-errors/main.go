package main

import (
	"errors"
	"fmt"
)

// Sometimes a low-level function fails,
// and as the error bubbles up to higher-level functions,
// you want to add more context without losing the original error.
// Since Go 1.13, you can wrap errors using the %w verb in fmt.Errorf:

var ErrNotAuthorized = errors.New("Not authorized: ")

func checkLogin(val bool) error {
	if val == false {
		return errors.New("Please log in as")
	}
	return nil
}
func main() {
	err := checkLogin(false)
	usrName := "Admin"
	finalError := fmt.Errorf("%w %w %v", ErrNotAuthorized, err, usrName)
	// This creates a logical chain for sentinel error "Not authorized: " wraps
	// "no more tea available" and also appends username "Admin".
	fmt.Println(finalError)

}

// Output:
// Not authorized: Please log in as Admin
