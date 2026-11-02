package main

import (
	"fmt"
	"sync"
)

func main() {
	aDone := make(chan struct{})
	bDone := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("A")
		close(aDone)
	}()
	go func() {
		defer wg.Done()
		<-aDone // B cannot print until A has printed.
		fmt.Println("B")
		close(bDone)
	}()
	go func() {
		defer wg.Done()
		<-bDone // C cannot print until B has printed.
		fmt.Println("C")
	}()

	wg.Wait()
}
