package main

import (
	"fmt"
)

func printLine(s string, ch chan <-struct{}) {
	for i:=0;i<5;i++ {
		fmt.Println(s)
	}
	ch <- struct{}{}
}

func main() {
	ch := make(chan struct{})
	go printLine("ping", ch)
	go printLine("pong", ch)
		<- ch
		<-ch
}
