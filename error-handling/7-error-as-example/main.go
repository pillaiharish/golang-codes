package main

import (
	"errors"
	"fmt"
)

// 1. Define a custom error type (a struct) with extra fields
type insufficientFundsError struct {
	balanceAmt int
	requestAmt int
}

// 2. Implement the built-in `error` interface by adding an Error() string method
// Note: We use a pointer receiver (*insufficientFundsError) so it can be modified if needed,
// and it's the standard practice for custom errors.
func (e *insufficientFundsError) Error() string {
	return fmt.Sprintf("Available balance: $%d, Requested amt: $%d", e.balanceAmt, e.requestAmt)
}

// 3. A function that returns a WRAPPED custom error
func withdraw(balance, request int) (int, error) {
	if balance < request {
		// Create the custom error
		customErr := &insufficientFundsError{balanceAmt: balance, requestAmt: request}

		// Wrap it with higher-level context using %w
		return 0, fmt.Errorf("Transaction declined %w", customErr)
	}
	return balance - request, nil
}

func main() {

	response, err := withdraw(50, 100)
	if err != nil {
		// 4. Declare a variable of the custom error type to hold the extracted error.
		// IT MUST BE A POINTER (&).
		var fundsErr *insufficientFundsError
		// 5. Use errors.As to search the error chain for this type.
		// We pass the error `err` and a pointer to our target variable `&fundsErr`.
		if errors.As(err, &fundsErr) {
			// Success! The error chain contains an insufficientFundsError.
			// We can now safely access its specific fields.
			shortFall := fundsErr.requestAmt - fundsErr.balanceAmt

			fmt.Println("Got custom error")
			fmt.Printf("You are short by $%d\n", shortFall)
			fmt.Println("Your full error: ", err)
		} else {
			// It was an error, but NOT an insufficientFundsError
			fmt.Printf("Some other error: %v, check again\n", err)
		}
	} else {
		fmt.Printf("Withdrawal of Amt: $%d successfull!\n", response)
	}

}

// Output:
// Got custom error
// You are short by $50
// Your full error:  Transaction declined Available balance: $50, Requested amt: $100
