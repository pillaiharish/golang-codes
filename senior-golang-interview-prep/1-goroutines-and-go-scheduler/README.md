[← Senior Go interview preparation index](../README.md) · [Personal notes](NOTES.md) · [Revision log](../REVISION-LOG.md)

> Study notes adapted from the supplied transcript. Technical examples are preserved for review; code blocks are illustrative unless separately validated.

# 1 — Goroutines and the Go Scheduler

## 1.1 What a goroutine really is

A goroutine is **not** simply an OS thread.

A goroutine is a **user-space concurrent function** managed by the Go runtime.

Key differences from OS threads:

| Aspect | OS Thread | Goroutine |
|---|---|---|
| Created by | Kernel | Go runtime |
| Stack size | Usually fixed and relatively large | Small, dynamically growable |
| Creation cost | Relatively expensive | Cheap |
| Scheduling | Kernel scheduler | Go runtime scheduler |
| Context switch | Expensive | Much cheaper |
| Multiplexing | Usually 1:1 | Many goroutines onto fewer OS threads |

A goroutine starts with a small stack, commonly a few KB, and grows/shrinks as needed. This is why Go can run tens or hundreds of thousands of goroutines without needing that many OS threads.

---

## 1.2 GMP: the core scheduling model

Go’s scheduler is commonly described using the **GMP model**:

- **G** = Goroutine
- **M** = Machine, i.e. an OS thread
- **P** = Processor, i.e. a logical processor required to execute Go code

### G — Goroutine

A goroutine contains:

- its stack
- instruction pointer / state
- scheduling metadata
- current status, such as runnable, running, waiting, or in a syscall

When you write:

```go
go doWork()
```

the runtime creates a new **G**.

That G is not immediately running on its own OS thread. It is placed into a run queue and scheduled onto a P/M pair.

---

### M — Machine

An **M** represents an OS thread.

It is the actual execution context that runs machine instructions.

But an M cannot execute Go code unless it has a **P**.

---

### P — Processor

A **P** is a logical processor.

It holds:

- a local run queue of runnable goroutines
- scheduling state
- caches needed for efficient execution

The number of Ps is controlled by:

```go
GOMAXPROCS
```

By default:

```go
GOMAXPROCS = number of logical CPU cores
```

So if you have 8 cores, Go usually has 8 Ps by default.

---

## 1.3 How scheduling works conceptually

At a high level:

1. A goroutine becomes runnable.
2. It is placed on a P’s local run queue, or the global run queue.
3. An M must acquire a P to run that goroutine.
4. The M executes the G.
5. If the G blocks, the runtime tries to keep the P busy by:
   - parking the G if it is waiting on something,
   - finding another runnable G,
   - or associating the P with another M if necessary.

This is how Go achieves **M:N scheduling**:

> many goroutines are multiplexed onto a smaller number of OS threads.

---

## 1.4 Goroutine states

A simplified view of goroutine states:

### 1. Runnable

The goroutine is ready to run and waiting for CPU time.

### 2. Running

The goroutine is currently executing on an M/P.

### 3. Waiting

The goroutine is blocked on something such as:

- channel send/receive
- mutex lock
- network I/O
- timer
- select

### 4. Syscall

The goroutine is inside a system call.

This distinction matters because **network I/O** and **blocking OS syscalls** behave differently in the runtime.

---

## 1.5 What happens when a goroutine blocks on network I/O?

This is a very common senior interview topic.

Go’s network I/O is typically integrated with the OS async notification mechanism:

- epoll on Linux
- kqueue on macOS/BSD
- I/O completion ports on Windows

When a goroutine performs network I/O:

1. The socket is usually set to non-blocking mode.
2. If the operation cannot complete immediately, the goroutine is parked.
3. The underlying file descriptor is registered with the runtime netpoller.
4. The P is freed to run other goroutines.
5. When the socket becomes ready, the netpoller makes the goroutine runnable again.

This is extremely important.

It means:

> A goroutine waiting on network I/O usually does **not** pin an OS thread forever.

This is one reason Go works well for high-concurrency network services.

### Example

```go
go func() {
    resp, err := http.Get("https://example.com")
    // while waiting for the response, this goroutine can be parked
    // and other goroutines can continue running
    _ = resp
    _ = err
}()
```

In a production telemetry service, this matters because you may have thousands of outbound HTTP/Kafka/network operations waiting at the same time, without needing thousands of blocked OS threads.

---

## 1.6 What happens during a blocking system call?

This is different from async network I/O.

If a goroutine makes a **truly blocking system call**, for example:

- some file I/O
- CGO calls
- certain OS operations

then the M may become blocked inside the kernel.

The runtime tries to keep the P productive:

- the P can be detached from the blocked M
- another M can pick up that P
- other runnable goroutines can continue executing

So Go has mechanisms to avoid stalling the whole scheduler just because one goroutine is stuck in a blocking syscall.

But if many goroutines make long blocking syscalls, the runtime may need to create more Ms, and that can increase resource usage.

---

## 1.7 Preemption

Older versions of Go relied heavily on cooperative scheduling. That created problems if a goroutine ran a tight CPU loop without function calls.

Modern Go supports **asynchronous preemption**.

In practice:

- function calls still create safe preemption points
- the runtime can also preempt long-running goroutines using signals

This helps prevent a single CPU-bound goroutine from starving others indefinitely.

### Important nuance

Preemption improves fairness, but it does **not** guarantee deterministic ordering.

Go does not promise:

- which runnable goroutine runs next
- when a goroutine will run
- that goroutines finish in creation order

If your logic depends on ordering, you must enforce it yourself, usually with channels or other synchronization.

---

## 1.8 Work stealing

The scheduler uses work stealing to balance load.

If a P has no local runnable goroutines, it may:

1. check the global run queue
2. check the netpoller
3. steal runnable goroutines from another P’s local queue

This improves utilization across cores.

---

## 1.9 Goroutine stacks

Goroutine stacks are not fixed-size like many OS threads.

They:

- start small
- grow when needed
- may shrink when possible

Historically Go moved from segmented stacks to contiguous stacks. Modern Go uses contiguous stacks and copies/adjusts them when growth is needed.

This design is part of why goroutines are cheap.

---

## 1.10 Goroutine leaks

A goroutine leak happens when a goroutine is created but never terminates, often because it is blocked forever.

Common causes:

### 1. Sending to a channel with no receiver

```go
func leak() {
    ch := make(chan int)
    go func() {
        ch <- 1 // blocks forever if no receiver
    }()
}
```

### 2. Receiving from a channel that never gets a send

```go
func leak() {
    ch := make(chan int)
    go func() {
        <-ch // blocks forever
    }()
}
```

### 3. Forgetting to cancel a context

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // if omitted, worker may run forever

go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // work
        }
    }
}()
```

### 4. HTTP calls without timeouts

```go
resp, err := http.Get(url) // may hang too long without timeout
```

### 5. Deadlocked select or nil channel receive

```go
var ch chan int
<-ch // blocks forever
```

### How to detect leaks in production

Useful signals:

- `runtime.NumGoroutine()` increasing over time
- `/debug/pprof/goroutine`
- goroutine profiles showing many goroutines stuck in the same stack
- tracing
- metrics/alerting on goroutine count

A senior answer should not just say “use pprof.” It should also say **what pattern you expect to see** in the profile.

For example:

- many goroutines stuck in `chan receive`
- many stuck in `select`
- many stuck in `http.(*Client).do`
- many stuck in `sync.(*Mutex).Lock`

That tells you the leak or bottleneck category.

---

## Interview questions and senior-level answers

## Question 1: Are goroutines lightweight threads? How do they differ?

### Good answer

Goroutines are often called lightweight threads, but that phrase can be misleading.

They are **not OS threads**. They are user-space execution units scheduled by the Go runtime.

Key differences:

- goroutines have much smaller, dynamic stacks
- goroutine creation is much cheaper
- the Go runtime schedules many goroutines onto a smaller pool of OS threads
- blocking behavior is handled differently by the runtime
- context switching between goroutines is usually much cheaper than kernel thread switching

### Senior-level extra

In production, this means a Go service can handle many concurrent connections or background tasks without a 1:1 thread explosion. But it also means leaks can quietly grow the heap and goroutine count, so lifecycle management matters.

---

## Question 2: If you create 100,000 goroutines, do you create 100,000 OS threads?

### Good answer

No.

Go multiplexes goroutines onto OS threads using the runtime scheduler.

If you create 100,000 goroutines, you typically do **not** create 100,000 OS threads.

The number of OS threads depends on:

- `GOMAXPROCS`
- how many goroutines are runnable
- how many are blocked
- whether they are blocked on network I/O or real blocking syscalls
- runtime decisions

For CPU-bound work, the number of actively running goroutines at any instant is usually bounded by the number of Ps.

For I/O-bound work, many goroutines can be waiting simultaneously, and Go can still run others with a limited number of threads.

### Senior-level extra

If all 100,000 goroutines are mostly waiting on I/O, that can be fine.

If they are all CPU-bound, adding more goroutines than useful parallelism can increase overhead without improving throughput.

If they are all blocked on real blocking syscalls, the runtime may spawn more OS threads, increasing resource pressure.

---

## Question 3: How does `GOMAXPROCS=4` affect CPU-bound and I/O-bound workloads?

### Good answer

`GOMAXPROCS=4` means Go uses 4 logical processors by default.

For **CPU-bound work**:

- at most 4 goroutines can be executing Go code in parallel at any instant
- creating more goroutines than that does not create more CPU parallelism
- it may even increase scheduling overhead

For **I/O-bound work**:

- many goroutines can still be in flight
- while some goroutines wait on I/O, others can run
- `GOMAXPROCS=4` does not limit total concurrency to 4, only the number of Ps available for executing Go code

### Senior-level extra

In a telemetry ingestion service:

- request handlers and queue workers may be I/O-bound
- JSON encoding/compression may be CPU-bound
- setting worker counts without understanding the bottleneck can hurt performance

A common senior mistake is to assume “more goroutines = faster.” That is only true if the bottleneck is not CPU-bound parallelism.

---

## Question 4: What happens when a goroutine makes a blocking system call?

### Good answer

It depends on the syscall.

For network I/O, Go typically uses non-blocking sockets and the runtime netpoller. The goroutine parks, and the P can run other goroutines.

For a truly blocking system call, the M may block. The runtime can detach the P from that M and associate it with another M so other goroutines can continue running.

If many goroutines block in syscalls, the runtime may create more OS threads.

### Senior-level extra

This distinction is important when integrating with:

- file systems
- CGO
- some databases/drivers
- external processes

If a dependency forces blocking syscalls, Go’s usual efficient I/O multiplexing may not fully apply.

---

## Question 5: Can a goroutine be starved? Is execution order guaranteed?

### Good answer

Yes, a goroutine can be starved in practice, although the runtime tries to be fair.

Execution order is **not guaranteed**.

Go does not guarantee:

- which goroutine runs first
- that goroutines run in creation order
- that all runnable goroutines get equal CPU time

Starvation-like behavior can happen due to:

- lock contention
- unfair channel patterns
- CPU-bound loops
- poor workload design
- scheduler pressure

Modern preemption reduces the risk of a single goroutine monopolizing a CPU forever, but it does not make scheduling deterministic.

### Senior-level extra

If ordering or fairness matters, do not rely on the scheduler.

Use explicit design:

- channels for sequencing
- priority queues if needed
- worker pools with bounded queues
- context cancellation for lifecycle control

---

## Question 6: How would you identify whether your Go service has a goroutine leak?

### Good answer

I would look at both metrics and profiles.

First, monitor:

- `runtime.NumGoroutine()`
- goroutine count over time
- memory growth
- request latency and queue depth

Then inspect:

- `/debug/pprof/goroutine?debug=1`
- `go tool pprof` goroutine profile
- traces

I would look for many goroutines stuck in the same function or state, such as:

- channel receive
- select
- mutex lock
- HTTP read
- waitgroup wait

That usually points to the leak source.

### Senior-level extra

In production, I would also correlate the leak with:

- traffic patterns
- request cancellation behavior
- downstream failures
- queue backpressure
- deployment changes

A goroutine leak is usually a symptom of a lifecycle bug, not a random runtime issue.

---

## Production connection

Imagine a telemetry service:

- HTTP handlers receive events
- each event is published to an internal queue
- worker goroutines send to Kafka
- downstream systems may become slow

If you do not understand scheduling and lifecycle management, several things can go wrong:

- handler goroutines pile up waiting on a full queue
- publisher goroutines block on slow Kafka sends
- context cancellation is not propagated
- goroutine count rises continuously
- heap grows
- p99 latency worsens

A senior engineer reasons about:

- how many workers should run
- whether the queue is bounded
- whether producers should block or drop
- how cancellation flows through the system
- how to observe queue depth and drop rate

That connects Module 1 directly to Module 2.

---

---

## Revision checkpoint

Try answering from memory *before* opening the relevant section. Mark the boxes only after you can explain or demonstrate the result.

- [ ] Explain G/M/P and why `GOMAXPROCS` limits Go execution parallelism, not total I/O concurrency.
- [ ] Walk through a goroutine parking on network I/O versus blocking inside a syscall.
- [ ] Explain why preemption does not imply deterministic execution order.
- [ ] Describe the profile evidence that would confirm a goroutine leak.
- [ ] Defend the worker count for a CPU-bound versus an I/O-bound workload.

**After review:** record your confidence, uncertainties, and evidence in [NOTES.md](NOTES.md) and add one entry to the [revision log](../REVISION-LOG.md).
