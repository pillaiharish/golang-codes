# Q8. What happens when a channel is closed?

[Concurrency index](../README.md) · [Q7: producer-consumer](../7-producer-consumer/README.md) · [Revision notes](../NOTES.md)

**Level 2 — channel fundamentals.**

## Task

Write independent tiny programs to demonstrate: receiving from a closed channel, sending to a closed channel, closing twice, and closing a nil channel. Predict whether each returns, blocks, or panics.

## Reference answers

| Operation | Outcome |
|---|---|
| Receive from a **closed, drained** channel | Immediately returns element zero value and `ok == false`. |
| Receive from a **closed, buffered** channel | Yields buffered values with `ok == true` until drained; then zero / `false`. |
| Send to a closed channel | Panics. |
| Close an already closed channel | Panics. |
| Close a nil channel | Panics. |
| Send to or receive from a nil channel | Blocks forever unless some external mechanism stops the goroutine/process. |

```go
ch := make(chan int, 1)
ch <- 7
close(ch)
fmt.Println(<-ch)       // 7
v, ok := <-ch
fmt.Println(v, ok)     // 0 false
```

Run each **panic** demonstration in a separate process so the other cases can execute. Never use panic-prone code in one shared demo expecting later lines to run.

## What-if scenarios

| Question | Answer |
|---|---|
| Why can a receiver safely range over a closed channel? | A close causes range to end after pending values have been read. |
| Should the receiver close the channel? | Usually no; the designated sender/owner knows when no sends remain. |
| How do many senders close safely? | Wait for all senders, then have a single coordinator close. |
| Can `sync.Once` make repeated `close` harmless? | It can serialize a known close owner, but cannot fix concurrent send-versus-close races by itself. |

**Recall:** Distinguish an empty but open channel, closed-and-buffered, closed-and-drained, and nil. Explain how `v, ok := <-ch` differs from checking for a zero payload.

References: [Go specification — close](https://go.dev/ref/spec#Close), [Go specification — receive](https://go.dev/ref/spec#Receive_operator).
