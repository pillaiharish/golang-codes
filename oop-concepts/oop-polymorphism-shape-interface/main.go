package main

import (
	"fmt"
	"math"
)

// Shape interface defines polymorphic behavior
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct represents a circle with a radius
type Circle struct {
	Radius float64
}

// Rectangle struct represents a rectangle with width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle "constructor"
func NewCircle(radius float64) Shape {
	// Negative values should not be accepted
	if radius < float64(0.1) {
		return nil
	}
	return &Circle{Radius: radius}
}

// Rectangle "constructor"
func NewRectangle(width, height float64) Shape {
	// Negative values should not be accepted
	if width < float64(0.1) || height < float64(0.1) {
		return nil
	}
	return &Rectangle{Width: width, Height: height}
}

// Circle's Area method
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle's Perimeter method
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle's Area method
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle's Perimeter method
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	shapes := []Shape{
		NewCircle(5),
		NewRectangle(5, 6),
	}

	result := make(map[string]float64)
	for _, shape := range shapes {
		switch s := shape.(type) {
		case *Circle:
			result["circleArea"] = s.Area()
			result["circlePerimeter"] = s.Perimeter()
		case *Rectangle:
			result["rectangleArea"] = s.Area()
			result["rectanglePerimeter"] = s.Perimeter()
		}
	}

	fmt.Println(result)
}

// Output:
// go run main.go
// map[circleArea:78.53981633974483 circlePerimeter:31.41592653589793 rectangleArea:30 rectanglePerimeter:22]
