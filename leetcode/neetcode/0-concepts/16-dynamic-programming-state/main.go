package main

import (
	"fmt"
)

func main() {
	type State struct{ I, J int }
	memo := make(map[State]int)
	k := State{I: 2, J: 1}
	v, exists := memo[k]
	fmt.Println(v, exists)
	memo[k] = 0
	v, exists = memo[k]
	fmt.Println(v, exists)
}
