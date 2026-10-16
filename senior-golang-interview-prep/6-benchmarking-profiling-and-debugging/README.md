[← Senior Go interview preparation index](../README.md) · [Personal notes](NOTES.md) · [Revision log](../REVISION-LOG.md)

> Study notes adapted from the supplied transcript. Technical examples are preserved for review; code blocks are illustrative unless separately validated.

# 6 — Benchmarking, Profiling, and Production Debugging

This module is about evidence.

A senior engineer does not say:

> “I think this is faster.”

They say:

> “Here is the benchmark, here is the profile, here is the before/after comparison, and here is how we will monitor it in production.”

---

## 6.1 Benchmarking fundamentals

Basic benchmark command:

```bash
go test -run '^$' -bench=. -benchmem ./...
```

- `-run '^$'` skips normal tests
- `-bench=.` runs all benchmarks
- `-benchmem` reports allocations

Example benchmark:

```go
func BenchmarkEncode(b *testing.B) {
    event := Event{
        ID:      "abc",
        Payload: []byte("hello"),
    }

    b.ReportAllocs()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _, err := encode(event)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

Better modern form where available:

```go
func BenchmarkEncode(b *testing.B) {
    event := Event{
        ID:      "abc",
        Payload: []byte("hello"),
    }

    b.ReportAllocs()

    for b.Loop() {
        _, _ = encode(event)
    }
}
```

`b.Loop()` is the newer style and helps avoid some timing pitfalls.

---

## 6.2 Benchmark traps

### Trap 1: Compiler optimizing away the work

If the result is unused, the compiler may eliminate the work.

Use a package-level sink:

```go
var sink []byte

func BenchmarkEncode(b *testing.B) {
    b.ReportAllocs()

    for b.Loop() {
        out, _ := encode(testEvent)
        sink = out
    }
}
```

---

### Trap 2: Measuring setup instead of the hot path

Do not put expensive setup inside the benchmark loop unless you intend to measure it.

Use `b.ResetTimer()` if necessary, but prefer clean setup outside the loop.

---

### Trap 3: Benchmarks that do not resemble production

A benchmark with one goroutine may look great, while production fails under concurrency.

For concurrent systems, use:

```go
b.RunParallel(func(pb *testing.PB) {
    for pb.Next() {
        // operation
    }
})
```

To control parallelism:

```go
b.SetParallelism(64)
```

This is especially relevant for Module 7, where you want throughput at concurrency 1, 4, 16, and 64.

---

### Trap 4: Comparing unstable results

Use multiple runs:

```bash
go test -run '^$' -bench=. -benchmem -count=10 ./...
```

Then compare with `benchstat`:

```bash
go install golang.org/x/perf/cmd/benchstat@latest

go test -run '^$' -bench=. -count=10 ./... > old.txt
# make change
go test -run '^$' -bench=. -count=10 ./... > new.txt

benchstat old.txt new.txt
```

A senior engineer does not trust one noisy benchmark run.

---

## 6.3 Race detector

Run tests with:

```bash
go test -race ./...
```

The race detector dynamically detects data races during execution.

### Important limitations

1. It only detects races on code paths that actually run.
2. It does not prove absence of races.
3. It does not detect deadlocks.
4. It does not detect logic-level concurrency bugs.
5. It adds significant CPU and memory overhead.
6. It should be used in tests and CI, not normally in production.

Good senior phrasing:

> “The race detector is necessary but not sufficient. It proves that exercised interleavings are race-free; it cannot prove all interleavings.”

To stress concurrency tests:

```bash
go test -race -count=20 ./...
```

or vary CPU count:

```bash
go test -race -cpu=1,4,16 ./...
```

---

## 6.4 CPU profiling

### Enable pprof in a service

```go
import _ "net/http/pprof"

go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

Security note: do not expose this endpoint publicly. Bind it to localhost, an internal admin port, or protect it.

### Capture CPU profile

```bash
curl -o cpu.prof http://localhost:6060/debug/pprof/profile?seconds=30
go tool pprof -http=:8080 cpu.prof
```

Or directly:

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile?seconds=30
```

### What CPU profile tells you

It shows where the program spends CPU time.

Typical useful views:

- flame graph
- top functions
- source listing
- call graph

Things to look for:

- application hotspots
- JSON encoding/decoding
- compression
- TLS
- regex
- `runtime.mallocgc`
- GC marking functions
- lock-related functions
- syscall time

---

## 6.5 Heap and allocation profiling

Heap profile is not one view. It has multiple useful perspectives.

### Currently retained memory

Default is often `inuse_space`:

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap
```

This helps with leaks and retained memory.

### Total allocations over time

Use `alloc_space`:

```bash
go tool pprof -http=:8080 "http://localhost:6060/debug/pprof/heap?seconds=0"
```

Inside pprof:

```text
sample_index=alloc_space
```

Or use:

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/allocs
```

This helps with GC pressure.

### Important distinction

| Profile view | Best for |
|---|---|
| `inuse_space` | memory leaks, retained heap |
| `inuse_objects` | object retention |
| `alloc_space` | allocation pressure, GC workload |
| `alloc_objects` | allocation frequency |

A senior answer should distinguish **allocation rate** from **retained memory**.

They are related but not the same problem.

---

## 6.6 Goroutine dumps

Useful when you suspect:

- goroutine leaks
- deadlocks
- blocked workers
- runaway fan-out

### Text dump

```bash
curl http://localhost:6060/debug/pprof/goroutine?debug=1
```

For more detailed stacks:

```bash
curl http://localhost:6060/debug/pprof/goroutine?debug=2
```

### Interactive profile

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/goroutine
```

### What to look for

Many goroutines stuck in the same place, such as:

- `chan receive`
- `chan send`
- `select`
- `sync.(*Mutex).Lock`
- `net/http.(*conn).serve`
- `runtime.gopark`
- `semacquire`

The pattern tells you the category:

- channel misuse
- missing receiver
- lock contention
- HTTP client hanging
- shutdown bug

---

## 6.7 Mutex and block profiling

These are essential for diagnosing latency that is not CPU-bound.

### Enable mutex profiling

```go
runtime.SetMutexProfileFraction(5)
```

Non-zero enables sampling. Lower values sample more. Use carefully in production.

### Enable block profiling

```go
runtime.SetBlockProfileRate(1000)
```

This samples blocking events such as:

- mutex acquisition
- channel sends/receives
- selects
- other synchronization stalls

### Capture profiles

```bash
curl -o mutex.prof http://localhost:6060/debug/pprof/mutex
curl -o block.prof http://localhost:6060/debug/pprof/block

go tool pprof -http=:8080 mutex.prof
go tool pprof -http=:8080 block.prof
```

### What they tell you

- Mutex profile: where lock contention is being caused and where locks are held.
- Block profile: where goroutines are waiting.

If CPU is low but latency is high, these profiles are often more useful than CPU profile.

---

## 6.8 Execution tracer

The execution tracer is powerful for concurrency and scheduler issues.

Capture:

```bash
curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
go tool trace trace.out
```

It can help with:

- scheduler latency
- goroutine wakeup latency
- GC phases
- syscalls
- network poller behavior
- starvation-like patterns

For advanced debugging, it is one of the strongest tools.

---

## 6.9 Using evidence to distinguish failure modes

This is a key senior skill.

### Case 1: CPU saturation

Evidence:

- high CPU usage
- CPU profile dominated by application functions
- many runnable goroutines
- throughput may rise with load until saturation

Likely causes:

- hot code paths
- inefficient algorithms
- too much serialization/compression
- excessive JSON work
- GC due to high allocation rate

---

### Case 2: Blocking / contention

Evidence:

- latency high
- CPU not high
- goroutines stuck in lock/channel/select
- block/mutex profile shows waiting
- trace shows goroutines parked

Likely causes:

- lock held too long
- full bounded channel
- missing receiver
- downstream stall propagating backwards
- shutdown path blocking workers

---

### Case 3: Allocation/GC overhead

Evidence:

- CPU profile shows `runtime.mallocgc`, GC marking functions
- heap alloc profile shows high allocation rate
- GC trace shows frequent GC and high GC CPU
- p99 spikes correlate with GC activity

Likely causes:

- per-event allocations
- repeated string formatting
- buffer misuse
- JSON encoding without reuse
- slice growth without preallocation

---

## Interview questions and answers

## Question 1: An API’s p99 latency suddenly doubles. How would you investigate?

### Strong senior answer

I would start by narrowing the scope.

1. **Confirm the change**
   - Is it all endpoints or one?
   - Did it happen after a deploy?
   - Is traffic higher?
   - Are downstream dependencies slower?
   - Are pod restarts, OOMs, or throttling happening?

2. **Check saturation metrics**
   - CPU
   - memory
   - GC CPU
   - goroutine count
   - queue depth
   - FD/socket usage
   - Kafka producer latency

3. **Compare profiles before and after**
   - CPU profile
   - heap alloc profile
   - mutex/block profile
   - goroutine profile

4. **Look for the dominant category**
   - CPU-bound: CPU profile shows app functions
   - allocation-bound: CPU profile shows runtime allocation/GC
   - contention-bound: mutex/block profile and parked goroutines
   - dependency-bound: spans waiting on Kafka/DB/network

5. **Reproduce**
   - load test
   - benchmark the suspect function
   - isolate with minimal repro

6. **Fix and verify**
   - measure p50/p95/p99 before/after
   - compare CPU and allocation profiles
   - add alerts/metrics to prevent regression

The important part is that I do not guess first. I classify the problem first.

---

## Question 2: A pod uses 90% CPU while request throughput remains flat. What do you inspect?

### Strong senior answer

First, take a CPU profile.

If CPU profile shows application code, then there is a hotspot.

If CPU profile shows runtime allocation or GC, then inspect heap allocation profile and GC behavior.

If CPU profile shows syscalls, inspect external calls, file I/O, DNS, or kernel interaction.

If CPU profile is not revealing much but CPU is still high, I would inspect:

- goroutine count
- background workers
- retry loops
- busy loops without blocking
- timers
- channel spin patterns
- GC trace
- `GOMAXPROCS` relative to container CPU quota

A common production issue is Go seeing too many CPUs inside a container, causing excessive threading and CFS throttling. Tools like `automaxprocs` or explicit `GOMAXPROCS` tuning can help.

The senior point is:

> High CPU with flat throughput usually means the service is doing extra work that is not producing useful request progress — GC, retries, contention, or background loops.

---

## Question 3: What can a heap profile tell you that a CPU profile cannot?

### Strong senior answer

A heap profile tells you about memory behavior:

- where objects are allocated
- what is currently retained
- which sites allocate the most bytes
- which sites allocate the most objects
- whether memory growth comes from retained objects or high churn

A CPU profile tells you where time is spent, but it does not directly show:

- retained memory
- leak sources
- object lifetime
- allocation frequency by call site
- whether memory growth is intentional cache growth or unintended retention

For GC pressure, I want allocation profiles. For leaks, I want retained in-use profiles, often compared over time.

---

## Question 4: How do you identify lock contention?

### Strong senior answer

I would look at:

1. Mutex profile:

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/mutex
```

2. Block profile:

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/block
```

3. Goroutine profile:

Many goroutines stuck in:

```text
sync.(*Mutex).Lock
runtime.gopark
semacquire
```

4. CPU profile:

If contention causes spinning or lock-related CPU overhead, it may appear there too.

5. Application metrics:

- latency rising while CPU is not saturated
- queue depth increasing
- throughput flattening as concurrency increases

Then I would fix by:

- shrinking critical sections
- sharding locks
- using atomics where appropriate
- using copy-on-write for read-mostly state
- moving I/O out of locks

---

## Question 5: A Go service’s memory usage grows continuously. How do you distinguish a genuine memory leak from expected heap growth?

### Strong senior answer

I would distinguish several possibilities:

1. **Retained heap growth**
   - `HeapInuse` grows
   - heap profile `inuse_space` grows
   - same allocation sites keep retaining more objects

2. **Goroutine leak**
   - goroutine count grows
   - goroutine profile shows many stacks stuck in same place

3. **Runtime holding memory**
   - `HeapIdle` is high
   - `HeapReleased` is low
   - RSS is high but live heap is stable

4. **Expected cache growth**
   - memory grows until a cache bound is reached
   - growth stabilizes
   - object counts match cache expectations

5. **GC target growth**
   - depending on `GOGC`, heap may grow to a larger target before collecting
   - not necessarily a leak

### Practical method

Take heap profiles at two times:

```bash
curl -o heap1.prof http://localhost:6060/debug/pprof/heap
sleep 600
curl -o heap2.prof http://localhost:6060/debug/pprof/heap
```

Compare:

```bash
go tool pprof -base heap1.prof -http=:8080 heap2.prof
```

If the diff shows retained growth in specific allocation sites, that is strong evidence of a leak or unintended retention.

If live heap is stable but RSS is high, it may be runtime memory retention rather than a true leak.

---

## Practical production debugging workflow

This is the workflow I would defend in an interview.

## Step 1: Establish baseline

Collect:

- p50/p95/p99 latency
- request rate
- error rate
- CPU usage
- memory usage
- GC CPU
- goroutine count
- queue depth
- Kafka publish latency/failures

Without baseline, “better” is meaningless.

---

## Step 2: Reproduce under load

Use:

```bash
go test -run '^$' -bench=. -benchmem ./...
```

For concurrency:

```go
b.RunParallel(func(pb *testing.PB) {
    for pb.Next() {
        // workload
    }
})
```

Use `b.SetParallelism(n)` for 1, 4, 16, 64 concurrency cases.

---

## Step 3: Profile

Depending on symptom:

| Symptom | First profile |
|---|---|
| high CPU | CPU profile |
| high allocations/GC | heap alloc profile |
| memory growth | heap inuse profile |
| high latency, low CPU | block/mutex/goroutine profile |
| suspected leak | goroutine + heap diff |
| scheduler weirdness | execution trace |

---

## Step 4: Form a hypothesis

Examples:

- “p99 is caused by JSON allocation pressure.”
- “Throughput flattens because workers contend on one mutex.”
- “Memory grows because payload subslices retain full buffers.”
- “Goroutines leak because shutdown never cancels publishers.”

---

## Step 5: Fix and verify

After the change:

- rerun benchmarks
- rerun race detector
- compare profiles
- compare GC behavior
- compare latency percentiles
- verify under realistic concurrency

---

## Step 6: Guard in CI

Good candidates for CI:

```bash
go test -race ./...
go vet ./...
go test -run '^$' -bench=. -benchmem ./...
```

For performance-sensitive code, store benchmark results and compare against baseline with `benchstat`.

---

## Commands you should know cold

These are the core commands from your study plan, with context.

```bash
go test -race ./...
```

Run tests with the race detector.

```bash
go test -run '^$' -bench=. -benchmem ./...
```

Run benchmarks only, with allocation stats.

```bash
go tool pprof -http=:8080 cpu.prof
```

Open saved CPU profile.

```bash
go build -gcflags="-m" .
```

Inspect escape analysis.

Additional useful commands:

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile?seconds=30
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/allocs
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/goroutine
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/mutex
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/block
go tool trace trace.out
```

---

## Senior-level production mindset

In senior Go interviews, the interviewer wants to hear that you can connect internals to system behavior.

For example:

> “If our telemetry publisher allocates a new buffer per event, the CPU profile will show allocation and GC overhead. The heap alloc profile will show the encode path. We can reduce allocations with buffer reuse and preallocation, then verify with `-benchmem`, pprof, and p99 latency. Only after that would I consider adjusting `GOGC` or `GOMEMLIMIT`.”

That answer is strong because it links:

- code behavior
- memory management
- profiling evidence
- production metrics
- engineering trade-offs

That is the level expected for a senior role.

---

---

## Revision checkpoint

Try answering from memory *before* opening the relevant section. Mark the boxes only after you can explain or demonstrate the result.

- [ ] Design a benchmark that excludes setup and resists compiler elimination.
- [ ] Explain what a clean `go test -race` run does and does not prove.
- [ ] Choose CPU, heap, allocs, goroutine, mutex, block, or trace for a symptom.
- [ ] Investigate a p99 latency regression with an explicit hypothesis and baseline.
- [ ] Prove an optimization using repeated benchmarks and production SLO metrics.

**After review:** record your confidence, uncertainties, and evidence in [NOTES.md](NOTES.md) and add one entry to the [revision log](../REVISION-LOG.md).
