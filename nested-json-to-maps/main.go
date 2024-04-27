package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{"user": {"name": "Harish", "details": {"age": 30, "location": "India"}}}`
	// var result map[string]interface{} // Access result["user"] // Output: map[details:map[age:30 location:India] name:Harish]
	var result map[string]map[string]interface{} // Access result["user"]["details"] // Output: map[age:30 location:India]
	// var result map[string]map[string]map[string]interface{} // Access result["user"]["details"]["age"] // Output: 30
	json.Unmarshal([]byte(jsonData), &result)
	fmt.Println(result["user"]["details"])
	// May get below error if trying to access age i.e. result["user"]["details"]["age"]
	// invalid operation: cannot index result["user"]["details"] (map index expression of type interface{})
}

// Output:
// map[age:30 location:India]
