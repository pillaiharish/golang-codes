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

// Output:
// harish:oop-composition harish$ go run main.go
// At 10 am: The animal is eating Chicken
// Tom is playing. pat on back means he is happy to see you
// Waiting for animal to be hungry...
// Feed Him.
