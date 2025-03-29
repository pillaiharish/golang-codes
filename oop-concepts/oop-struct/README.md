# Pets Package

This repository contains a Go package `pets` that demonstrates the use of structs and methods, showcasing Object-Oriented Programming (OOP) principles in Go. The package models a `Dog` struct with attributes and behaviors, such as feeding and playing.

---

## Introduction

Go does not have traditional classes like other OOP languages, but it achieves similar functionality using structs and methods. This `pets` package demonstrates these concepts by modeling a `Dog` entity with attributes (fields) and behaviors (methods).

---

## Features

- Define a `Dog` struct with attributes: `Name`, `Color`, and `Breed`.
- Implement methods for:
  - Feeding the dog (`Feed`).
  - Playing with the dog (`PlayWithDog`).
- Showcase encapsulation and behavior modeling in Go.

---

## Usage

### Dog Struct

The `Dog` struct represents a dog with three attributes:
```
type Dog struct {
Name string // The name of the dog
Color string // The color of the dog's fur
Breed string // The breed of the dog
}
```
#### Fields

- **Name**: A string representing the dog's name.
- **Color**: A string representing the color of the dog's fur.
- **Breed**: A string representing the breed of the dog.

### Feed Method

The `Feed` method simulates feeding the dog:
```
func (d Dog) Feed(food string) string {
return fmt.Sprintf("%s is eating %s", d.Name, food)
}
```
#### Parameters

- **food**: A string representing the type of food being given to the dog.

#### Returns

A formatted string indicating that the dog is eating the specified food.

### PlayWithDog Method

The `PlayWithDog` method simulates playing with the dog:
```
func (d Dog) PlayWithDog(activity string) string {
response := ""
switch strings.ToUpper(activity) {
case ("WAG"):
response = "wags tail and is happy"
case ("PAT ON BACK"):
response = "wants more pat on back"
default:
response = "is happy to see you"
}
return fmt.Sprintf("%s is playing. %s means he %s", d.Name, activity, response)
}
```
#### Parameters

- **activity**: A string representing the activity you want to do with the dog (e.g., "wag" or "pat on back").

#### Returns

A formatted string indicating how the dog is responding to the activity.

---

## Example Output

Here’s how you can use the `pets` package in your main application:
```
package main

import (
"fmt"
"oop-struct/pets"
)

func main() {
pet := pets.Dog{Name: "Tom", Color: "White", Breed: "American Bully"}
fmt.Printf("At 10 am: %s \n", pet.Feed("Chicken"))
fmt.Println(pet.PlayWithDog("pat on back"))
}
```

### Output:
```
Tom is playing. pat on back means he wants more pat on back
```
---

## Contributing

Contributions are welcome! If you'd like to improve this package or add new features, feel free to fork this repository and submit a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.