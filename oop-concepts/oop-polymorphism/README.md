# Pet Polymorphism Example

This project demonstrates the concept of polymorphism in Go through a pet management system. The code showcases how different animal types (Dog and Cat) can share common behaviors while implementing their unique characteristics.

## Overview
This application models a pet management system where different types of animals (Dog and Cat) implement a common `Pet` interface. This allows for treating different types of pets uniformly while maintaining their specific behaviors.

## Polymorphism Concepts
Polymorphism is a core concept in object-oriented programming that allows objects of different types to be treated as objects of a common type. This project demonstrates several key aspects of polymorphism:

### 1. Interface-based Polymorphism
The `Pet` interface defines a contract that all pet types must fulfill:

```
type Pet interface {
    Feed(food string) string
    GiveAttention(activity string) string
    IsHungry() bool
}
```

Both `Dog` and `Cat` types implement this interface, allowing them to be used interchangeably where a `Pet` is expected.

### 2. Struct Embedding (Composition)
Rather than using inheritance (which Go doesn't have), this project uses composition through struct embedding:

```
type Dog struct {
Name string
    Color string
    Breed string
    Animal // Embedded struct
}
```

This allows `Dog` and `Cat` to inherit behavior from the `Animal` struct while adding their own fields and methods.

### 3. Method Overriding
While `Dog` and `Cat` inherit the `Feed` and `IsHungry` methods from `Animal`, they provide their own implementation of `GiveAttention`, allowing for type-specific behavior:

```
func (d Dog) GiveAttention(activity string) string {
    // Dog-specific implementation
}

func (c Cat) GiveAttention(activity string) string {
    // Cat-specific implementation
}
```


### 4. Runtime Polymorphism
In `main.go`, we store different pet types in a single slice of the interface type (`[]pets.Pet`). When methods are called on these objects, the appropriate implementation is selected at runtime based on the actual type:

```
var animals []pets.Pet
animals = append(animals, pets.NewDog("Tom", "Black & White", "American Bully"))
animals = append(animals, pets.NewCat("Sphynx", "White", "Mixed"))

for _, animal := range animals {
    // The correct implementation is called based on the actual type
    animal.GiveAttention("pat on back")
}
```

## Project Structure
The project is organized into the following files:
- `pets/animal.go`: Defines the base `Animal` struct with common behaviors.
- `pets/pet.go`: Defines the `Pet` interface that all pet types must implement.
- `pets/dog.go`: Implements the `Dog` struct with dog-specific behaviors.
- `pets/cat.go`: Implements the `Cat` struct with cat-specific behaviors.
- `main.go`: Contains the main application that demonstrates polymorphic behavior.

## How It Works
1. The base `Animal` struct provides common behaviors like `Feed` and `IsHungry`.
2. Specific pet types (`Dog` and `Cat`) embed the `Animal` struct to inherit its behaviors.
3. Each pet type implements its unique version of `GiveAttention`.
4. Factory functions (`NewDog` and `NewCat`) return instances as the `Pet` interface type.
5. The main program creates different pets and treats them uniformly through the `Pet` interface.


