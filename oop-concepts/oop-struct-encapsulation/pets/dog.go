package pets

import (
	"fmt"
	"strings"
	"time"
)

type Dog struct {
	Name      string
	Color     string
	Breed     string
	lastSlept time.Time
}

func (d Dog) needsSleep() bool {
	// if the dog sleeps more than 4 hrs then it needs sleep
	return time.Now().Sub(d.lastSlept) > 4*time.Hour
}

func (d Dog) sleep() {
	d.lastSlept = time.Now() // dog is sleeping now
}

func (d Dog) Feed(food string) string {
	return fmt.Sprintf("%s is eating %s", d.Name, food)
}

func (d Dog) PlayWithDog(activity string) string {
	// check if dog needs sleep or not
	if d.needsSleep() { // if yes then let it sleep
		d.sleep()
		return "Your dog is tired and needs sleep"
	}
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

// since we cannot instantiate lastSlept while creating Dog object in main package
// We will do it here  
func NewDog(name, color, breed string, lastSlept time.Time) Dog {
	return Dog{
		Name:      name,
		Color:     color,
		Breed:     breed,
		lastSlept: lastSlept,
	}
}
