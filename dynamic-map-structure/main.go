package main

import "fmt"

func main() {
	userData := make(map[string]interface{})
	userData["name"] = "Harish"
	userData["age"] = 29
	userData["preferences"] = []string{"coffee", "tea"}

	fmt.Println(userData)
}

// Output:
// map[age:29 name:Harish preferences:[coffee tea]]
