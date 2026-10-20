package main

import "fmt"

func main() {
	ct := NewCounter()
	ct.Increment("a")
	ct.Increment("a")
	ct.Increment("a")
	ct.Increment("a")
	fmt.Println(ct.GetKey("a")) // 4
}
