package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("error reading josn file", err)
	}
	var read_json map[string]interface{}
	err = json.Unmarshal(file, &read_json)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(read_json)
	fmt.Println("Method supported: ", read_json["method"])
}

// Output
// map[api_endpoint:http://example.com method:GET rate_limit:false ttl:3600]
// Method supported:  GET
