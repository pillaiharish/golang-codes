package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	mode := flag.String("mode", "rwmutex", "mutex, rwmutex, syncmap, or unsafe")
	flag.Parse()
	var mu sync.RWMutex
	values := make(map[int]int)
	var concurrent sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			switch *mode {
			case "syncmap":
				concurrent.Store(key, key*2)
				_, _ = concurrent.Load(key)
			case "unsafe":
				values[key] = key * 2 // intentionally unsafe concurrent writes
			case "mutex":
				mu.Lock()
				values[key] = key * 2
				_, _ = values[key]
				mu.Unlock()
			default:
				mu.Lock()
				values[key] = key * 2
				mu.Unlock()
				mu.RLock()
				_, _ = values[key]
				mu.RUnlock()
			}
		}(i)
	}
	wg.Wait()
	if *mode == "syncmap" {
		length := 0
		concurrent.Range(func(_, _ any) bool { length++; return true })
		fmt.Println("sync.Map entries:", length)
	} else {
		fmt.Println("map entries:", len(values))
	}
}
