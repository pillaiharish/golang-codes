# Go Polymorphism using Shapes like Circle and Rectangle 

This repository demonstrates **polymorphism** in Go using **interfaces** and **structs**.

## What is Polymorphism?

Polymorphism allows different data types to be treated as the same interface type. In Go, this is achieved through **interfaces**.

When multiple types implement the same interface, you can write functions that operate on the interface instead of a specific type.

## Example

We define a `Shape` interface with `Area()` and `Perimeter()` methods. Then we create two types:
- `Circle`
- `Rectangle`

Both implement the `Shape` interface.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

## Implementation
- Circle struct with methods `Area()` and `Perimeter()`

- Rectangle struct with methods `Area()` and `Perimeter()`

We create instances of these structs using pseudo-constructors `NewCircle()` and `NewRectangle()` and store them in a slice of type `Shape`.
```
shapes := []Shape{
    NewCircle(5),
    NewRectangle(5, 6),
}
```

Then we calculate and print their areas and perimeters using the interface methods.

## How It Demonstrates Polymorphism
We define one interface (Shape)

Implement it with multiple structs (Circle and Rectangle)

Write code that treats all Shape types uniformly