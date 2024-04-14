package main

import (
	"fmt"
	"time"
)

type data struct {
	alphabets []string
}

func main() {
	d := data{
		alphabets: []string{"a", "b", "c", "d", "e"},
	}
	for i, v := range d.alphabets {
		go fmt.Println(i, v)
	}
	time.Sleep(1 * time.Second)
}

// output
// 4 e
// 2 c
// 0 a
// 3 d
// 1 b
