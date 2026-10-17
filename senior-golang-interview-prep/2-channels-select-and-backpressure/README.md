[← Senior Go interview preparation index](../README.md) · [Personal notes](NOTES.md) · [Revision log](../REVISION-LOG.md)

**Focused doubts:** [5. Bounded buffers, blocking, and producer-consumer policy](../doubts/05-buffers-blocking-and-backpressure.md)

> Study notes adapted from the supplied transcript. Technical examples are preserved for review; code blocks are illustrative unless separately validated.

# 2 — Channels, Select, and Backpressure

This is an important module for senior Go platform interviews because it maps directly to:

- event pipelines
- telemetry ingestion
- Kafka producers/consumers
- bounded queues
- backpressure
- lossy vs durable delivery

---

## 2.1 Channels are not just queues

A channel in Go is both:

1. a communication mechanism
2. a synchronization mechanism

That dual nature is what trips people up.

A channel can:

- transfer values
- block until a sender/receiver is ready
- signal completion by closing
- fan out work
- fan in results
- implement bounded queues

---

## 2.2 Unbuffered vs buffered channels

### Unbuffered channel

```go
ch := make(chan int)
```

An unbuffered channel requires a **rendezvous**:

- a send blocks until a receiver is ready
- a receive blocks until a sender is ready

This gives strong synchronization semantics.

### Example

```go
ch := make(chan int)

go func() {
    ch <- 42 // blocks until another goroutine receives
}()

v := <-ch // completes the rendezvous
fmt.Println(v)
```

### Buffered channel

```go
ch := make(chan int, 3)
```

A buffered channel allows sends to proceed without an immediate receiver, up to the buffer capacity.

```go
ch := make(chan int, 2)

ch <- 1
ch <- 2
// ch <- 3 would block here until a receive happens
```

Buffered channels are useful for decoupling producers and consumers, but they introduce capacity planning and backpressure questions.

---

## 2.3 Channel direction

You can restrict channel direction in function signatures.

### Send-only

```go
func publish(ch chan<- int) {
    ch <- 1
}
```

### Receive-only

```go
func consume(ch <-chan int) {
    v := <-ch
    _ = v
}
```

This improves API design by making ownership and intent explicit.

A producer should usually receive a send-only channel.

A consumer should usually receive a receive-only channel.

---

## 2.4 Closing channels

Closing a channel has specific semantics.

### Reading from a closed channel

If there are still buffered values, you can read them.

After the buffer is drained, reads return the zero value and `ok == false`.

```go
ch := make(chan int, 1)
ch <- 10
close(ch)

fmt.Println(<-ch) // 10
v, ok := <-ch
fmt.Println(v, ok) // 0 false
```

### Sending to a closed channel

This panics.

```go
ch := make(chan int)
close(ch)
ch <- 1 // panic: send on closed channel
```

### Closing a nil channel

This also panics.

```go
var ch chan int
close(ch) // panic
```

### Reading from a nil channel

Blocks forever.

```go
var ch chan int
<-ch // blocks forever
```

### Sending to a nil channel

Blocks forever.

```go
var ch chan int
ch <- 1 // blocks forever
```

---

## 2.5 Channel ownership and safe closing

A senior-level rule of thumb:

> The sender side that controls production should usually own closing the channel.

But with multiple producers, this becomes tricky.

### Safe pattern 1: single producer closes the channel

```go
func produce(ch chan<- int) {
    defer close(ch)
    for i := 0; i < 3; i++ {
        ch <- i
    }
}
```

### Dangerous pattern: multiple producers close independently

If multiple goroutines send to the same channel and more than one may close it, you risk:

- panic from sending on closed channel
- panic from closing an already closed channel

### Safer multi-producer pattern

Use a separate done/stop signal and a coordinator.

Example:

```go
func worker(ctx context.Context, jobs <-chan int, done chan<- struct{}) {
    defer func() { done <- struct{}{} }()
    for {
        select {
        case <-ctx.Done():
            return
        case j, ok := <-jobs:
            if !ok {
                return
            }
            _ = j
        }
    }
}
```

A common senior pattern is:

- producers send to a data channel
- producers do **not** close the data channel
- a separate closer waits for all producers to finish, then closes it

Example with WaitGroup:

```go
func fanIn(ctx context.Context, producers int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    for i := 0; i < producers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 3; j++ {
                select {
                case <-ctx.Done():
                    return
                case out <- id*10 + j:
                }
            }
        }(i)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

This is a clean pattern because:

- producers never close `out`
- a single goroutine closes `out`
- consumers can range safely

---

## 2.6 `select`

`select` lets a goroutine wait on multiple channel operations.

```go
select {
case v := <-ch1:
    // receive from ch1
case ch2 <- 10:
    // send to ch2
case <-time.After(time.Second):
    // timeout
default:
    // nonblocking fallback
}
```

Important properties:

- If multiple cases are ready, Go chooses one pseudo-randomly.
- If no case is ready and there is no `default`, select blocks.
- If `default` is present, select does not block.
- Nil channel cases are ignored.

---

## 2.7 Fairness in `select`

### Question: Does `select` guarantee fairness?

No.

If multiple cases are ready, Go chooses one using a pseudo-random policy.

It does **not** guarantee:

- round-robin
- FIFO
- strict fairness

If fairness matters, you need to design it explicitly.

---

## 2.8 Nonblocking operations

A nonblocking send:

```go
select {
case ch <- event:
    // accepted
default:
    // dropped
}
```

A nonblocking receive:

```go
select {
case v := <-ch:
    // got value
default:
    // no value available
}
```

This is extremely important for bounded queues and backpressure.

---

## 2.9 Backpressure

Backpressure is what happens when a producer is faster than a consumer.

You have three main choices:

1. **Block the producer**
2. **Drop the event**
3. **Buffer durably**

Each has trade-offs.

---

### Option 1: Block the producer

```go
ch <- event
```

Pros:

- no loss
- natural pressure upstream

Cons:

- can cause request latency to rise
- can exhaust goroutines if many producers block
- can turn a downstream slowdown into a full service outage

Best when:

- losing data is unacceptable
- upstream can tolerate delays
- the system has enough headroom

---

### Option 2: Drop the event

```go
select {
case ch <- event:
    accepted.Add(1)
default:
    dropped.Add(1)
}
```

Pros:

- protects service availability
- keeps latency stable
- prevents unbounded queue growth

Cons:

- data loss
- may violate correctness requirements

Best when:

- data is telemetry/metrics/traces where some loss is acceptable
- you can sample
- availability matters more than completeness

---

### Option 3: Buffer durably

Examples:

- local disk queue
- Kafka itself
- persistent outbox
- WAL-style buffer

Pros:

- survives restarts/crashes
- better delivery guarantees

Cons:
- more complexity
- higher latency
- storage and replay concerns

Best when:

- events are billing, audit, security, or high-value signals
- at-least-once delivery is required

---

## 2.10 Bounded queues

A bounded queue is essential in production.

Unbounded in-memory queues are dangerous because:

- memory grows without limit
- GC pressure increases
- latency worsens
- the process may OOM

A bounded queue forces an explicit policy:

- accept
- block
- drop
- reject with backpressure signal

This is a senior architectural decision, not just a coding detail.

---

## Interview questions and senior-level answers

## Question 1: What happens when you write to a full buffered channel?

### Good answer

The sender blocks until space becomes available.

If the send is inside a `select` with a `default`, then the `default` branch executes instead of blocking.

If the channel is closed, sending panics.

### Example

Blocking send:

```go
ch := make(chan int, 1)
ch <- 1
ch <- 2 // blocks until a receive happens
```

Nonblocking send:

```go
select {
case ch <- 2:
    fmt.Println("sent")
default:
    fmt.Println("dropped")
}
```

---

## Question 2: What happens when you read from a closed channel?

### Good answer

Reading from a closed channel does **not** panic.

If buffered values remain, they can still be read.

Once drained, reads return the zero value immediately.

Using the two-value form:

```go
v, ok := <-ch
```

- `ok == true` while values remain
- `ok == false` after the closed channel is drained

### Example

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
close(ch)

fmt.Println(<-ch) // 1
fmt.Println(<-ch) // 2

v, ok := <-ch
fmt.Println(v, ok) // 0 false
```

---

## Question 3: What happens when you send to a closed or nil channel?

### Send to closed channel

Panic.

```go
ch := make(chan int)
close(ch)
ch <- 1 // panic
```

### Send to nil channel

Blocks forever.

```go
var ch chan int
ch <- 1 // blocks forever
```

### Extra: close(nil)

Also panics.

```go
var ch chan int
close(ch) // panic
```

This is a classic code-output trap.

---

## Question 4: Does `select` guarantee fairness when multiple cases are ready?

### Good answer

No.

If multiple cases are ready, Go chooses one pseudo-randomly.

There is no guarantee of:

- round-robin
- priority
- FIFO
- strict fairness

If fairness or priority is required, it must be implemented explicitly.

### Senior-level extra

For example, if you need priority between a control channel and a data channel, you may need nested selects or an explicit scheduling loop. But that adds complexity and should be justified.

---

## Question 5: How would you implement a nonblocking queue with a drop counter?

### Good answer

Use a buffered channel and a `select` with `default`.

```go
type Queue struct {
    ch      chan Event
    dropped atomic.Int64
}

type Event struct {
    ID string
}

func NewQueue(capacity int) *Queue {
    return &Queue{
        ch: make(chan Event, capacity),
    }
}

func (q *Queue) TryPublish(e Event) bool {
    select {
    case q.ch <- e:
        return true
    default:
        q.dropped.Add(1)
        return false
    }
}
```

This gives:

- bounded memory
- nonblocking publish
- explicit drop accounting

### Senior-level extra

In production, I would also export:

- queue length
- drop rate
- publish latency
- consumer lag

Those metrics are essential to know whether 9% loss is stable, improving, or catastrophic.

---

## Question 6: How would you prevent a slow Kafka producer from blocking an API?

This is highly relevant to senior Go platform engineering interviews.

### Good answer

I would isolate the API request path from the Kafka publishing path.

Typical design:

1. API handler receives event
2. API handler does minimal work
3. event is placed into a bounded in-memory queue
4. separate publisher workers send to Kafka
5. if the queue is full, decide explicitly:
   - drop
   - reject
   - return degraded response
   - spill to durable buffer

### Example

```go
select {
case q.ch <- event:
    w.WriteHeader(http.StatusAccepted)
default:
    dropped.Add(1)
    w.WriteHeader(http.StatusTooManyRequests)
}
```

### Senior-level extra

If Kafka is slow, you do not want every HTTP handler goroutine to block on Kafka sends.

Instead, you want:

- bounded concurrency
- bounded queue depth
- timeouts
- cancellation
- metrics
- possibly batching

A slow downstream dependency should degrade gracefully, not consume all service resources.

---

## Question 7: When should the system block producers, drop messages, or buffer them durably?

### Good answer

It depends on the data’s value and the system’s SLOs.

### Block producers when:

- data loss is unacceptable
- producers can slow down
- upstream has natural backpressure
- latency is less important than completeness

Example:

- financial transactions
- audit logs
- security events

### Drop messages when:

- some loss is acceptable
- availability and low latency matter more
- the data is high-volume and redundant
- sampling is acceptable

Example:

- debug telemetry
- performance metrics
- traces where sampling is expected

### Buffer durably when:

- delivery must survive crashes/restarts
- at-least-once semantics are required
- downstream outages are expected
- replay is needed

Example:

- billing events
- usage metering
- compliance logs

### Senior-level extra

The best answer includes **measurement**:

- acceptable loss rate
- p99 latency target
- queue depth threshold
- producer throughput
- consumer throughput
- downstream error rate

Without those, the decision is just opinion.

---

## Mini coding question

You asked about this code:

```go
func main() {
    ch := make(chan int, 1)

    ch <- 10

    select {
    case ch <- 20:
        fmt.Println("sent")
    default:
        fmt.Println("dropped")
    }

    close(ch)

    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

## Output

```text
dropped
10
0
```

## Why?

Step by step:

### Step 1

```go
ch := make(chan int, 1)
```

The channel has capacity 1.

### Step 2

```go
ch <- 10
```

This succeeds immediately because the buffer has space.

Channel now contains:

```text
[10]
```

### Step 3

```go
select {
case ch <- 20:
    fmt.Println("sent")
default:
    fmt.Println("dropped")
}
```

The send `ch <- 20` cannot proceed because the buffer is full.

Because there is a `default` case, the select does not block.

So it prints:

```text
dropped
```

### Step 4

```go
close(ch)
```

The channel is closed, but the buffered value `10` is still there.

### Step 5

```go
fmt.Println(<-ch)
```

Reads the buffered value:

```text
10
```

### Step 6

```go
fmt.Println(<-ch)
```

Now the closed channel is drained, so it returns the zero value for `int`:

```text
0
```

---

## How behavior changes if we remove the buffer

```go
ch := make(chan int)
ch <- 10
```

Now the channel is unbuffered.

There is no receiver at that moment, so:

```go
ch <- 10
```

blocks forever.

The program deadlocks.

Output: none.

---

## How behavior changes if we remove the `default` case

Original with buffer but no default:

```go
select {
case ch <- 20:
    fmt.Println("sent")
}
```

Because the buffer is full and there is no receiver, this select blocks forever.

`close(ch)` is never reached.

So the program deadlocks.

---

## How behavior changes if we remove `close(ch)`

Then after printing:

```go
fmt.Println(<-ch) // 10
```

the next read:

```go
fmt.Println(<-ch)
```

blocks forever because the channel is empty and not closed.

So the program deadlocks.

---

## A useful senior pattern: context-aware blocking publish

If you want a publish method that can block but still respect cancellation:

```go
func Publish(ctx context.Context, ch chan<- Event, e Event) error {
    select {
    case ch <- e:
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

This is blocking, but cancellable.

If you want nonblocking publish, use `default`:

```go
func TryPublish(ch chan<- Event, e Event) bool {
    select {
    case ch <- e:
        return true
    default:
        return false
    }
}
```

A senior engineer should clearly distinguish these two behaviors.

---

## A useful senior pattern: worker pool with bounded input

```go
func runWorkers(ctx context.Context, jobs <-chan Event, workers int) {
    var wg sync.WaitGroup

    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for {
                select {
                case <-ctx.Done():
                    return
                case job, ok := <-jobs:
                    if !ok {
                        return
                    }
                    process(job)
                }
            }
        }()
    }

    wg.Wait()
}

func process(e Event) {
    // send to Kafka, HTTP, etc.
}
```

This gives:

- bounded concurrency
- clean shutdown
- no inherent goroutine leak if contexts and channels are managed correctly

---

## Senior-level takeaways from Module 2

When discussing Kafka telemetry pipelines, the underlying engineering questions are:

- Can you design bounded pipelines?
- Do you understand backpressure?
- Can you avoid goroutine leaks?
- Can you reason about loss vs durability?
- Can you protect an API from a slow downstream dependency?
- Can you measure whether your design is actually better?

The best answers combine:

- channel semantics
- context cancellation
- worker pools
- metrics
- explicit failure policy

---

---

## Revision checkpoint

Try answering from memory *before* opening the relevant section. Mark the boxes only after you can explain or demonstrate the result.

- [ ] Predict sends and receives for nil, open, full, buffered, and closed channels.
- [ ] Explain safe channel closing with multiple producers and the coordinator pattern.
- [ ] Explain `select` readiness, `default`, and the absence of strict fairness.
- [ ] Choose block, drop, reject, or durable buffer for three different event types.
- [ ] Predict the mini coding example output and its three deadlock variations.

**After review:** record your confidence, uncertainties, and evidence in [NOTES.md](NOTES.md) and add one entry to the [revision log](../REVISION-LOG.md).
