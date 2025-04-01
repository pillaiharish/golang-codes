package main

import (
	"fmt"
	"oop-polymorphism/pets"
	"time"
)

func main() {
	// slice of animals
	var animals []pets.Pet
	animals = append(animals, pets.NewDog("Tom", "Black & White", "American Bully"))
	animals = append(animals, pets.NewCat("Sphynx", "White", "Mixed"))

	for _, animal := range animals {
		fmt.Printf("At 10 am: %s \n", animal.Feed("Chicken"))
		fmt.Println(animal.GiveAttention("pat on back"))
		if animal.IsHungry() {
			fmt.Println(animal.Feed("Steak"))
		} else {
			fmt.Println("Waiting for animal to be hungry...")
			time.Sleep(3 * time.Second)
			fmt.Println("Feed Him.")
		}
	}

}

// Output:
// harish:oop-polymorphism harish$ go run main.go
// At 10 am: The animal is eating Chicken
// Tom is playing. pat on back means he is happy to see you
// Waiting for animal to be hungry...
// Feed Him.
// At 10 am: The animal is eating Chicken
// Sphynx is playing. pat on back means he is not happy to see you
// The animal is eating Steak
