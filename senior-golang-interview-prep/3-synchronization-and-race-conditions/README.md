[← Senior Go interview preparation index](../README.md) · [Personal notes](NOTES.md) · [Revision log](../REVISION-LOG.md)

**Focused doubts:** [4. Increment statements and for-loop traps](../doubts/04-increments-and-for-loops.md) · [5. Bounded buffers, blocking, and producer-consumer policy](../doubts/05-buffers-blocking-and-backpressure.md)

> Study notes adapted from the supplied transcript. Technical examples are preserved for review; code blocks are illustrative unless separately validated.

# 3 — Synchronization and Race Conditions

This module is where most candidates reveal whether they've actually debugged concurrency in production or only read about it. The trap is treating "synchronization" as a syntax topic. It isn't. It's about **two distinct problems** that people constantly conflate:

1. **Atomicity** — can an operation be torn or interleaved?
2. **Visibility and ordering** — will another goroutine ever *see* my write, and in what order?

A senior answer names both. Most answers only name the first.

---

## 3.1 What a data race actually is

Go's memory model defines a **data race** precisely:

> Two goroutines access the same memory location concurrently, **at least one access is a write**, and **neither access happens-before the other**.

That last clause is the one everyone skips. It's not "they touched it at the same time." It's "there is no ordering edge between them."

### The canonical broken counter

```go
var counter int
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        counter++ // NOT atomic: load, add, store
    }()
}
wg.Wait()
fmt.Println(counter) // often < 1000
```

`counter++` is three machine operations. Two goroutines can both load `41`, both add one, both store `42`. One increment is **lost**.

Typical observed output: `873`, `941`, `998` — different every run. That variance is itself a signal: **non-determinism in a "simple" counter is never a compiler quirk, it's a missing ordering edge.**

And with `-race`:

```text
WARNING: DATA RACE
Write at 0x000000000000 by goroutine 8:
  main.main.func1()
Previous write at 0x000000000000 by goroutine 7:
  main.main.func1()
```

### The critical caveat — say this in the interview

> The race detector is **dynamic**. It only reports races on interleavings your program actually produced during that run. A clean `-race` run is necessary evidence, **not proof of absence**.

This single sentence separates a senior candidate from a mid-level one, because it explains why you still need design discipline, invariants, and tests under contention.

---

## 3.2 `sync.Mutex`

```go
type Counter struct {
    mu sync.Mutex
    n  int64
}

func (c *Counter) Inc() {
    c.mu.Lock()
    c.n++
    c.mu.Unlock()
}

func (c *Counter) Load() int64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.n
}
```

### The traps, in the order they actually bite you in production

**1. Go mutexes are not reentrant.**

```go
func (c *Counter) IncTwice() {
    c.mu.Lock()
    c.Inc() // deadlock: same goroutine, second Lock
    c.mu.Unlock()
}
```

This is the single most common self-inflicted deadlock. The rule: **never call a method that takes the same lock while you hold it.** Keep locked helpers unexported and named to make the contract obvious (`incLocked`).

**2. Never copy a mutex.**

```go
c := Counter{}
c2 := c // go vet: copies lock value
```

A copied mutex has independent, meaningless state. `go vet`'s `copylocks` catches most cases — which is why you run `go vet` in CI, not just `-race`.

**3. Never hold a lock across I/O.**

```go
c.mu.Lock()
resp, _ := http.Get(url) // 2 seconds of everyone else blocked
c.mu.Unlock()
```

This converts a slow downstream into a **global stall**. It's the #1 cause of "our p99 is 4 seconds but CPU is idle." Move the I/O outside the critical section, take the lock only to publish the result.

**4. Never hold a lock while sending on a channel.**

```go
c.mu.Lock()
c.ch <- event // blocks if queue is full
c.mu.Unlock()
```

If the consumer needs `c.mu` to drain the queue → **deadlock**. This pattern is shockingly common in Go pipeline code because each half looks obviously correct in isolation.

**5. Lock ordering.**

If you ever need two locks, define a **global acquisition order** and enforce it everywhere. `lockA → lockB` in one place and `lockB → lockA` in another is a deadlock that reproduces once a week under specific load.

### On `defer` for unlock

`defer c.mu.Unlock()` is safer on multi-return/error paths and costs a small but real amount in hot loops. My practical rule: **`defer` by default; explicit `Unlock` only in a measured hot path where you've proven the difference matters.** Never optimize this before you've profiled.

---

## 3.3 `sync.RWMutex` — and why it's usually the wrong default

```go
type Cache struct {
    mu sync.RWMutex
    m  map[string]string
}

func (c *Cache) Get(k string) (string, bool) {
    c.mu.RLock()
    v, ok := c.m[k]
    c.mu.RUnlock()
    return v, ok
}

func (c *Cache) Set(k, v string) {
    c.mu.Lock()
    c.m[k] = v
    c.mu.Unlock()
}
```

### The three things people don't know

**1. `RWMutex` is writer-preferring.**

When a writer calls `Lock()` and must wait, **subsequent `RLock()` calls also block.** This exists to prevent writer starvation — and it's also the mechanism by which one slow writer turns into a full-system stall. A read-heavy system with an unlucky long-held write lock performs *worse* than a plain `Mutex`.

**2. `RLock` is not recursive, and upgrading deadlocks.**

```go
func (c *Cache) GetOrLoad(k string) string {
    c.mu.RLock()
    v, ok := c.m[k]
    if !ok {
        c.mu.RUnlock()
        c.Set(k, fetch(k)) // fine
        // but if you forgot RUnlock and called Set -> deadlock
    }
    c.mu.RUnlock()
    return v
}
```

There is **no lock upgrade** in Go. `RLock` → `Lock` on the same goroutine is a guaranteed deadlock.

**3. The read path is not free.**

`RLock` performs atomic updates on shared reader counters. With many cores doing many short reads, those cache lines bounce between cores — and a plain `Mutex`, whose uncontended path is a single CAS, can be **faster**.

> This is the direct answer to "is RWMutex always faster for many readers?" — **No.** I'll come back to it formally below.

---

## 3.4 `sync.WaitGroup`

```go
var wg sync.WaitGroup
for i := 0; i < n; i++ {
    wg.Add(1)              // in the PARENT, before go
    go func(i int) {
        defer wg.Done()    // immediately, so panics/returns still count
        work(i)
    }(i)
}
wg.Wait()
```

### Failure modes worth naming out loud

| Mistake | Consequence |
|---|---|
| `wg.Add(1)` **inside** the goroutine | `Wait()` may return before any goroutine registers — work still in flight |
| Missing `Done` on an early return / panic | `Wait()` hangs forever |
| Copying the `WaitGroup` | broken semantics; `go vet` copylocks |
| Counter goes negative | panic: `sync: negative WaitGroup counter` |
| Reusing a `WaitGroup` while `Wait` is returning | undefined; misuse |

The `Add`-inside-the-goroutine bug is the one I'd expect a senior engineer to identify instantly, because it produces a specific and maddening symptom class: **tests that pass 95% of the time**, and shutdown code that returns before flushing.

*Side note worth having ready:* since **Go 1.22**, `for` loop variables are per-iteration, so `go func() { use(i) }()` is now safe without shadowing. If an interviewer asks about the classic capture bug, mention the version change — it signals you're current.

---

## 3.5 `sync/atomic`

Modern typed API (Go 1.19+):

```go
var n atomic.Int64

n.Add(1)
n.Load()
n.CompareAndSwap(old, new)
```

Older `atomic.AddInt64(&x, 1)` still works; prefer the typed forms — they prevent the classic mistake of mixing atomic and non-atomic access to the same variable, which the compiler can't catch but the race detector can.

### What atomics give you

- **Atomicity** of a single memory operation.
- **Ordering**: Go's memory model (formalized around 1.19) treats atomic operations as synchronization edges and gives them sequential-consistency semantics. An atomic store that a later atomic load observes establishes happens-before.

### What atomics do **not** give you

**Multi-variable invariants.** This is the important one.

```go
// Intention: never exceed limit
if c.total.Load() < limit {
    c.total.Add(1) // 500 goroutines pass the check simultaneously
}
```

Every operation is atomic. The **program is still wrong.** This is a *race condition* — check-then-act — with zero data races. The fix is a `CompareAndSwap` loop, or just a mutex:

```go
for {
    cur := c.total.Load()
    if cur >= limit {
        return false
    }
    if c.total.CompareAndSwap(cur, cur+1) {
        return true
    }
}
```

And note the honest trade-off: CAS loops can **livelock** under extreme contention, because they retry. A mutex *queues*. Under brutal contention, a mutex's fairness can beat a CAS loop's optimism. That's a genuinely senior observation.

### The pattern I actually reach for on read-mostly state

`atomic.Pointer[T]` — copy-on-write. Readers never lock; the writer builds a whole new snapshot and swaps one pointer.

```go
type Config struct {
    Endpoints []string
    Timeout   time.Duration
}

var cfg atomic.Pointer[Config]

func Load() *Config { return cfg.Load() }        // lock-free, zero contention
func Update(c *Config) { cfg.Store(c) }          // one writer
```

For configuration, feature flags, route tables, and lookup tables in a telemetry service, this beats both `Mutex` and `RWMutex` — and knowing it is a strong signal.

---

## 3.6 `sync.Once`

```go
var once sync.Once
var producer *KafkaProducer

func Get() *KafkaProducer {
    once.Do(func() { producer = newProducer() })
    return producer
}
```

### The traps

**1. Other callers *block* until `Do` returns.** If your initializer hangs on a Kafka broker that isn't up, **every caller hangs**. That's a startup deadlock with a very confusing symptom.

**2. There is no error return, and it will not retry.** If `newProducer()` fails, `Once` is consumed forever — you've permanently cached a broken nil. The fix:

```go
var (
    once     sync.Once
    producer *KafkaProducer
    perr     error
)

func Get() (*KafkaProducer, error) {
    once.Do(func() { producer, perr = newProducer() })
    return producer, perr
}
```

Now the failure is *observable*, even though it still won't retry. If you need retry-on-failure, `Once` is the wrong primitive — say so.

**3. Re-entrancy deadlocks.** Calling `once.Do` from inside its own function is a self-deadlock, same family as the non-reentrant mutex.

---

## 3.7 `sync.Pool` — and what it is *not*

`sync.Pool` is a **per-P scratch-object reuse mechanism to reduce allocation and GC pressure.** It is **not a cache.**

```go
var bufPool = sync.Pool{
    New: func() any { return new(bytes.Buffer) },
}

func encode(e TelemetryEvent) []byte {
    b := bufPool.Get().(*bytes.Buffer)
    b.Reset()
    defer func() { bufPool.Put(b) }()

    json.NewEncoder(b).Encode(e)
    return append([]byte(nil), b.Bytes()...) // copy out before Put
}
```

### Why it scales

Each P keeps a **private slot**, so `Get`/`Put` on the hot path touch no shared state. That's the whole design: avoid the very cache-line contention we've been discussing all module.

### The traps

- **Objects can be discarded at any GC.** Since Go 1.13 the pool is cleared on *every* GC. You get no reuse guarantee, no warm-pool guarantee, no lifetime guarantee. Never build correctness on it.
- **Use-after-`Put` is a data race waiting to happen.** You `Put` the buffer, still hold the pointer, keep writing — another goroutine `Get`s it. Often the race detector *won't* catch this cleanly, because both accesses may be "valid" in isolation. This is exactly the class of bug that motivates the later question: *"can a program be race-free and still be broken?"*
- **Pooled objects pin memory.** A pool of 1 MiB buffers × GOMAXPROCS × several pools is retained heap that looks like a leak but isn't.
- **Right use:** byte buffers for encoding/serialization, request-scoped scratch structs, `[]byte` for compression. **Wrong use:** caching expensive-to-build clients, connections, or anything with a real lifecycle.

---

## 3.8 Race condition vs data race vs deadlock — three different diseases

Candidates lose marks by using these interchangeably. Keep them separate:

| Term | Definition | Detected by `-race`? |
|---|---|---|
| **Data race** | Concurrent access, one write, no happens-before | Yes, if the interleaving occurs |
| **Race condition** | Correctness depends on event timing/interleaving | **Often no** |
| **Deadlock** | No progress; everyone waits | **No** — it just hangs |

And deadlocks have Go-specific flavors worth naming, because naming them proves experience:

- lock-order inversion
- mutex re-entrancy
- send on a channel with no receiver, ever
- `WaitGroup` never `Done`
- **holding a mutex while sending on a full channel** ← the signature Go deadlock
- `select` with only nil-channel cases

---

## 3.9 Happens-before and memory visibility

This is the conceptual core. **You synchronize not only to make operations atomic, but to create ordering edges that make writes visible.** Compilers and CPUs reorder. Without an edge, another goroutine may never observe your write at all.

### The flag bug

```go
var done bool
var data string

go func() {
    data = "ready"
    done = true
}()

for !done {
}
fmt.Println(data) // may print ""
```

No data race is *needed* for this to break — the write to `data` can be reordered after the write to `done`, or `done` may never become visible in the spinning goroutine's view. The compiler is also free to hoist `done` out of the loop entirely.

Correct versions: close a channel, use a `Mutex`, use `atomic.Bool`, or use a `WaitGroup`.

### The edges you should be able to recite

- `go f()` **happens-before** the execution of `f`
- a channel **send** happens-before the corresponding **receive** completes (for capacity `C`, the *k*th send hb the *k*th receive; the *k*th receive hb the *k*+*C*th send's completion)
- **`close(ch)`** happens-before a receive that returns the zero value because the channel is closed
- **`mu.Unlock()`** happens-before a subsequent **`mu.Lock()`** on the same mutex
- **`wg.Done()`** happens-before **`wg.Wait()`** returning
- **`once.Do(f)`** returning happens-before subsequent `Do` returns
- an **atomic store** happens-before an **atomic load** that observes it

### The corollary people get wrong

> **Write before you send.** A channel send publishes the state that existed *before* the send. Data mutated *after* the send has no ordering relationship with the receiver.

```go
go func() {
    result = compute() // must be before
    ch <- result       // publishes it
}()
```

---

## 3.10 Channels vs mutexes — the defensible answer

Not the slogan. The slogan is "share memory by communicating," and it's a good instinct and a bad universal.

**Use a mutex when** the state is genuinely *shared*: counters, caches, connection pools, maps, invariants spanning multiple fields. It's simpler, cheaper uncontended (one CAS), and much easier to audit for correctness.

**Use a channel when** you're *transferring ownership*: distributing work to workers, building pipeline stages, signaling cancellation, expressing lifecycle and shutdown.

My one-line rule, which I'd defend in an interview:

> **If the data has one owner, move it with a channel. If the data is genuinely shared, guard it with a mutex.**

On performance, be blunt: an uncontended mutex is roughly a CAS. A channel send/receive involves queueing and potentially waking a parked goroutine. **A channel-per-operation design for a hot counter is almost always a pessimization** — which is exactly what Module 3's coding challenge will prove numerically.

And on `sync.Map`: it's for two specific shapes — **read-mostly with a mostly-stable key set**, or **disjoint key sets per goroutine**. Under churning or write-heavy keys it's frequently *slower* than `map + RWMutex`, because of the read-only-map/dirty-map promotion path (dirty gets promoted to read-only after enough misses). Reaching for `sync.Map` by default is a mid-level tell.

---

## 3.11 Measuring and fixing lock contention

You cannot fix contention you haven't measured. Enable the profiles:

```go
runtime.SetMutexProfileFraction(5)  // sample lock wait
runtime.SetBlockProfileRate(1000)   // sample blocking: mutex, chan, etc.
```

```bash
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/mutex
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/block
```

The mutex profile reports **wait time attributed to acquisition, and where the lock was held.** That second part matters: the metric that kills you is **hold time × frequency**, not lock count.

Fixes, roughly in order of value:

1. **Shrink the critical section** — move I/O, allocation, and encoding outside the lock. This fixes more real incidents than everything else combined.
2. **Shard** — `N` independent mutexes indexed by `hash(key) % N`. Cuts contention by roughly `N`. Ideal for per-key telemetry counters.
3. **Replace with atomics** for single-value counters and flags.
4. **Copy-on-write with `atomic.Pointer`** for read-mostly config.
5. **Batch** — accumulate locally, flush under the lock occasionally.

---

## 3.12 Safe maps

Go maps are **not** safe for concurrent writes, and concurrent read-during-write is a data race that can **crash**, not merely corrupt.

| Option | Use when | Avoid when |
|---|---|---|
| `map` + `Mutex` | default; short critical sections | long-held sections |
| `map` + `RWMutex` | long read sections, rare writes | short reads (overhead + writer stall) |
| `sync.Map` | read-mostly, stable keys; disjoint per-goroutine keys | churning/write-heavy keys |
| Sharded map + `Mutex` | high write contention, key-addressable | small maps (shard overhead dominates) |

For a telemetry service aggregating per-series counters under load, **sharding by key hash** is usually the right answer, and saying so — with the reason — is a strong senior response.

---

## Interview questions and answers

## Q1: Mutex vs RWMutex — is RWMutex always faster for many readers?

**No.** Three reasons:

1. **The read path isn't free.** `RLock` performs atomic updates on shared reader counters. Many cores doing short reads means those cache lines bounce — and a plain `Mutex`, uncontended, is a single CAS. For a map lookup or a short field read, `Mutex` frequently wins.
2. **Writer preference creates a global stall.** A pending `Lock()` blocks *new* `RLock()` calls. One slow writer in a read-heavy system is worse than a plain mutex would have been.
3. **It only pays off under specific conditions:** read critical sections long relative to lock overhead, writes rare, and writes short.

And the better answer continues: for read-mostly **immutable** snapshots, the right answer is often **neither** — `atomic.Pointer[T]` copy-on-write gives lock-free reads and never blocks the writer at all.

---

## Q2: Can two goroutines safely increment the same integer without synchronization?

**No.** `i++` is load → add → store. Interleaving loses updates.

But go deeper than "use a mutex":

- It's a **data race**, which means **undefined behavior**, not merely "a slightly wrong number." The compiler may reorder, hoist, or tear. The wrong value is the *benign* symptom.
- Even made atomic, if the increment is part of a **larger invariant**, you can still be wrong — the check-then-act case above.

Correct options: `atomic.Int64` for a plain counter, `Mutex` when the increment is one part of a multi-field update.

---

## Q3: Difference between an atomic counter and a mutex-protected counter?

| | Atomic | Mutex |
|---|---|---|
| Scope | one location, one operation | arbitrary critical section, many variables |
| Uncontended cost | one locked instruction | one CAS |
| Contended behavior | cache-line transfer per op; CAS loops retry | queues and parks (syscall, scheduler latency) |
| Expresses invariants | no, without CAS loops | yes, naturally |
| Failure mode | can't express check-then-act; CAS livelock | deadlock; held too long |
| Memory ordering | yes (sync edge) | yes (`Unlock` hb `Lock`) |

**Decision rule:** atomic for counters, gauges, and flags. Mutex the moment the invariant spans more than one word, or the critical section does anything non-trivial. Reaching for atomics to "make it fast" without measuring is how you end up with a subtly wrong program that benchmarks beautifully.

---

## Q4: What happens if you call `WaitGroup.Add` too late?

`Wait()` can return **before the goroutine has registered**, so your program proceeds with work still in flight.

The dangerous form:

```go
go func() {
    wg.Add(1)      // too late
    defer wg.Done()
    work()
}()
wg.Wait()          // may return immediately
```

Symptoms, and these are the ones that make it memorable: **tests that pass 95% of the time**, shutdown returning before a flush completes, and data races on the result variable that appear only under CI load. Related misuse: calling `Add` after the counter has reached zero while `Wait` is returning is undefined, and driving the counter negative panics with `sync: negative WaitGroup counter`.

**Rule: `Add` synchronously in the parent, before `go`. `defer Done()` as the first statement inside.**

---

## Q5: Can a program be free of data races but still contain concurrency bugs?

**Yes — and this is the question that separates levels.**

The race detector checks one narrow, specific property. It does not check your design.

Concrete cases:

- **Check-then-act with perfect atomics.** `if v.Load() < limit { v.Add(1) }` — every operation atomic, no data race, program still exceeds the limit.
- **Non-deterministic `select`.** Two ready cases, pseudo-random choice. Flaky ordering, no race reported.
- **Deadlock or livelock.** No data race at all — it simply never finishes.
- **Goroutine leak.** Entirely race-free; the process just exhausts memory over a week.
- **Broken composed invariant.** Two individually-correct atomic counters whose *sum* is supposed to be conserved.
- **Use-after-`Put` on a `sync.Pool` object**, where each access looks valid in isolation.
- **Wrong channel owner closing** — panics at runtime, no race involved.

And the meta-point: **`-race` is dynamic.** It only observes interleavings your tests produced. Clean output means "not found in this run," never "absent."

---

## Q6: How would you implement a thread-safe cache?

Design, not syntax:

```go
type entry struct {
    once   sync.Once
    val    Value
    err    error
    expiry time.Time
}

type Cache struct {
    mu    sync.RWMutex
    m     map[string]*entry
    ttl   time.Duration
    fetch func(context.Context, string) (Value, error)
}

func (c *Cache) Get(ctx context.Context, key string) (Value, error) {
    c.mu.RLock()
    e, ok := c.m[key]
    c.mu.RUnlock()

    if ok && time.Now().Before(e.expiry) {
        e.once.Do(func() { e.val, e.err = c.fetch(ctx, key) })
        return e.val, e.err
    }

    ne := &entry{expiry: time.Now().Add(c.ttl)}
    c.mu.Lock()
    if existing, ok := c.m[key]; ok && time.Now().Before(existing.expiry) {
        c.mu.Unlock()
        ne = existing
    } else {
        c.m[key] = ne
        c.mu.Unlock()
    }
    ne.once.Do(func() { ne.val, ne.err = c.fetch(ctx, key) })
    return ne.val, ne.err
}
```

The decisions that matter, and that I'd articulate:

1. **Never hold the write lock across the fetch.** That's the mistake that turns a cache into a global stall. Use a per-entry `sync.Once` so concurrent misses on the same key produce **one** fetch — single-flight, for free.
2. **Don't return pointers into cached state** unless values are immutable. Returning `*Value` lets a caller mutate the cache without holding any lock — a data race the detector may never surface, because it's a logic bug.
3. **TTL:** lazy expiry on read plus a janitor goroutine bound to a root `context` so shutdown is clean.
4. **Bound the size** with eviction, or it's a memory leak with extra steps.
5. **Instrument it:** hit/miss ratio, in-flight count, eviction rate, fetch latency. A cache you can't measure is a liability.
6. **Not `sync.Map`** by default — the key set here churns.

---

## Coding challenge: three counters, 1,000 goroutines

## The correctness test first

```go
func TestCounter_1000Goroutines(t *testing.T) {
    for _, tc := range []struct {
        name string
        c    Counter
    }{
        {"mutex", &MutexCounter{}},
        {"atomic", &AtomicCounter{}},
        {"channel", NewChannelCounter()},
    } {
        t.Run(tc.name, func(t *testing.T) {
            var wg sync.WaitGroup
            for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func() {
                    defer wg.Done()
                    tc.c.Inc()
                }()
            }
            wg.Wait()
            if got := tc.c.Load(); got != 1000 {
                t.Fatalf("got %d, want 1000", got)
            }
        })
    }
}
```

And the broken baseline, for contrast:

```go
type BrokenCounter struct{ n int64 }
func (c *BrokenCounter) Inc() { c.n++ } // data race
```

Run with `go test -race`: it fails the assertion **and** reports the race. Note that the assertion failure is *probabilistic* — on a single-core machine it might pass. That's worth saying: **a race that doesn't reproduce is not a race that doesn't exist.**

## The three implementations

```go
type Counter interface {
    Inc()
    Load() int64
}

// 1. Mutex
type MutexCounter struct {
    mu sync.Mutex
    n  int64
}

func (c *MutexCounter) Inc() {
    c.mu.Lock()
    c.n++
    c.mu.Unlock()
}
func (c *MutexCounter) Load() int64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.n
}

// 2. Atomic
type AtomicCounter struct{ n atomic.Int64 }

func (c *AtomicCounter) Inc()   { c.n.Add(1) }
func (c *AtomicCounter) Load() int64 { return c.n.Load() }

// 3. Owner goroutine — no shared memory at all
type ChannelCounter struct {
    inc  chan struct{}
    quit chan struct{}
}

func NewChannelCounter() *ChannelCounter {
    c := &ChannelCounter{
        inc:  make(chan struct{}),
        quit: make(chan struct{}),
    }
    go c.run()
    return c
}

func (c *ChannelCounter) run() {
    var n int64
    for {
        select {
        case <-c.inc:
            n++
        case <-c.quit:
            return
        }
    }
}

func (c *ChannelCounter) Inc() {
    select {
    case c.inc <- struct{}{}:
    case <-c.quit:
    }
}

func (c *ChannelCounter) Load() int64 {
    reply := make(chan int64, 1)
    // request/reply round trip just to observe the value
    ...
}
```

**Stop and notice something important about #3.** `Inc()` is **fire-and-forget**: it returns before the increment has been applied. And `Load()` needs a **full request/reply round trip** — two context switches — just to read a number. So it isn't even the same *contract* as the other two. That's a design observation, not a performance one, and it's the kind of thing that earns marks: **the channel-owner version is asynchronous by construction, and making it synchronous costs you a round trip per operation.**

## The benchmark

```go
var sink int64 // prevents dead-code elimination

func benchInc(b *testing.B, inc func()) {
    b.ReportAllocs()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            inc()
        }
    })
    sink = 1
}

func BenchmarkMutex(b *testing.B) {
    c := &MutexCounter{}
    benchInc(b, c.Inc)
}

func BenchmarkAtomic(b *testing.B) {
    c := &AtomicCounter{}
    benchInc(b, c.Inc)
}

func BenchmarkChannel(b *testing.B) {
    c := NewChannelCounter()
    defer close(c.quit)
    benchInc(b, c.Inc)
}
```

Two harness points worth stating, because getting these wrong invalidates everything:

- **Use `b.RunParallel`, not 1,000 hand-spawned goroutines.** Manual goroutine creation cost dominates the measurement and you'd be benchmarking the scheduler's spawn path, not the counter. (`b.RunParallel` is also the correct tool for measuring *contention*, which is the whole point.)
- **Sink the result.** Without `sink`, the compiler can eliminate the work entirely and you'll benchmark nothing at a suspiciously impressive 0.2 ns/op.

## Expected shape of results

Exact numbers are machine- and core-count-dependent; the *shape* and the *why* are what matter.

| Implementation | Typical high-contention ns/op | Allocs/op | Dominant cost |
|---|---|---|---|
| **Atomic** | fastest | 0 | single locked instruction; bounded by cache-line ownership transfer |
| **Mutex** | slower | 0 | cheap CAS uncontended; **parking + syscalls** when contended |
| **Channel owner** | order(s) of magnitude slower | 0 | rendezvous + scheduler wake-up per op; **all work on one goroutine** |

### Why they differ — the actual explanation

**Atomic.** One hardware read-modify-write on one cache line. Every core wants exclusive ownership of that line, so throughput is bounded by cache-line transfer latency, serialized. Minimal instruction count, no syscalls, no parking.

**Mutex.** Uncontended, `Lock` is essentially a CAS — nearly free. Contended, the runtime **parks goroutines**, which means futex syscalls, context switches, and scheduler latency. That's the gap. But note the flip side: a mutex **queues**, which under extreme contention is a fairness property a CAS loop doesn't have.

**Channel owner.** Every single increment costs a channel send, a wake-up of the owner, and a receive. You've paid scheduler participation for *every operation*. And critically — **adding cores does not help, because one goroutine does all the work.** It cannot scale with `GOMAXPROCS`. It's a hard serialization point by design.

### The observation that makes this a senior answer

> **At low concurrency, all three are within noise of each other.** The divergence appears as core count and goroutine count rise. So if someone tells you "atomics are faster than mutexes," the honest reply is: *for this workload, under this contention, on this machine — and here's the profile that shows it.*

And the design conclusion:

- **Single counter, hot path** → `atomic`. Ship it.
- **Invariant over multiple fields, or anything non-trivial in the section** → `Mutex`, and accept the cost, because correctness is worth more than 40 ns.
- **Channel ownership** → not a performance technique. It's a **correctness and composability** technique: explicit ownership, no shared mutable state, clean shutdown. You accept a large per-op cost to make the design easier to reason about and to make lifecycle obvious. Choosing it for a hot counter is a category error.

That last sentence is the answer I'd want from a senior Go engineer: *knowing that the elegant pattern is the wrong tool here.*

---

---

## Revision checkpoint

Try answering from memory *before* opening the relevant section. Mark the boxes only after you can explain or demonstrate the result.

- [ ] Distinguish a data race, a logical race condition, and a deadlock.
- [ ] Explain the Go happens-before edges created by channels, locks, and atomics.
- [ ] Choose Mutex, RWMutex, atomic, sync.Map, or a sharded map and justify the trade-off.
- [ ] Spot WaitGroup, sync.Once, sync.Pool, and mutex misuse.
- [ ] Run and explain the three-counter correctness and benchmark exercise.

**After review:** record your confidence, uncertainties, and evidence in [NOTES.md](NOTES.md) and add one entry to the [revision log](../REVISION-LOG.md).

## Related concurrency practice

Continue with the [senior concurrency topic guides](../concurrency/topics/README.md) and [unsolved implementation tasks](../concurrency/tasks/README.md). Record actual test and recall evidence in the [concurrency notes](../concurrency/NOTES.md).
