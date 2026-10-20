package main

import (
	"fmt"
)

func main() {
	lo, hi := 2, 10
	mid := lo + (hi-lo)/2
	fmt.Println(lo, mid, hi)
	hi = mid
	fmt.Println("remaining width", hi-lo)
}
