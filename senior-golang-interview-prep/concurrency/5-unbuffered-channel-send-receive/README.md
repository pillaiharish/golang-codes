# Q5. Send and receive on an unbuffered channel

[Concurrency index](../README.md) · [Q6: buffering](../6-buffered-channel-capacity/README.md) · [Revision notes](../NOTES.md)

**Level 2 — channel fundamentals.**

## Task

Create `ch := make(chan int)`. Start a goroutine that sends `42`; have `main` receive and print it. Explain which operation blocks and why.

## Reference answer

```go
package main
import "fmt"
func main() {
    ch := make(chan int)
    go func() { ch <- 42 }()
    fmt.Println(<-ch)
}
```

An unbuffered send can complete only when a matching receive participates (and vice versa). The receive in `main` waits until the sender can rendezvous. The successful send/receive also establishes a synchronization edge for memory visibility.

## What-if scenarios

| Question | Answer |
|---|---|
| What if `ch <- 42` happens in `main` immediately before `<-ch`? | `main` blocks at send: there is no concurrent receiver. With no runnable goroutines, the runtime normally reports a deadlock. |
| What if the goroutine sends twice but `main` receives once? | The sender can remain blocked on its second send; returning from `main` ends the process. |
| Can an unbuffered channel be used purely for a signal? | Yes; `chan struct{}` avoids transferring a payload. |
| Does a send guarantee the receiver has finished processing? | No. It guarantees a matching receive, not that subsequent receiver work has completed. |

**Recall:** Distinguish rendezvous, acknowledgment, and task completion. Explain why introducing a buffer changes blocking but not the need to design lifecycle.

References: [Go specification — channel types](https://go.dev/ref/spec#Channel_types), [Go memory model](https://go.dev/ref/mem).
