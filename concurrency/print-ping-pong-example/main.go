package main

import (
	"fmt"
	"sync"
)

func printLine(s string, wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0;i<5;i++{
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printLine("ping", &wg)

	go printLine("pong", &wg)
	wg.Wait()
}
