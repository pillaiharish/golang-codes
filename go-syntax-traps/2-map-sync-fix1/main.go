package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	m := make(map[string]int)
	var mu sync.Mutex
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			m["a"] = i
			mu.Unlock()
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			m["b"] = i
			mu.Unlock()
		}
	}()
	time.Sleep(1 * time.Microsecond)
	mu.Lock()
	fmt.Println("Map contents:", m)
	mu.Unlock()

}

// Output1: When we use sleep of time.Sleep(1 * time.Millisecond)
// harish $ go run -race main.go
// Map contents: map[a:987 b:735]

// Output2: When we dont use any sleep
// harish $ go run -race main.go
// Map contents: map[]

// Output3: When we use time.Sleep(1 * time.Second)
// harish $ go run -race main.go
// Map contents: map[a:999 b:999]

// Output4: When we use time.Sleep(1 * time.Nanosecond)
// harish $ go run -race main.go
// Map contents: map[b:13]
// harish $ go run -race main.go
// Map contents: map[b:1]
// harish $ go run -race main.go
// Map contents: map[]
// harish $ go run -race main.go
// Map contents: map[]
// harish $ go run -race main.go
// Map contents: map[a:0 b:39]

// Output5: When we use time.Sleep(1 * time.Microsecond)
// harish $ go run -race main.go
// Map contents: map[a:0]
// harish $ go run -race main.go
// Map contents: map[a:14 b:72]
// harish $ go run -race main.go
// Map contents: map[a:0]
// harish $ go run -race main.go
// Map contents: map[b:12]
// harish $ go run -race main.go
// Map contents: map[]
// harish $ go run -race main.go
// Map contents: map[]
// harish $ go run -race main.go
// Map contents: map[a:0]
// harish $ go run -race main.go
// Map contents: map[a:7 b:24]
// harish $ go run -race main.go
// Map contents: map[a:52 b:115]
// harish $ go run -race main.go
// Map contents: map[a:33]
