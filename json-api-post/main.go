package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{Name: "abhi", Age: 30}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Sending data... ", user)

	// Simulating sending JSON data to an API endpoint
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var jsonBody map[string]interface{}
	if resp.StatusCode == http.StatusCreated {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		err = json.Unmarshal(bodyBytes, &jsonBody)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Created resource with ID ", jsonBody["id"])
		fmt.Println("Received response...")
		fmt.Println(bodyString)
	}

}

// Output:
// Sending data...  {abhi 30}
// Created resource with ID  101
// Received response...
// {
//   "name": "abhi",
//   "age": 30,
//   "id": 101
// }
