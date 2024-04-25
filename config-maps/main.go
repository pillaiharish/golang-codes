package main

import (
	"fmt"
)

func main() {
	config := map[string]string{ // map[string]interface{} is also valid here
		"api_endpoint": "http://www.go.dev",
		"method":       "GET",
	}
	fmt.Println("API endpoint is:", config["api_endpoint"])
	fmt.Println("Method supported:", config["method"])

}

// Output:
// API endpoint is: http://www.go.dev
// Method supported: GET
