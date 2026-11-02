# Q6. Understand buffered channels — worked solution

[Back to Q6 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q5: unbuffered channel](../../5-unbuffered-channel-send-receive/README.md) · [Revision notes](../../NOTES.md)

**Level 2 — channel fundamentals.**

## Task

Create a capacity-two channel. Send two integers without starting a receiver. Predict what happens on a third send, then add a receiver to unblock it.

## Reference answer

```go
ch := make(chan int, 2)
ch <- 1 // buffered: succeeds
ch <- 2 // buffered: succeeds

go func() { fmt.Println(<-ch) }()
ch <- 3 // proceeds only once capacity becomes available
close(ch)
for v := range ch { fmt.Println(v) }
```

With no concurrent receiver, the third send would block indefinitely. A receive releases a slot, allowing it to continue. FIFO order applies to values in a channel; scheduling between independent goroutines is still nondeterministic.

## What-if scenarios

| Question | Answer |
|---|---|
| What if capacity is zero? | The channel is unbuffered: send and receive must rendezvous. |
| What if the buffer is empty and no sender exists? | A receive blocks unless the channel has been closed. |
| Is capacity two a promise that the application never blocks? | No. The third outstanding send blocks when full. |
| Does a buffered channel cap the number of spawned goroutines? | No. It caps queued elements, not goroutines waiting to send. |
| What if the buffer has values when closed? | Receivers drain buffered values first, then receive the zero value with `ok == false`. |

**Recall:** Explain `len(ch)` versus `cap(ch)`, what work is considered *accepted*, and how buffer size affects backpressure and memory use.

References: [Go specification — channels](https://go.dev/ref/spec#Channel_types), [Go blog — pipelines](https://go.dev/blog/pipelines).
