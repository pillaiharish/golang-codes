package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	mode := flag.String("mode", "mutex", "mutex, atomic, or unsafe (deliberate race)")
	flag.Parse()
	var wg sync.WaitGroup
	var mu sync.Mutex
	var plain int
	var count atomic.Int64

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			switch *mode {
			case "atomic":
				count.Add(1)
			case "unsafe":
				plain++ // deliberately racy; try go run -race . -mode=unsafe
			default:
				mu.Lock()
				plain++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	if *mode == "atomic" {
		fmt.Println("counter:", count.Load())
	} else {
		fmt.Println("counter:", plain)
	}
}
