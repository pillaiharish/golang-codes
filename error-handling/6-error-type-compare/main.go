package main

import (
	"errors"
	"fmt"
)

// Inspecting Wrapped Errors (errors.Is and errors.As)
// If you wrap an error, a simple == comparison will fail because the outer error is not the exact same object as the inner error. Go provides two powerful functions to inspect error chains:
// errors.Is → compare error identity/meaning
// errors.As → extract error TYPE

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

	// A. errors.Is(err, target)
	// Checks if the error, or any error in its wrap chain, matches a specific target (usually a sentinel error).
	fmt.Println("Check if ErrNotAuthorized data is present in returned error: ",
		errors.Is(finalError, ErrNotAuthorized))

	// B. errors.As(err, &target)
	// While errors.Is checks for equality, errors.As checks for type.
	// It is used with custom error types (structs that implement the error interface)
	// to extract the underlying error value so you can read its fields.
	// (This is the natural next step after learning errors.Is).
	// var validate
	// fmt.Println(errors.As(ErrNotAuthorized, finalError))

	// further examples of errors.As() is present in 7-error-as-example package.

}

// Output:
// Not authorized: Please log in as Admin
