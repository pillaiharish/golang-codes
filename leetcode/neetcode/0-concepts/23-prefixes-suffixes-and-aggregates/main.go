package main

import (
	"fmt"
)

func main() {
	values := []int{4, 1, 7}
	prefix := make([]int, len(values)+1)
	for i, v := range values {
		prefix[i+1] = prefix[i] + v
	}
	fmt.Println(prefix)
	fmt.Println(prefix[3] - prefix[1])
}
