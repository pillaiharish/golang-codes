# Q33. Debug goroutines with pprof

[Concurrency index](../README.md) · [Q29: leaks](../29-detect-goroutine-leak/README.md) · [Revision notes](../NOTES.md)

**Level 6 — production debugging.**

## Task

Run a small local HTTP server with pprof enabled, deliberately block a goroutine on a channel or mutex, and inspect goroutine profiles/stacks to locate the blocking site.

## Reference approach

```go
import (
    "log"
    "net/http"
    _ "net/http/pprof"
)

func main() {
    log.Fatal(http.ListenAndServe("127.0.0.1:6060", nil))
}
```

Run the server only in a local disposable test process, then inspect:

```sh
go tool pprof http://127.0.0.1:6060/debug/pprof/goroutine
curl 'http://127.0.0.1:6060/debug/pprof/goroutine?debug=2'
```

The `debug=2` endpoint provides goroutine stack traces with wait locations; a pprof goroutine profile helps find frequently blocked stack signatures. Correlate profiles with request logs, queue depth, active worker counts and cancellation events.

## Scenario Q&A

| Question | Answer |
|---|---|
| Does a large goroutine count prove a leak? | No. Inspect growth over time, expected concurrency and blocked stacks. |
| What is the difference between a goroutine profile and a CPU profile? | Goroutine profile shows current goroutine stacks/states; CPU profile shows where CPU time is spent. |
| Why not expose pprof publicly? | Profiles can reveal sensitive stack/operational information and impose diagnostic overhead. Bind locally or protect tightly. |
| Can pprof identify who should close a channel? | It locates blocked sites, but ownership and intended lifecycle still require code/design investigation. |

**Recall:** Describe a debugging walkthrough for a Kafka sender blocked behind a full queue: measurements, stack evidence, hypothesis, mitigation, and regression test.

References: [`net/http/pprof`](https://pkg.go.dev/net/http/pprof), [Go diagnostics](https://go.dev/doc/diagnostics).
