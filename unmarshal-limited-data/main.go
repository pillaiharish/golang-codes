package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Incoming struct {
		Name     string
		Age      int
		IsOnline bool
	}
	var msg Incoming
	// var msg map[string]interface{}
	j_data := `{"Name":"Harish", "age":30,"location": "India","IsOnline":true, "ManyMoreData":"NotConsidered"}`
	err := json.Unmarshal([]byte(j_data), &msg)
	if err != nil {
		fmt.Println("Erro in unmarshalling data", err)
		return
	}
	fmt.Println(msg)
	fmt.Println("Username is: ", msg.Name)
	fmt.Println("User age is: ", msg.Age)
	fmt.Println("Status: ", msg.IsOnline)
}

// Output:
// {Harish 30 true}
// Username is:  Harish
// User age is:  30
// Status:  true
