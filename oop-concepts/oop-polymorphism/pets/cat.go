package pets

import (
	"fmt"
	"strings"
	"time"
)

type Cat struct {
	Name  string
	Color string
	Breed string
	Animal
}

func (c Cat) GiveAttention(activity string) string {
	response := ""
	switch strings.ToUpper(activity) {
	case ("cuddle"):
		response = "aggressive if someone cuddles it."
	case ("pat on back"):
		response = "scratches you back."
	default:
		response = "is not happy to see you"
	}
	return fmt.Sprintf("%s is playing. %s means he %s", c.Name, activity, response)
}

func NewCat(name, color, breed string) Pet {
	return Cat{
		Name:   name,
		Color:  color,
		Breed:  breed,
		Animal: Animal{lastAte: time.Now()},
	}
}
