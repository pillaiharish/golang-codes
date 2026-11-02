# Thirty Go concurrency theory questions and short answers

[Concurrency index](README.md) · [Study plan](STUDY-PLAN.md) · [Evidence log](NOTES.md)

**Practice:** Hide the answer column and explain the mechanism, one counterexample, and one safe pattern before reading it. These are concise recall cues; the linked numbered exercises carry deeper scenarios.

| # | Interview question | Short answer / key trap |
|---|---|---|
| 1 | What is a goroutine? | A lightweight concurrently executing function managed by the Go runtime; creation with `go f()` does not wait for completion. |
| 2 | How does it differ from an OS thread? | Many goroutines are multiplexed over OS threads by the scheduler; they have small growable stacks and runtime-managed blocking/scheduling. |
| 3 | What is a channel? | A typed communication mechanism with send, receive, optional buffering and close semantics; it can also provide synchronization. |
| 4 | Buffered versus unbuffered? | Unbuffered operations rendezvous; buffered sends proceed while space exists, receives while values exist. Capacity does not mean unlimited progress. |
| 5 | Send on nil channel? | Blocks forever (unless another surrounding `select` case or process event allows progress). |
| 6 | Receive from nil channel? | Blocks forever; a nil case in `select` is disabled. |
| 7 | Close a nil channel? | Panics. |
| 8 | Send on closed channel? | Panics. |
| 9 | Receive from closed channel? | Drains buffered values first; once drained returns element zero value with `ok=false`. |
| 10 | Who closes a channel? | The owner who can guarantee no future sends, usually the sole sender or a coordinator after all senders exit. |
| 11 | How avoid closing it twice? | Make close ownership exclusive or coordinate with one lifecycle transition; also prevent send-versus-close races. |
| 12 | What is a data race? | Conflicting memory accesses without required synchronization, at least one write. Use `-race` and code-level memory-model reasoning. |
| 13 | How to detect races? | `go test -race`/`go run -race` on representative concurrent paths; clean observed runs don't prove universal safety. |
| 14 | What is a deadlock? | Work cannot make progress because goroutines wait in a cycle or on events that never occur. Not all early exits or leaks are deadlocks. |
| 15 | What is a goroutine leak? | A goroutine remains alive unintentionally, often blocked forever, retaining stack and referenced memory. |
| 16 | Channels versus mutexes? | Channels express ownership transfer, signaling and pipelines; mutexes protect shared mutable invariants. Choose by contract, not slogan. |
| 17 | What does `select` do? | Chooses one ready channel operation; blocks if none ready unless `default`; multiple ready cases are selected pseudo-randomly. |
| 18 | Why use `context.Context`? | Propagates request-scoped cancellation, deadlines and values; workers must explicitly cooperate. |
| 19 | Purpose of `sync.WaitGroup`? | Registers and joins tasks via `Add`, `Done`, `Wait`; it neither orders workers nor serializes conflicting writes. |
| 20 | Purpose of `sync.Mutex`? | Exclusive access to a critical section and its protected invariant, with lock/unlock synchronization. |
| 21 | Purpose of `sync.RWMutex`? | Multiple concurrent readers or one writer; performance depends on contention and workload, not guaranteed faster than Mutex. |
| 22 | Purpose of `sync.Once`? | Execute initialization at most once and safely publish its result; it does not automatically retry failed/panicking initialization. |
| 23 | What is a worker pool? | A bounded set of workers consuming a job queue with explicit admission, result, shutdown and failure contracts. |
| 24 | What is fan-out/fan-in? | Fan-out distributes work to multiple workers; fan-in merges result streams, closing only after all senders finish. |
| 25 | What is backpressure? | Downstream saturation slows/rejects upstream work; define queue capacity, rejection/drop policy and metrics. |
| 26 | What does “share memory by communicating” mean? | Prefer transferring ownership/data through channels where appropriate rather than allowing uncontrolled shared mutation. |
| 27 | What is “communicate by sharing memory”? | Shared state plus mutexes/atomics is also valid when ownership and invariants are explicit; don't turn a design proverb into a ban. |
| 28 | What is happens-before? | A partial ordering from sequenced operations and synchronization edges that determines visibility and race freedom. `Sleep` is not such an edge. |
| 29 | How debug a blocked goroutine? | Inspect goroutine stack profiles, channels/locks, queue depth, request traces and cancellation ownership; reproduce with controlled barriers. |
| 30 | How implement graceful shutdown? | Stop admission, establish sender/close ownership, drain or cancel with a deadline, join cooperative work, close sinks and account for unresolved work. |

## Follow-up practice prompts

1. **Nil versus closed:** Explain why `select` on a nil channel disables a case but a closed channel may be permanently ready.
2. **WaitGroup versus mutex:** Write a counterexample where `wg.Wait()` joins all goroutines yet a data race still occurred.
3. **Semaphore versus rate:** A service starts 10 jobs/second, each taking 30 seconds. What prevents runaway active work?
4. **Cancellation versus termination:** What evidence proves a worker exited after `ctx.Done()` was closed?
5. **Production incident:** What would you measure to distinguish a full Kafka telemetry queue, a stalled sender, and a slow broker acknowledgment?

## Concept cross-links

- [Q1–Q4](README.md#level-1-goroutine-basics): lifecycle, order and panic.
- [Q5–Q10](README.md#level-2-channel-fundamentals): channel behavior and deadlocks.
- [Q11–Q15](README.md#level-3-select-timeouts-and-cancellation): `select` and cancellation.
- [Q16–Q20](README.md#level-4-shared-state-and-synchronization): races and coordination.
- [Q21–Q28](README.md#level-5-core-concurrency-patterns): common patterns.
- [Q29–Q33](README.md#level-6-production-correctness): production lifecycle and debugging.

Sources: [Go language specification](https://go.dev/ref/spec), [Go memory model](https://go.dev/ref/mem), [`sync` package](https://pkg.go.dev/sync), [Go concurrency pipelines](https://go.dev/blog/pipelines).
