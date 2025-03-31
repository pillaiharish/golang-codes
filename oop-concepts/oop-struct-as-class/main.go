package main

import (
	"fmt"
)

// Define the struct
type Rectangle struct {
	Length int
	Width  int
}

// Method to calculate area
func (r Rectangle) GetArea() int {
	return r.Length * r.Width
}

// Method to calculate perimeter
func (r Rectangle) GetPerimeter() int {
	return 2 * (r.Length + r.Width)
}

func (r Rectangle) FindLargest(area, perimeter int) string {
	if area < perimeter {
		return "perimeter"
	} else {
		return "area"
	}
}

func main() {
	// Example input
	rect := Rectangle{Length: 2, Width: 4}

	// Output results
	fmt.Printf("Area: %d, Perimeter: %d\n", rect.GetArea(), rect.GetPerimeter())

	// Find largest of two
	fmt.Printf("Largest of two is: %s \n", rect.FindLargest(rect.GetArea(), rect.GetPerimeter()))
}

// Output:
// $ go run main.go
// Area: 8, Perimeter: 12
// Largest of two is: perimeter
