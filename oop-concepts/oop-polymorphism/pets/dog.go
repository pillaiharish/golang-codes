package pets

import (
	"fmt"
	"strings"
	"time"
)

type Dog struct {
	Name  string
	Color string
	Breed string
	Animal
}

func (d Dog) GiveAttention(activity string) string {
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

func NewDog(name, color, breed string) Pet {
	return Dog{
		Name:   name,
		Color:  color,
		Breed:  breed,
		Animal: Animal{lastAte: time.Now()},
	}
}
