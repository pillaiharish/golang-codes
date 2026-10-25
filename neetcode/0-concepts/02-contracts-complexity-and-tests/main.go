package main

import (
	"fmt"
)

func main() {
	for n := 1; n <= 4; n++ {
		visits := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				visits++
			}
		}
		fmt.Println(n, visits)
	}
}
