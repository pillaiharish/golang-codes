# Structs as Classes and Struct Embedding in Go

This explains the concepts of using structs as classes in Go, along with the idea of struct embedding. The provided code demonstrates how to define a `Rectangle` struct, implement methods for it, and utilize it in a simple program.

---

## Introduction

In Go, structs are used to create complex data types that group together related data. While Go does not have traditional classes like other object-oriented programming languages, structs can be used to achieve similar functionality. Methods can be defined on structs, allowing them to exhibit behavior akin to classes.

## Structs as Classes

A struct in Go is a composite data type that groups together variables (fields) under a single name. The following example defines a `Rectangle` struct with two fields: `Length` and `Width`.

```
type Rectangle struct {
    Length int
    Width int
}
```


By defining methods on the `Rectangle` struct, we can encapsulate behavior related to rectangles, similar to how methods are defined within classes in traditional OOP languages.

## Methods on Structs

Methods can be associated with structs in Go. The following methods are defined for the `Rectangle` struct:

### GetArea Method

Calculates the area of the rectangle:

```
func (r Rectangle) GetArea() int {
    return r.Length * r.Width
}
```


### GetPerimeter Method

Calculates the perimeter of the rectangle:

```
func (r Rectangle) GetPerimeter() int {
    return 2 * (r.Length + r.Width)
}
```


### FindLargest Method

Compares area and perimeter and returns which is larger:

```
func (r Rectangle) FindLargest(area, perimeter int) string {
    if area < perimeter {
        return "perimeter"
    } else {
        return "area"
    }
}
```


## Struct Embedding

Struct embedding allows you to include one struct within another. This is useful for creating more complex types that share common fields or behaviors without using inheritance. For example:

```
type Shape struct {
    Color string
}

type ColoredRectangle struct {
    Rectangle // Embedding Rectangle
    Shape // Embedding Shape
}
```

In this example, `ColoredRectangle` inherits fields and methods from both `Rectangle` and `Shape`, allowing for more complex behaviors while maintaining simplicity.


## Conclusion

We explored how structs in Go can be used similarly to classes in traditional OOP languages. We defined a `Rectangle` struct with methods for calculating area and perimeter, and compared these values. Additionally, we touched upon the concept of struct embedding for creating more complex types.