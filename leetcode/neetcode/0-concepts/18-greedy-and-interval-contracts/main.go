package main

import (
	"fmt"
)

func main() {
	aStart, aEnd := 1, 4
	bStart, bEnd := 4, 7
	closedOverlap := aStart <= bEnd && bStart <= aEnd
	halfOpenOverlap := aStart < bEnd && bStart < aEnd
	fmt.Println(closedOverlap, halfOpenOverlap)
}
