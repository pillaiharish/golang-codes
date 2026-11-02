package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	duration := flag.Duration("duration", 10*time.Second, "local profiling observation window")
	flag.Parse()
	release := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-release // deliberately blocked until cleanup
	}()

	server := &http.Server{Addr: "127.0.0.1:6060"}
	go func() { _ = server.ListenAndServe() }()
	fmt.Println("Inspect: curl 'http://127.0.0.1:6060/debug/pprof/goroutine?debug=2'")
	fmt.Println("Profiling window:", *duration)
	time.Sleep(*duration) // time to inspect, not a join mechanism
	close(release)
	wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil { fmt.Println("shutdown error:", err) }
}
