package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

// Printing correct data type of error
// 1) Using fmt.Printf("%T", err) (Recommended)
// The %T verb prints the dynamic type of any interface value.
// 2) Using reflect lib

func main() {

	err1 := errors.New("Page not found")
	fmt.Printf("Type of err1: %T \n", err1)

	_, err2 := os.Open("no_file_exists.txt")
	fmt.Printf("Type of err2: %T\n", err2)

	fmt.Println("Using reflect err1: ", reflect.TypeOf(err1))
	fmt.Println("Using reflect err2: ", reflect.TypeOf(err2))
}

// Output:
// Type of err1: *errors.errorString
// Type of err2: *fs.PathError
// Using reflect err1:  *errors.errorString
// Using reflect err2:  *fs.PathError
