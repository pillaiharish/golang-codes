package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Invoice struct {
		Book      string
		Price     int
		Available bool
	}
	var purchase Invoice = Invoice{
		Book:      "Stay Hungry Stay Foolish",
		Price:     1000,
		Available: true,
	}
	json_purchase, err := json.Marshal(purchase)
	if err != nil {
		fmt.Println("Error in Marshal data", err)
		return
	}

	// Only println will return the  []byte value convert to string
	fmt.Println(string(json_purchase)) // like json.stringify()

	// Initializing raw Map
	data := map[string]interface{}{
		"name": "Harish",
		"age":  30,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	// Uncomment below to verify if json marshalling worked for jsonData
	// var map_read map[string]interface{}
	// json.Unmarshal(jsonData, &map_read)
	// fmt.Println(map_read["name"])

}

// Output:
// {"Book":"Stay Hungry Stay Foolish","Price":1000,"Available":true}
// {"age":30,"name":"Harish"}
// Harish
