package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}
	i := 0
	if i++; i < 3 {
		fmt.Println("init", i)
	}
	var callbacks []func() int
	for i := 0; i < 3; i++ {
		callbacks = append(callbacks, func() int { return i })
	}
	for _, f := range callbacks {
		fmt.Println("declared", f())
	}
	var shared int
	var reused []func() int
	for shared = 0; shared < 3; shared++ {
		reused = append(reused, func() int { return shared })
	}
	for _, f := range reused {
		fmt.Println("reused", f())
	}
}
