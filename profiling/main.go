package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	h, err := os.Create("heap.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}

	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// Simulate some load
	for i := 0; i < 1000000; i++ {
		_ = time.Now().UnixNano()
	}

	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(h); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	defer h.Close()

	select {}
}
