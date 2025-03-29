package pets

import (
	"fmt"
	"strings"
)

type Dog struct {
	Name  string
	Color string
	Breed string
}

func (d Dog) Feed(food string) string {
	return fmt.Sprintf("%s is eating %s", d.Name, food)
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
