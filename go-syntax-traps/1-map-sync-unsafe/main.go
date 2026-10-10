package main

import (
	"fmt"
	"time"
)

// The Unsafe Way (Causes a Crash)
// If you run this code, Go will eventually panic and crash because
// both goroutines are modifying the map at the same time:
func main() {
	// Maps aren't safe for unsynchronized concurrent writes
	m := make(map[string]int)
	go func() {
		for i := 0; i < 1000; i++ {
			m["a"] = i
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			m["b"] = i
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")

}

// harish $ go run main.go
// fatal error: concurrent map writes

// goroutine 8 [running]:
// internal/runtime/maps.fatal({0x10047efe1?, 0x0?})
//         /opt/homebrew/Cellar/go/1.26.6/libexec/src/runtime/panic.go:1181 +0x20
// main.main.func2()
//         ~/golang-codes/go-syntax-traps/map-sync-traps/main.go:23 +0x44
// created by main.main in goroutine 1
//         ~/golang-codes/go-syntax-traps/map-sync-traps/main.go:21 +0xa0
// exit status 2
// harish $ go run -race main.go
// ==================
// WARNING: DATA RACE
// Write at 0x00c0000a40f0 by goroutine 7:
//   runtime.mapaccess2_faststr()
//       /opt/homebrew/Cellar/go/1.26.6/libexec/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
//   main.main.func1()
//       ~/golang-codes/go-syntax-traps/map-sync-traps/main.go:18 +0x4c

// Previous write at 0x00c0000a40f0 by goroutine 8:
//   runtime.mapaccess2_faststr()
//       /opt/homebrew/Cellar/go/1.26.6/libexec/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
//   main.main.func2()
//       ~/golang-codes/go-syntax-traps/map-sync-traps/main.go:23 +0x4c

// Goroutine 7 (running) created at:
//   main.main()
//       ~/golang-codes/go-syntax-traps/map-sync-traps/main.go:16 +0x84

// Goroutine 8 (finished) created at:
//   main.main()
//       ~/golang-codes/go-syntax-traps/map-sync-traps/main.go:21 +0xe4
// ==================
// Finished
// Found 1 data race(s)
// exit status 66
// harish $
