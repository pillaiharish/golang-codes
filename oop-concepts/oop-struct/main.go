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
