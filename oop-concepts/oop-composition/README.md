# OOP Composition in Go

This folder demonstrates the concept of **composition** in Go using structs. Instead of relying on inheritance (as in traditional object-oriented programming), Go uses composition to build complex types by combining simpler ones. This example showcases how composition can be used to model animals and pets.


---

## Concepts

### What is Composition?

Composition is a design principle that allows you to combine multiple types to create more complex ones. In Go:
- **Struct Embedding**: A struct can embed another struct, inheriting its fields and methods.
- **Code Reuse**: Composition promotes code reuse and avoids duplication.
- **Flexibility**: It allows you to build modular and maintainable code.

---

## Code Explanation

### Animal Struct

The `Animal` struct represents a generic animal with basic properties and behaviors.

```
package pets

import (
"fmt"
"time"
)

type Animal struct {
sound string
lastAte time.Time
}

func (a Animal) Feed(food string) string {
a.lastAte = time.Now()
return fmt.Sprintf("The animal is eating %s", food)
}

func (a Animal) Sound(s string) string {
a.sound = s
return a.sound
}

func (a Animal) IsHungry() bool {
return time.Since(a.lastAte) > 2*time.Second
}
```


#### Key Features:
1. **Fields**:
   - `sound`: Represents the sound the animal makes.
   - `lastAte`: Tracks the last feeding time.
2. **Methods**:
   - `Feed`: Updates feeding time and returns a message.
   - `Sound`: Sets the sound of the animal.
   - `IsHungry`: Checks if the animal is hungry based on elapsed time since last feeding.

---

### Dog Struct

The `Dog` struct embeds the `Animal` struct, allowing it to inherit its properties and methods while adding dog-specific attributes and behaviors.

```
package pets

import (
    "fmt"
    "strings"
    "time"
)

type Dog struct {
    Name string
    Color string
    Breed string
    Animal
}

func (d Dog) PlayWithDog(activity string) string {
    response := ""
    switch strings.ToUpper(activity) {
    case ("wag"):
        response = "wags tail and is happy"
    case ("pat on back"):
        response = "wants more pat on back"
    default:
        response = "is happy to see you"
    }
    return fmt.Sprintf("%s is playing. %s means he %s", d.Name, activity, response)
}

func NewDog(name, color, breed string) Dog {
return Dog{
    Name: name,
    Color: color,
    Breed: breed,
    Animal: Animal{lastAte: time.Now()},
    }
}
```


#### Key Features:
1. **Fields**:
   - `Name`, `Color`, `Breed`: Attributes specific to a dog.
2. **Embedded Struct**:
   - The embedded `Animal` struct provides access to its fields and methods.
3. **Methods**:
   - `PlayWithDog`: Defines dog-specific behavior based on activities like "wag" or "pat on back".
4. **Constructor**:
   - `NewDog`: A helper function to initialize a new dog instance.

---

### Main Function

The main function demonstrates how to create a new dog and interact with it using composition.

```
package main

import (
"fmt"
"oop-composition/pets"
"time"
)

func main() {
    pet := pets.NewDog(
        "Tom",
        "White",
        "American Bully",
    )
    fmt.Printf("At 10 am: %s \n", pet.Feed("Chicken"))
    fmt.Println(pet.PlayWithDog("pat on back"))
    if pet.IsHungry() {
        fmt.Println(pet.Feed("Steak"))
    } else {
        fmt.Println("Waiting for animal to be hungry...")
        time.Sleep(3 * time.Second)
        fmt.Println("Feed Him.")
    }
}
```


## Conclusion

Composition in Go provides a flexible way to build complex types by combining simpler ones. By using struct embedding, we can create rich behaviors while maintaining clean and manageable code. This approach aligns with Go's philosophy of simplicity and clarity.

