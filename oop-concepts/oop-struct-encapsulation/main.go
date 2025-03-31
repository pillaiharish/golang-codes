package main

import (
	"fmt"
	"oop-struct-encapsulation/pets"
	"time"
)

func main() {

	// instead of using pets.Dog like we create in normal struct we will use the NewDog method
	// pet := pets.Dog{

	//Testing with sleepTime
	// sleepTime := time.Now()
	sleepTime := time.Now().Add(time.Duration(-5) * time.Hour)

	pet := pets.NewDog(
		"Bubzee",
		"Black",
		"Rottweiler",
		sleepTime,
		// below throws cannot refer to unexported field lastSlept in struct literal of type
		// lastSlept: time.Now(),	// error is expected as Encapsulation is applied at struct level
	)
	// Same can be verified at method level
	// pet.needsSleep undefined (cannot refer to unexported method needsSleep)
	// pet.needsSleep()		// error is expected as Encapsulation is applied at

	fmt.Printf("At 10 am: %s \n", pet.Feed("Dog food"))
	fmt.Println(pet.PlayWithDog("pat on back"))
}

//  Output:
// TestCase 1: when lastseelp is just now or dog woke just now
// harish:oop-struct-encapsulation harish$ go run main.go
// At 10 am: Bubzee is eating Dog food
// Bubzee is playing. pat on back means he is happy to see you
//
// TestCase 2: updating last sleep to 5 hours
// harish:oop-struct-encapsulation harish$ go run main.go
// At 10 am: Bubzee is eating Dog food
// Your dog is tired and needs sleep
