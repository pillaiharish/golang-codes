package main

import (
	"fmt"
)

// type People struct {
// 	Name string
// 	City string
// }

func main() {

	// people := []People{
	// 	{Name: "Harish", City: "Bengaluru"},
	// 	{Name: "Kumar", City: "Hyderabad"},
	// 	{Name: "Abhi", City: "Ahmedabad"},
	// }
	people := []struct {
		Name string
		City string
	}{
		{"Harish", "Bengaluru"},
		{"Kumar", "Hyderabad"},
		{"Abhi", "Bengaluru"},
	}
	// City is key and value is tuple of name and city
	cityMap := make(map[string][]string)
	// Group the names  of people residng in same city
	for _, person := range people {
		cityMap[person.City] = append(cityMap[person.City], person.Name)
	}
	// group by city name
	fmt.Println(cityMap)
}

// Output:
// map[Bengaluru:[Harish Abhi] Hyderabad:[Kumar]]
