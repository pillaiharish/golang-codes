package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var msg map[string]interface{}
	j_data := `{"Name":"Harish", "age":30}`
	err := json.Unmarshal([]byte(j_data), &msg)
	if err != nil {
		fmt.Println("Erro in unmarshalling data", err)
		return
	}
	fmt.Println(msg)
	fmt.Println("Username is: ", msg["Name"])
	fmt.Println("User age is: ", msg["age"])
}

// Output:
// map[Name:Harish age:30]
// Username is:  Harish
// User age is:  30
