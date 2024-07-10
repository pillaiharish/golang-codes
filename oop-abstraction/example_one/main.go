/*
Abstraction involves hiding complex implementation details and showing only the necessary features of an object.

	In Go, this can be achieved using interfaces.
*/
package main

import "fmt"

type Shape interface {
	Area() int64
}

type Rectangle struct {
	length, width int64
}

func (r Rectangle) Area() int64 {
	return r.length * r.width
}

func main() {
	rectangle := Rectangle{length: 10, width: 25}
	fmt.Println("Area of rectangle is ", rectangle.Area())

}

/* 
Output:
Area of rectangle is  250
*/