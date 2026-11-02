# Q10. Create and fix deadlocks — worked solution

[Back to Q10 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q5: unbuffered send](../../5-unbuffered-channel-send-receive/README.md) · [Q7: producer-consumer](../../7-producer-consumer/README.md) · [Revision notes](../../NOTES.md)

**Level 2 — channel fundamentals.**

## Task

Create four small cases, predict the outcome, and fix them:

1. Send on an unbuffered channel without a receiver.
2. Receive from an open empty channel without a sender.
3. Range over a channel when the producer never closes it.
4. Start a goroutine but return from `main` before joining it.

## Reference answers

| Broken case | Mechanism | Fix |
|---|---|---|
| `ch := make(chan int); ch <- 1` | Sender blocks waiting for a receiver. | Start a matching receiver, or redesign ownership/buffering. |
| `ch := make(chan int); <-ch` | Receiver blocks waiting for a value/close. | Provide a sender or close from the owner. |
| `for v := range ch` without close | After final value, range waits for a future value/close. | Producer/coordinator closes after last send. |
| `go work(); return` from `main` | Program exits before work is joined. | Use `WaitGroup` or a completion signal. **This is not itself a deadlock**; it is an early-exit lifecycle bug. |

If no goroutine can progress, the Go runtime may detect a deadlock. A goroutine blocked on network I/O or a live timer can make runtime deadlock detection less straightforward; passing a test once is not a proof of leak freedom.

## Follow-up scenarios

- **A buffered channel of size 1:** The first send may succeed; a second send blocks if nothing receives.
- **A never-closing channel with cancellation:** Use `select` on a cancellation signal to let a blocked worker stop.
- **Two mutexes locked in opposite orders:** A classic circular-wait deadlock; establish a consistent acquisition order.
- **A test that hangs:** Use a test timeout, goroutine stack dumps, and deterministic barriers to identify the wait cycle.

**Recall:** For each bug draw the goroutines, the channel operation they await, and the event that would unblock it. Explain who owns that event.

References: [Go blog — pipelines and cancellation](https://go.dev/blog/pipelines), [Go memory model](https://go.dev/ref/mem).
