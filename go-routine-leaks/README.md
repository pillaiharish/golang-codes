## What is a goroutine leak?

A **goroutine leak** happens when a goroutine is started but never finishes because it becomes blocked forever or runs forever without an exit condition.

A leaked goroutine continues to occupy:

- its stack memory,
- any variables it references,
- possibly OS resources such as network connections,
- scheduler bookkeeping.

A small leak may be harmless, but many leaked goroutines can cause high memory usage, high CPU usage, slow scheduling, file descriptor exhaustion, or eventual failure.

Example:

```go
func leak() {
    ch := make(chan int)

    go func() {
        ch <- 42 // blocked forever: nobody receives from ch
    }()
}
```

When `leak()` returns, the spawned goroutine is still alive, blocked on `ch <- 42`. It will never exit.

---

# How to know whether goroutines leaked?

## 1. Watch `runtime.NumGoroutine()`

The simplest signal is: the goroutine count keeps increasing and does not come down.

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func leak() {
    ch := make(chan int)
    go func() {
        ch <- 1 // no receiver, leaks
    }()
}

func main() {
    fmt.Println("before:", runtime.NumGoroutine())

    for i := 0; i < 100; i++ {
        leak()
    }

    time.Sleep(time.Second)

    fmt.Println("after:", runtime.NumGoroutine())
}
```

Possible output:

```text
before: 1
after: 101
```

If after requests, tests, or workload completion the count remains high, you likely have a leak.

---

## 2. Use `pprof` goroutine profile

This is the most common way to find **where** goroutines are stuck.

Add this to your program:

```go
package main

import (
    "fmt"
    "net/http"
    _ "net/http/pprof"
    "runtime"
    "time"
)

func leak() {
    ch := make(chan int)

    go func() {
        ch <- 42 // blocked forever
    }()
}

func main() {
    go http.ListenAndServe("localhost:6060", nil)

    for i := 0; i < 100; i++ {
        leak()
    }

    time.Sleep(time.Second)
    fmt.Println("goroutines:", runtime.NumGoroutine())

    select {} // keep program alive
}
```

Then inspect:

```bash
# Human-readable summary
curl "http://localhost:6060/debug/pprof/goroutine?debug=1"

# Full stack traces of all goroutines
curl "http://localhost:6060/debug/pprof/goroutine?debug=2"
```

Or use `go tool pprof`:

```bash
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

Inside pprof:

```text
(pprof) top
(pprof) traces
(pprof) list leak
```

You may see something like:

```text
goroutine 19 [chan send]:
main.leak.func1()
        /path/main.go:14
```

That tells you: there are goroutines blocked while sending to a channel.

---

## 3. Dump all goroutine stacks manually

If you cannot use HTTP pprof, you can print stacks directly:

```go
package main

import (
    "os"
    "runtime"
    "time"
)

func leak() {
    ch := make(chan int)
    go func() {
        ch <- 1
    }()
}

func dumpGoroutines() {
    buf := make([]byte, 1<<20)
    n := runtime.Stack(buf, true) // true = all goroutines
    os.Stdout.Write(buf[:n])
}

func main() {
    for i := 0; i < 10; i++ {
        leak()
    }

    time.Sleep(time.Second)
    dumpGoroutines()
}
```

Look for repeated stacks like:

```text
goroutine 8 [chan send]:
main.leak.func1()
```

---

## 4. Use `goleak` in tests

For unit tests, Uber’s `goleak` library is very useful.

Install:

```bash
go get go.uber.org/goleak
```

Example:

```go
package main

import (
    "testing"

    "go.uber.org/goleak"
)

func leak() {
    ch := make(chan int)
    go func() {
        ch <- 1
    }()
}

func TestNoLeaks(t *testing.T) {
    defer goleak.VerifyNone(t)

    leak()
}
```

This test will fail because `leak()` leaves a goroutine running.

You can also use:

```go
func TestMain(m *testing.M) {
    goleak.VerifyTestMain(m)
}
```

---

## 5. Use the execution tracer

For complicated cases, use `runtime/trace`.

```go
package main

import (
    "os"
    "runtime/trace"
)

func main() {
    f, _ := os.Create("trace.out")
    trace.Start(f)
    defer trace.Stop()

    // your code here

    trace.Stop()
}
```

Run:

```bash
go run main.go
go tool trace trace.out
```

The trace UI can show goroutine creation, blocking, scheduling, and completion.

---

# Full example: leak, detection, and fix

## Leaking version

```go
package main

import (
    "fmt"
    "net/http"
    _ "net/http/pprof"
    "runtime"
    "time"
)

func leak() {
    ch := make(chan int)

    go func() {
        ch <- 42 // blocked forever because nobody receives
    }()
}

func main() {
    go http.ListenAndServe("localhost:6060", nil)

    for i := 0; i < 50; i++ {
        leak()
    }

    time.Sleep(time.Second)
    fmt.Println("goroutines:", runtime.NumGoroutine())

    select {}
}
```

Run:

```bash
go run main.go
```

Then:

```bash
curl "http://localhost:6060/debug/pprof/goroutine?debug=2"
```

You will see many goroutines blocked in `chan send`.

---

## Fixed version: ensure someone receives

```go
func fixed() {
    ch := make(chan int)

    go func() {
        ch <- 42
    }()

    v := <-ch
    _ = v
}
```

Now the goroutine can complete.

---

## Fixed version: use context cancellation

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context) {
    ch := make(chan int)

    go func() {
        select {
        case ch <- 42:
            fmt.Println("sent")
        case <-ctx.Done():
            fmt.Println("cancelled")
            return
        }
    }()
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    worker(ctx)

    time.Sleep(3 * time.Second)
}
```

If nobody receives from `ch`, the goroutine exits when the context is cancelled.

---

# Common situations that cause goroutine leaks

## 1. Sending to a channel with no receiver

```go
func leak() {
    ch := make(chan int)

    go func() {
        ch <- 1 // blocked forever
    }()
}
```

### Why it leaks

Unbuffered channels block until both sender and receiver are ready. If there is no receiver, the sender blocks forever.

### Fix

Either receive:

```go
func fixed() {
    ch := make(chan int)

    go func() {
        ch <- 1
    }()

    <-ch
}
```

Or use a buffered channel if appropriate:

```go
func fixed() {
    ch := make(chan int, 1)

    go func() {
        ch <- 1
    }()
}
```

But be careful: buffered channels only delay the problem if the buffer fills up and nobody reads.

Better:

```go
func fixed(ctx context.Context) {
    ch := make(chan int)

    go func() {
        select {
        case ch <- 1:
        case <-ctx.Done():
            return
        }
    }()
}
```

---

## 2. Receiving from a channel with no sender

```go
func leak() {
    ch := make(chan int)

    go func() {
        v := <-ch // blocked forever
        _ = v
    }()
}
```

### Why it leaks

A receive from an unbuffered channel blocks until a sender sends. If nobody sends and the channel is never closed, the goroutine waits forever.

### Fix

Send a value:

```go
func fixed() {
    ch := make(chan int)

    go func() {
        v := <-ch
        _ = v
    }()

    ch <- 1
}
```

Or close the channel if appropriate:

```go
func fixed() {
    ch := make(chan int)

    go func() {
        v, ok := <-ch
        _ = v
        _ = ok
    }()

    close(ch)
}
```

Or use cancellation:

```go
func fixed(ctx context.Context) {
    ch := make(chan int)

    go func() {
        select {
        case v := <-ch:
            _ = v
        case <-ctx.Done():
            return
        }
    }()
}
```

---

## 3. Worker pool channels are never closed

```go
type Job struct{}

func handle(Job) {}

func leak() {
    jobs := make(chan Job)

    for i := 0; i < 10; i++ {
        go func() {
            for job := range jobs {
                handle(job)
            }
        }()
    }

    // jobs is never closed
    // if no more jobs are sent, workers wait forever on range jobs
}
```

### Why it leaks

`range jobs` blocks waiting for more jobs. If the channel is never closed, workers never exit.

### Fix

Close the jobs channel when no more jobs will be sent.

```go
package main

import (
    "fmt"
    "sync"
)

type Job struct{}

func handle(Job) {}

func fixed() {
    jobs := make(chan Job)

    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()

            for job := range jobs {
                handle(job)
            }
        }()
    }

    for i := 0; i < 5; i++ {
        jobs <- Job{}
    }

    close(jobs)

    wg.Wait()
    fmt.Println("workers finished")
}
```

---

## 4. Forgetting to cancel a context

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func leak() {
    ctx := context.Background()

    go func() {
        <-ctx.Done() // context.Background() is never done
        fmt.Println("done")
    }()
}

func main() {
    leak()
    time.Sleep(time.Second)
}
```

### Why it leaks

`context.Background()` is never cancelled. The goroutine waits forever.

### Fix

Use a cancellable context and call `cancel`.

```go
func fixed() {
    ctx, cancel := context.WithCancel(context.Background())

    go func() {
        <-ctx.Done()
        fmt.Println("done")
    }()

    cancel()
}
```

Or use a timeout:

```go
func fixed() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    go func() {
        <-ctx.Done()
        fmt.Println("done or timed out")
    }()
}
```

---

## 5. HTTP requests without timeout

```go
package main

import (
    "net/http"
)

func leak() {
    go func() {
        resp, err := http.Get("http://example.com")
        if err != nil {
            return
        }
        defer resp.Body.Close()
    }()
}
```

### Why it can leak

If the remote server hangs, the goroutine can block for a very long time or forever, depending on OS/network behavior.

### Fix

Use a client with timeout:

```go
package main

import (
    "net/http"
    "time"
)

var client = &http.Client{
    Timeout: 5 * time.Second,
}

func fixed() {
    go func() {
        resp, err := client.Get("http://example.com")
        if err != nil {
            return
        }
        defer resp.Body.Close()
    }()
}
```

Better: use request context.

```go
package main

import (
    "context"
    "net/http"
    "time"
)

func fixed(parent context.Context) {
    ctx, cancel := context.WithTimeout(parent, 5*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://example.com", nil)
    if err != nil {
        return
    }

    go func() {
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            return
        }
        defer resp.Body.Close()
    }()
}
```

---

## 6. Blocking network I/O without deadlines

```go
package main

import (
    "net"
)

func leak() {
    go func() {
        conn, err := net.Dial("tcp", "example.com:80")
        if err != nil {
            return
        }
        defer conn.Close()

        buf := make([]byte, 1024)
        conn.Read(buf) // can block for a long time
    }()
}
```

### Fix

Use deadlines or context-aware dialers.

```go
func fixed() {
    go func() {
        conn, err := net.Dial("tcp", "example.com:80")
        if err != nil {
            return
        }
        defer conn.Close()

        conn.SetReadDeadline(time.Now().Add(5 * time.Second))

        buf := make([]byte, 1024)
        _, err = conn.Read(buf)
        if err != nil {
            return
        }
    }()
}
```

---

## 7. Nil channel blocks forever

```go
func leak() {
    var ch chan int

    go func() {
        <-ch // blocked forever
    }()
}
```

### Why it leaks

Reading from or writing to a `nil` channel blocks forever.

### Fix

Initialize the channel before using it.

```go
func fixed() {
    ch := make(chan int)

    go func() {
        <-ch
    }()

    ch <- 1
}
```

Sometimes people intentionally use `nil` channels in `select` to disable a case. That is fine only if another case can eventually unblock the `select`.

Example of dangerous use:

```go
func leak() {
    var ch chan int

    go func() {
        select {
        case <-ch: // nil channel, never ready
        }
    }()
}
```

---

## 8. Waiting on a mutex that is never unlocked

```go
package main

import (
    "sync"
)

func leak() {
    var mu sync.Mutex

    mu.Lock()

    go func() {
        mu.Lock() // blocked forever if first Lock is never unlocked
        defer mu.Unlock()
    }()
}
```

### Why it leaks

The second goroutine waits forever for the mutex.

### Fix

Unlock the mutex.

```go
func fixed() {
    var mu sync.Mutex

    mu.Lock()

    go func() {
        mu.Lock()
        defer mu.Unlock()
    }()

    mu.Unlock()
}
```

In real code, use `defer` carefully:

```go
func doWork(mu *sync.Mutex) {
    mu.Lock()
    defer mu.Unlock()

    // work
}
```

---

## 9. Infinite loop without shutdown condition

```go
func leak() {
    go func() {
        for {
            // do something
        }
    }()
}
```

### Why it leaks

The goroutine never exits.

### Fix

Provide a stop channel or context.

```go
func fixed(stop <-chan struct{}) {
    go func() {
        for {
            select {
            case <-stop:
                return
            default:
                // do work
            }
        }
    }()
}
```

Better with ticker:

```go
package main

import (
    "time"
)

func fixed(stop <-chan struct{}) {
    go func() {
        ticker := time.NewTicker(time.Second)
        defer ticker.Stop()

        for {
            select {
            case <-ticker.C:
                // do periodic work
            case <-stop:
                return
            }
        }
    }()
}
```

---

## 10. Using `select {}` forever

```go
func leak() {
    go func() {
        select {} // blocks forever
    }()
}
```

### Why it leaks

An empty `select` blocks forever.

Sometimes this is intentional in `main()` to keep a program alive, but inside worker goroutines it usually means a leak.

### Fix

Use a done/context channel.

```go
func fixed(stop <-chan struct{}) {
    go func() {
        select {
        case <-stop:
            return
        }
    }()
}
```

---

## 11. WaitGroup misuse

```go
package main

import (
    "sync"
)

func leak() {
    var wg sync.WaitGroup

    wg.Add(1)

    go func() {
        // forgot wg.Done()
    }()

    wg.Wait() // blocks forever
}
```

### Why it leaks/blocks

If `wg.Done()` is not called, `wg.Wait()` never returns.

If `wg.Wait()` is inside a goroutine, that goroutine leaks. If it is in `main`, the program hangs.

### Fix

Use `defer wg.Done()`.

```go
func fixed() {
    var wg sync.WaitGroup

    wg.Add(1)

    go func() {
        defer wg.Done()
    }()

    wg.Wait()
}
```

Also, call `wg.Add(1)` before starting the goroutine, not inside it.

Bad:

```go
var wg sync.WaitGroup

go func() {
    wg.Add(1) // race: Wait may happen before Add
    defer wg.Done()
}()

wg.Wait()
```

Good:

```go
var wg sync.WaitGroup

wg.Add(1)

go func() {
    defer wg.Done()
}()

wg.Wait()
```

---

## 12. Producer keeps sending after consumer stopped

```go
package main

import "fmt"

func leak() {
    ch := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            ch <- i // may block forever if consumer stops early
        }
    }()

    v := <-ch
    fmt.Println(v)

    // function returns, consumer stopped
    // producer may be blocked on second send
}
```

### Why it leaks

The consumer reads one value and exits. The producer then blocks on the next send.

### Fix

Use context cancellation.

```go
package main

import (
    "context"
    "fmt"
)

func fixed(ctx context.Context) {
    ch := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            select {
            case ch <- i:
            case <-ctx.Done():
                return
            }
        }
    }()

    select {
    case v := <-ch:
        fmt.Println(v)
    case <-ctx.Done():
    }
}
```

Caller:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

fixed(ctx)
```

---

## 13. `time.Tick` in a loop that should stop

```go
func leak() {
    go func() {
        for range time.Tick(time.Second) {
            // do something
        }
    }()
}
```

### Why it leaks

`time.Tick` provides no way to stop the ticker. The goroutine will keep running forever unless the whole program exits.

### Fix

Use `time.NewTicker` and stop it.

```go
func fixed(stop <-chan struct{}) {
    go func() {
        ticker := time.NewTicker(time.Second)
        defer ticker.Stop()

        for {
            select {
            case <-ticker.C:
                // do periodic work
            case <-stop:
                return
            }
        }
    }()
}
```

---

# How to identify the leaking location in practice

When you suspect a leak, follow this workflow:

## Step 1: Confirm the count grows

```go
fmt.Println(runtime.NumGoroutine())
```

Run the operation many times:

```go
for i := 0; i < 1000; i++ {
    doSomething()
}
```

Then check again after some idle time:

```go
time.Sleep(5 * time.Second)
fmt.Println(runtime.NumGoroutine())
```

If the count stays high, investigate.

---

## Step 2: Capture goroutine stacks

Use:

```bash
curl "http://localhost:6060/debug/pprof/goroutine?debug=2" > goroutines.txt
```

Then search for repeated patterns:

```bash
grep -A 5 "chan send" goroutines.txt
grep -A 5 "chan receive" goroutines.txt
grep -A 5 "select" goroutines.txt
grep -A 5 "IO wait" goroutines.txt
grep -A 5 "semacquire" goroutines.txt
```

Common meanings:

| Stack state | Likely cause |
|---|---|
| `chan send` | goroutine blocked sending to a channel |
| `chan receive` | goroutine blocked receiving from a channel |
| `select` | goroutine waiting in a `select` statement |
| `IO wait` | network/file I/O waiting |
| `semacquire` | waiting on synchronization primitive, often mutex |
| `sleep` | sleeping, possibly infinite or too long |

---

## Step 3: Find repeated stacks

If you see this repeated many times:

```text
goroutine 100 [chan send]:
main.leak.func1()
```

then the leak source is likely inside `main.leak.func1`.

If you see this:

```text
goroutine 200 [chan receive]:
main.worker()
```

then workers are waiting for jobs or signals that never arrive.

---

# Common leak patterns and fixes

## Pattern: channel send without receiver

Leak:

```go
go func() {
    ch <- value
}()
```

Fix:

```go
select {
case ch <- value:
case <-ctx.Done():
    return
}
```

---

## Pattern: channel receive without sender

Leak:

```go
go func() {
    value := <-ch
}()
```

Fix:

```go
select {
case value := <-ch:
    _ = value
case <-ctx.Done():
    return
}
```

---

## Pattern: worker waiting for jobs forever

Leak:

```go
go func() {
    for job := range jobs {
        handle(job)
    }
}()
```

Fix:

```go
close(jobs)
```

or:

```go
go func() {
    for {
        select {
        case job := <-jobs:
            handle(job)
        case <-ctx.Done():
            return
        }
    }
}()
```

---

## Pattern: HTTP call without timeout

Leak:

```go
go func() {
    resp, err := http.Get(url)
    ...
}()
```

Fix:

```go
client := &http.Client{
    Timeout: 5 * time.Second,
}
```

or:

```go
req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
resp, err := client.Do(req)
```

---

# Prevention checklist

When starting a goroutine, always ask:

1. How will this goroutine exit?
2. What happens if the parent request is cancelled?
3. What happens if the channel receiver disappears?
4. What happens if the channel sender disappears?
5. Is this goroutine waiting on a network call without timeout?
6. Is there a `select` with a `ctx.Done()` or `stop` case?
7. If using channels, who closes them?
8. If using `WaitGroup`, is `Done()` always called?
9. If using tickers, are they stopped?
10. Can this goroutine be exercised in a test with `goleak`?

---

# Summary

A **goroutine leak** is a goroutine that never terminates, usually because it is blocked forever on:

- channel send,
- channel receive,
- `select`,
- mutex,
- network I/O,
- nil channel,
- infinite loop,
- uncancelled context,
- missing shutdown signal.

To find leaked goroutines:

- check `runtime.NumGoroutine()`,
- use `net/http/pprof` and inspect `/debug/pprof/goroutine`,
- dump stacks with `runtime.Stack`,
- use `goleak` in tests,
- use `runtime/trace` for difficult cases.

The stack traces are the key: they show exactly where goroutines are stuck, such as `chan send`, `chan receive`, `select`, or `IO wait`.