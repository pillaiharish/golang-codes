package main

import (
	"fmt"
)

func main() {
	// using a map to make a set
	set := make(map[string]bool)

	// Adding elements to the set
	set["member1"] = true
	set["member2"] = true

	//try to add a duplicate entry
	if !set["member1"] {
		set["member1"] = false
	}

	fmt.Println(set)
}

// Output:
// map[member1:true member2:false]
