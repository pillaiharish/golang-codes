package main

import (
	"fmt"
)

func main() {
	var x uint32 = 10
	mask := uint32(1) << 1
	fmt.Printf("%04b %04b %04b\n", x&mask, x|1, x^mask)
	fmt.Println(x >> 1)
}
