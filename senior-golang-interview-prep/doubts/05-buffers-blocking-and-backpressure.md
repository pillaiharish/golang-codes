# 5. Bounded buffers, blocking, and producer-consumer policy

**Follow-up:** [Slice length, growth, and explicit limits](07-slice-length-capacity-growth-and-limits.md).

[All doubts](README.md) · [Previous](04-increments-and-for-loops.md) · [Next: I/O](06-network-io-and-blocking-syscalls.md)

## Direct answer: three separate questions

1. **Capacity:** how much waiting work may be stored? Bounded means a defined maximum; unbounded means no configured finite queue limit, not infinite physical memory.
2. **Operation behavior:** when an operation cannot proceed, does it wait, return failure, drop, or wait until a deadline?
3. **Handoff:** is there a buffer between producer and consumer, or must they rendezvous?

These are descriptive terms, not Go keywords. `chan`, `select`, and `default` are language keywords; bounded, buffered, and blocking describe behavior. In Go, a channel's capacity is fixed when made: all buffered channels are bounded. There is no built-in unbounded channel. A dynamically growing queue needs additional implementation and synchronization.

## Combination matrix

Assume an open, non-nil queue/channel. “Nonblocking” below means the operation does not wait for queue readiness, not a formal lock-free progress guarantee.

| Design | Producer when unavailable | Consumer when empty | Example / trade-off |
|---|---|---|---|
| Bounded buffered, blocking | Wait for space | Wait for an item | Plain channel sends/receives; propagates backpressure |
| Bounded buffered, nonblocking enqueue | Return failure or apply explicit drop policy | May wait or try independently | `select` send with `default`; bounds queued work but may lose/reject work |
| Bounded buffered, timed/cancelable | Wait until space or cancellation/deadline | Policy independently chosen | Select on queue plus context |
| Unbuffered, blocking | Wait for a matching receive | Wait for a matching send | `make(chan T)`; zero queued items |
| Unbuffered, nonblocking attempt | Fail unless a receiver can rendezvous now | Fail unless a sender can rendezvous now | Try send/receive with default; no burst absorption |
| Unbounded queue, wait-on-empty receive | No configured full state; enqueue may still wait for locks/resources | Wait for an item | Growing queue plus condition variable; overload becomes memory growth |
| Unbounded queue, try-receive API | Enqueue resource/lock costs still exist | Return empty immediately | “Unbounded” says nothing about locking or scheduling |

Blocking versus nonblocking must be specified **per operation**. It is common to have a nonblocking producer and a blocking consumer on the same channel. A buffered channel can block; an unbuffered operation can be attempted without blocking.

## Three producer APIs

```go
func Send(ctx context.Context, q chan<- int, value int) error {
    select {
    case q <- value:
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func TrySend(q chan<- int, value int) bool {
    select {
    case q <- value:
        return true
    default:
        return false
    }
}

func TryReceive(q <-chan int) (value int, ready bool, open bool) {
    select {
    case value, open = <-q:
        return value, true, open
    default:
        return 0, false, true
    }
}
```

For `TryReceive`, `open` is meaningful only when `ready` is true. Closed is a ready receive, not an empty-but-open queue. These examples assume the caller owns a safe channel lifecycle: **a send to a closed channel panics even with a default case**. Select does not protect against concurrent close.

For cancelable send, if cancellation and sending are both ready, either can be chosen. This is not strict “no sends after cancellation” ordering. A stricter lifecycle contract needs additional coordination.

Run [the queue example](examples/queues/main.go), `go run ./examples/queues`. It deterministically demonstrates a full bounded queue rejecting an item, an unbuffered try-send failing without a receiver, and canceled waiting. No sleeps are used to infer readiness.

## Why default is not inherently a drop

```go
if !TrySend(queue, eventID) {
    // Caller decides: reject, count a drop, retry with a bound, or persist.
}
```

An unsuccessful try-send becomes loss only when the caller discards the work. Unbounded retry spins waste CPU. Launching a goroutine for each failed/blocking enqueue moves the backlog into goroutines and can defeat the intended memory bound. A channel's internal synchronization also means a default case does not establish lock-free or wait-free algorithmic guarantees.

## Full/empty/nil/closed behavior

| State | Plain send | Plain receive | Try-operation with default |
|---|---|---|---|
| Full buffered, open | Wait | Receive available item | Send can return failure |
| Empty buffered, open | Send if capacity available | Wait | Receive can return unavailable |
| Unbuffered, open | Needs receiver | Needs sender | Succeeds only if rendezvous is available |
| Nil | Wait forever | Wait forever | Case disabled; default can run |
| Closed | Panic | Drain queued values, then zero value and `ok=false` | Closed send can panic; receive remains ready |

Do not use `if len(q) < cap(q) { q <- v }` as a concurrent try-send: another producer may fill the channel between check and send. `select` performs the readiness decision as part of the channel operation.

## Capacity is not an end-to-end resource bound

A channel of 1024 slice values bounds the number of queued slice descriptors, not their payload bytes. Payloads can vary in size or retain large backing arrays. Memory also includes producer-held items, waiting goroutines, in-flight consumers, client-library batches, retries, and downstream buffers. For a real byte budget, enforce payload size and outstanding-byte limits in addition to queue length.

An in-memory blocking queue does not guarantee delivery: the process can crash after enqueue or after dequeue but before downstream acknowledgement. Durability, acknowledgement, retries, idempotency, ordering, and shutdown draining are separate contracts.

## Apply it to a single-reader telemetry-to-Kafka pipeline

With many producers, a channel of 1024, and one goroutine doing synchronous downstream sends:

- Producers using `select/default` have a bounded, nonblocking admission path.
- The consumer blocks waiting for an event when empty, then waits for downstream completion.
- The channel can fill while that one worker waits on Kafka, even if the Go runtime has many available threads.
- Increasing the channel delays overflow; it does not improve the sustained sink service rate.

Let arrival rate be lambda items/s and sustainable drain rate be mu items/s. During sustained overload, backlog grows approximately at `lambda - mu`. Starting with occupancy `q0`, approximate time to fill is `(capacity - q0)/(lambda - mu)` when lambda > mu. For example, 10,000 arrivals/s and 8,000 completions/s fill an initially empty 1024-item queue in about 0.512 s. This fluid approximation ignores burstiness and varying service times; average-rate equality does not prevent burst drops.

For an arrival burst lasting T, estimate backlog from excess arrivals over completed service, then check payload memory and acceptable queue delay. Little's law `L = lambda * W` describes stable-system averages under its assumptions; it is not a worst-case sizing guarantee.

Potential changes depend on the contract:

| Need | Candidate | Main cost / verification |
|---|---|---|
| Protect request latency; telemetry loss acceptable | Bounded try-send, sampling, explicit drop metrics | Measure loss; avoid logging each drop synchronously |
| Absorb finite bursts | Increase buffer within byte/latency budget | More retention and queue delay |
| Improve sustained throughput | Batch downstream sends; bounded worker concurrency | Verify client safety, ordering, ack semantics and actual throughput |
| Require durable delivery | Durable spool/broker and bounded retries | Disk limits, recovery and idempotency complexity |
| Backpressure upstream | Cancelable blocking enqueue | Request latency and blocked-producer budget |

Track accepted, rejected/full, canceled, downstream failed, retried, completed, queue depth, queue age, and in-flight work. If a single combined drop counter is intentional, document what it combines; it cannot reveal the dominant cause alone. A send accepted by the channel is not the same as an acknowledged Kafka record.

## Interview checkpoints

- Describe a queue using capacity, enqueue policy, dequeue policy, and durability separately.
- Explain why an unbuffered try-send can return immediately.
- Explain how 1024 queued pointers can retain far more than 1024 small objects.
- Explain why more goroutines or a larger buffer may fail to fix a slow sink.

Related: [channels module](../2-channels-select-and-backpressure/README.md), [publisher exercise](../7-bounded-concurrent-telemetry-publisher/README.md).

Sources: [channel types](https://go.dev/ref/spec#Channel_types), [send statements](https://go.dev/ref/spec#Send_statements), [select](https://go.dev/ref/spec#Select_statements), [memory model](https://go.dev/ref/mem), [pipeline cancellation](https://go.dev/blog/pipelines). Capacity calculations and design trade-offs above are worked engineering examples, not runtime guarantees.
