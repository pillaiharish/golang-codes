package main

import (
	"fmt"
)

func main() {
	counts := make(map[string]int)
	for _, label := range []string{"red", "blue", "red"} {
		counts[label]++
	}
	v, ok := counts["green"]
	fmt.Println(counts["red"], v, ok)
	delete(counts, "blue")
	fmt.Println(len(counts))
}
