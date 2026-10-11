package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	breed string
}

func (Dog) Speak() string {
	return "bark bark"
}

func main() {
	var s Speaker = Dog{breed: "Golden Retriever"}

	// We have 's', which is a Speaker. We can't do s.breed here.

	// Type Assertion: Try to extract the original Dog struct back out of 's'
	d, ok := s.(Dog)
	if ok {
		// Success! Now we have the concrete Dog struct again.
		fmt.Println("Retrieved breed:", d.breed) // Output: Golden Retriever
	}
}
