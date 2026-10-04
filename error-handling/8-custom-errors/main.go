package main

import (
	"errors"
	"fmt"
)

// Define a custom error type (a struct) with extra fields
type APIError struct {
	Code    int
	Message string
}

// Implemented custom Error interface to return string error.
func (a *APIError) Error() string {
	return fmt.Sprintf("API Error Code: %d with Msg: %s", a.Code, a.Message)
}

// Define API Error call
func callAPIErr() error {
	custErr := &APIError{Code: 404, Message: "Page Not Found"}
	return custErr
}

func main() {
	err := callAPIErr()
	var valAPIErr *APIError
	if errors.As(err, &valAPIErr) {

		fmt.Printf("Error received: %v\n", valAPIErr.Error())
		fmt.Println("HTTP Code:", valAPIErr.Code)
		fmt.Println("Error Message:", valAPIErr.Message)
	}
}

// Output:
// Error received: API Error Code: 404 with Msg: Page Not Found
// HTTP Code: 404
// Error Message: Page Not Found
