package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

type Result struct {
	Index int
	URL string
	Status int
	Err error
}

func main() {
	// Local-only HTTP server makes the exercise reproducible offline.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/slow" { time.Sleep(250 * time.Millisecond) }
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	urls := []string{server.URL + "/a", server.URL + "/slow", server.URL + "/b", server.URL + "/c"}
	results := make([]Result, len(urls))
	client := &http.Client{}
	permits := make(chan struct{}, 2)
	var wg sync.WaitGroup
	for i, url := range urls {
		permits <- struct{}{} // bounded admission before spawning
		wg.Add(1)
		go func(index int, target string) {
			defer wg.Done()
			defer func() { <-permits }()
			ctx, cancel := context.WithTimeout(context.Background(), 120*time.Millisecond)
			defer cancel()
			result := Result{Index: index, URL: target}
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
			if err != nil {
				result.Err = err
				results[index] = result
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				result.Err = err
			} else {
				result.Status = resp.StatusCode
				resp.Body.Close()
			}
			results[index] = result // one writer per distinct index
		}(i, url)
	}
	wg.Wait()
	for _, r := range results { fmt.Printf("%d %s status=%d err=%v\n", r.Index, r.URL, r.Status, r.Err) }
}
