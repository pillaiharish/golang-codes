package main

import "fmt"

type Speaker interface {
	Speaks() string
}

type myDog struct {
	breed string
}

type myCat struct {
	color string
}

func (d *myDog) Speaks() string {
	return "Woof Woof!"
}

func (myCat) Speaks() string {
	return "Meow Meow!"
}

func main() {
	var dogSpeak Speaker
	d1 := &myDog{breed: "Bull dog"}
	dogSpeak = d1
	fmt.Println("My", d1.breed, " goes:", dogSpeak.Speaks())
	// In below d2 = dog will not work as dog belongs to the method of set *Dog:

	// var d2 Speaker
	// dog := myDog{breed: "Pomerian"}
	// d2 = dog    // COMPILATION ERROR! myDog does not implement Speaker (Speak method has pointer receiver)

	var cat1 Speaker
	cat1 = myCat{color: "Orange"}
	fmt.Println("My cat says:", cat1.Speaks())

	var cat2 Speaker
	cat2 = &myCat{color: "Night"}
	fmt.Println("My black cat says", cat2.Speaks())
}

// Summary: Which one should you choose?
// A) Use Pointer Receivers if:
//    - The method needs to modify the fields inside the struct.
//    - The struct is large and you want to avoid copying it on every method call.
//    - Consequence: Any interface requiring this method can only accept pointers to your struct (*T).
// B) Use Value Receivers if:
//    - The struct is small (like basic coordinates or configuration flags).
//    - You want maximum flexibility so both values and pointers can satisfy interfaces.

// Output
