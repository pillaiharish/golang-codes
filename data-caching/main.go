package main

import (
	"fmt"
	"time"
)

var cache = make(map[string]interface{})

// Query for data if not found
func expensiveOperation(key string) interface{} {
	//Run query to fetch data
	time.Sleep(2 * time.Second)
	fmt.Println("Executing Query...")
	return "Result for " + key
}

// Get the the value for given key
func getValue(key string) interface{} {
	value, found := cache[key]
	if !found {
		value = expensiveOperation(key)
		cache[key] = value
	}
	return value
}

func main() {
	fmt.Println("Fetching Data...")
	// Since first value is not present getValue() will call expensiveOperation()
	fmt.Println(getValue("myData")) // slow output
	fmt.Println("Fetching Data again...")
	// now value present in cache, expensiveOperation() will not be called
	fmt.Println(getValue("myData")) // immediate output

}

// Output:
// Fetching Data...
// Executing Query...
// Result for myData
// Result for myData
