package main

import (
	"fmt"
	"time"
)

func DetectBot(s string) {
	fmt.Println(s)
}

func main() {
	go DetectBot("Are you a robot?")
	time.Sleep(1 * time.Second)
	DetectBot("Are you Human?")
}
