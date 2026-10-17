[← Senior Go interview preparation index](../README.md) · [Personal notes](NOTES.md) · [Revision log](../REVISION-LOG.md)

**Focused doubts:** [5. Bounded buffers, blocking, and producer-consumer policy](../doubts/05-buffers-blocking-and-backpressure.md) · [6. Goroutine blocking on network I/O versus blocking system calls](../doubts/06-network-io-and-blocking-syscalls.md)

> Study notes adapted from the supplied transcript. Technical examples are preserved for review; code blocks are illustrative unless separately validated.

# 4 — Context, Cancellation, and Graceful Shutdown

45 minutes. Shorter module, but it's the one that decides whether your Module 7 assignment is senior-grade or not.

## 4.1 Context is a cancellation *tree*, not a bag of values

```go
ctx := context.Background()          // root, never cancelled
ctx = context.WithCancelValue(...)   // derived
```

`WithCancel`, `WithTimeout`, `WithDeadline` all build a node in a tree. **Cancellation propagates downward, never upward.**

```go
parent, cancelParent := context.WithCancel(context.Background())
child, cancelChild := context.WithCancel(parent)

cancelParent() // child.Done() is now closed too
cancelChild()  // parent is completely unaffected
```

Cancellation is **one-way and irreversible**. There is no un-cancel.

### The leak people actually cause

```go
func handle() {
    ctx, _ := context.WithTimeout(parent, 5*time.Second)
    doWork(ctx) // cancel never called
}
```

Two costs: the child stays registered in the parent's children set until the parent is cancelled, and the **timer stays alive until the deadline fires**. Neither is catastrophic; both are how you get a slow, boring memory growth curve that looks like a leak.

> **Rule: `defer cancel()` immediately after creating a cancellable context — even for `WithTimeout`.** The returned cancel is not optional bookkeeping.

---

## 4.2 Cancellation is cooperative

This is the single most important sentence in the module:

> **Cancelling a context does not stop anything. It announces that someone should stop.**

A goroutine stops only if it **chooses** to look:

```go
go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            work()
        }
    }
}()
```

A goroutine in a tight CPU loop with no `select`, or inside a non-cancellable blocking syscall, **will not stop**. Context is a signal, not a preemption.

### The `for range` trap

```go
for v := range ch {   // ctx.Done() is invisible here
    process(v)
}
```

`range` over a channel only ends when the channel is **closed**. If you need cancellation, you must write the `select` yourself:

```go
for {
    select {
    case <-ctx.Done():
        return
    case v, ok := <-ch:
        if !ok {
            return
        }
        process(v)
    }
}
```

And the symmetric trap on the **send** side — a worker blocked sending into a full queue needs the same treatment:

```go
select {
case out <- result:
case <-ctx.Done():
    return ctx.Err()
}
```

Forgetting the send-side select is the classic "my shutdown hangs" bug.

---

## 4.3 Don't put a context in a struct

```go
type Publisher struct {
    ctx context.Context // almost always wrong
}
```

Contexts are **per-request/per-operation scope**, not object lifetime. Storing one creates exactly the bug you'd expect: a request-scoped context outlives the request, and the object silently dies when the request is cancelled — or worse, keeps a cancelled context forever and every later operation fails instantly.

**Pass contexts as the first parameter.** The one legitimate exception is a long-lived *root* shutdown context owned by the service itself — and that's a different thing with a different name, which is why I'd call it `shutdownCtx` and keep it clearly separate.

---

## 4.4 `context.WithoutCancel` — the detail that signals currency

Go 1.21 added it, and it solves a real problem: **work that must finish after the HTTP response has been sent.**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    enqueue(r.Context(), event)

    // this must survive the request being cancelled by the client
    bg := context.WithoutCancel(r.Context())
    go func() {
        auditLog(bg, event) // keeps values, drops cancellation
    }()
}
```

Without it, a client disconnect kills your audit write. With it, you keep the request's values (request ID, trace context) while detaching from its lifetime. Knowing this exists is a small, cheap signal of being current.

---

## 4.5 The graceful shutdown ladder

For a telemetry publisher, shutdown is a **sequence**, and the order is the whole design:

```go
func (s *Service) Run(ctx context.Context) error {
    rootCtx, cancel := context.WithCancel(ctx)
    defer cancel()

    var wg sync.WaitGroup

    // every goroutine the service owns is registered here
    for i := 0; i < s.workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            s.worker(rootCtx)
        }()
    }

    <-rootCtx.Done() // shutdown signalled

    // 1. Stop accepting new work — Publish now returns an error
    s.closed.Store(true)

    // 2. Close ingress so consumers can drain rather than abort
    close(s.queue)

    // 3. Wait for in-flight work, with a deadline
    done := make(chan struct{})
    go func() { wg.Wait(); close(done) }()

    select {
    case <-done:
        log.Println("drained cleanly")
    case <-time.After(s.drainTimeout):
        log.Println("drain timeout — forcing exit, some events lost")
    }

    // 4. Flush/close the Kafka producer — it has its own internal buffer
    return s.producer.Close(flushCtx)
}
```

### Why each step is where it is

**Close ingress before draining.** If you cancel workers first, they abandon queued events. If you close the queue first, workers see the drain and finish what's already buffered. Order matters enormously, and getting it backwards is the difference between "lost everything queued" and "lost nothing."

**Set a `closed` flag before closing the channel.** Otherwise a producer races the close and panics with `send on closed channel`. This is exactly the channel-ownership problem from Module 2 reappearing at system scale.

**Always bound the drain.** An unbounded `wg.Wait()` means one stuck Kafka broker hangs your pod forever, and Kubernetes escalates to `SIGKILL` — which loses *more* than a graceful timeout would have. A bounded drain is strictly better than an infinite one.

**Close the producer last.** The Kafka client has its **own** internal buffer and batching. Closing your channel doesn't flush it. If you exit without `Close()`, you lose whatever was batched and not yet ACKed — a genuinely nasty, easy-to-miss source of silent loss.

### The two-signal pattern

In production I want a second signal to mean *stop trying*:

- first `SIGTERM` → graceful: stop accepting, drain, flush, exit
- second `SIGTERM` / deadline exceeded → forced: close everything, exit now

Because the failure mode you're protecting against is a pod that refuses to die and gets `SIGKILL`ed at the worst possible moment.

### HTTP alongside it

`srv.Shutdown(ctx)` stops accepting connections and waits for active ones — it does **not** kill in-flight handlers. So the ordering is: stop accepting HTTP → let handlers finish or time out → then drain the queue → then close the producer. And after the deadline, `srv.Close()` as the hammer.

---

## Interview questions and answers

## Q1: What happens when a parent context is cancelled?

All **descendants** are cancelled too — cancellation flows down the tree and never up. Their `Done()` channels close and `Err()` returns `context.Canceled`, or `DeadlineExceeded` if a deadline was the cause.

It's one-way and irreversible. Cancelling a child leaves the parent untouched.

And the practical corollary: **always `defer cancel()`**, including on `WithTimeout` — otherwise the child stays registered in the parent and the timer stays armed until the deadline. Small, slow, and exactly the kind of thing that shows up as a mystery growth curve six weeks later.

---

## Q2: Does cancelling a context automatically stop a goroutine?

**No. Cancellation is cooperative.** It closes a channel; it doesn't preempt anything.

A goroutine stops only if it selects on `ctx.Done()`. It will **not** stop if it's:

- in a tight CPU loop with no `select`
- blocked in a non-cancellable syscall
- in a `for range ch` that never gets closed
- blocked sending to a full channel without a `select` on `Done()`

The senior framing: context is a **broadcast signal**, and honoring it is the receiver's responsibility. So the real question in a code review isn't "did we cancel?" — it's "**who is listening, and where can they fail to notice?**"

---

## Q3: How do you stop a worker blocked on a channel receive?

Wrap the receive in a `select`:

```go
select {
case <-ctx.Done():
    return
case v, ok := <-ch:
    if !ok {
        return
    }
    process(v)
}
```

You cannot interrupt an already-committed receive — the `select` must be there *before* it blocks.

Two things I'd add unprompted:

- **Do the same on the send side.** A worker blocked sending into a full queue needs `select { case out <- v: case <-ctx.Done(): return }`. This is the most commonly forgotten half and the usual cause of "shutdown hangs."
- **Closing the channel is an alternative, but it changes the contract.** Close means *drain then finish*; cancel means *abandon now*. Conflating them is how you silently drop accepted events during a deploy. Pick deliberately.

---

## Q4: How do you ensure no goroutine is left running after shutdown?

By **ownership and verification**, not hope.

1. **Own every goroutine at the top level** with one `WaitGroup`, `Add` in the parent, `defer Done()` inside.
2. **One root context** for the whole service; cancel it to signal everyone.
3. **Bound the wait** with a deadline, then force.
4. **Verify**, don't assume — in tests, check `runtime.NumGoroutine()` before and after (with a small retry, since goroutine exit isn't instantaneous), or use a leak-detection helper. In production, alert on goroutine count over time.

The key insight to state: **you can't verify absence by reading the code.** Every `go` statement is a potential leak, so the discipline is that no goroutine is created without a registered owner and a defined exit condition. That's an architectural rule, and it's the one I'd want a senior engineer to enforce in review.

---

## Q5: How would you shut down a Kafka publishing service without unnecessarily losing accepted messages?

This is a production-grade shutdown question, and the answer has to confront an uncomfortable truth first:

> **"Accepted" and "durable" are different promises.** If you ACK to the client at *enqueue* time, you've promised something a pod crash can't honor. No shutdown procedure fixes that — only the ACK point does.

So the answer has two parts.

**The shutdown sequence**, in order:

1. Set `closed` — `Publish` now rejects, so nobody races the close.
2. Close the ingress channel — workers **drain** rather than abort.
3. Wait on the `WaitGroup` **with a deadline**.
4. `Close()` the Kafka producer so its internal batch buffer flushes.
5. Exit.

Ordering is the design. Cancel workers before draining and you discard the queue. Skip producer `Close()` and you lose the batched-but-unACKed tail — a small, consistent, maddening loss that looks like random Kafka flakiness.

**The honest policy question:**

- If you ACK at enqueue → best-effort. Fine for debug telemetry, provided the SLA says so and you **measure the loss** during deploys.
- If you need at-least-once → **ACK only after Kafka confirms.** You pay latency, and you accept duplicates — which means consumers must be idempotent.
- If you need both low latency and durability → a **durable spill** (local WAL / outbox) that survives restart, with replay on boot. More complexity, real latency, and the only way to actually keep the promise.

And then the part most candidates omit: **you must measure it.** Emit a counter for "events accepted but not delivered during shutdown," and alert on it. Otherwise "graceful" is an adjective, not a property.

That last move — *turning an adjective into a metric* — is the strongest thing you can do in this answer, and it's exactly the mindset Module 7 is testing.

---

---

## Revision checkpoint

Try answering from memory *before* opening the relevant section. Mark the boxes only after you can explain or demonstrate the result.

- [ ] Explain cancellation propagation across a context tree and why cancellation is cooperative.
- [ ] Make a blocked channel send or receive responsive to cancellation.
- [ ] Explain `context.WithoutCancel` and the risk of breaking cancellation propagation.
- [ ] State the graceful shutdown order for producers, queue, workers, and downstream sink.
- [ ] Explain how a deadline changes accepted-but-not-delivered accounting.

**After review:** record your confidence, uncertainties, and evidence in [NOTES.md](NOTES.md) and add one entry to the [revision log](../REVISION-LOG.md).
