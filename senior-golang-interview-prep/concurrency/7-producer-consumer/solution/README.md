# Q7. Build a simple producer-consumer — worked solution

[Back to Q7 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q8: closed channels](../../8-closed-channel-behavior/README.md) · [Revision notes](../../NOTES.md)

**Level 2 — channel fundamentals.**

## Task

One producer sends `1..10` into a channel, closes it after the final send, and one consumer prints all received values. Explain who owns the close operation and when the consumer stops.

## Reference answer

```go
package main
import "fmt"
func main() {
    ch := make(chan int)
    go func() {
        defer close(ch) // the sole producer owns closure
        for i := 1; i <= 10; i++ { ch <- i }
    }()
    for v := range ch { fmt.Println(v) }
}
```

`range` receives until the producer closes the channel and any buffered values are drained. One producer means send order is preserved. The main goroutine's range also acts as the join for this simple pipeline: it finishes after production is closed and drained.

## What-if scenarios

| Question | Answer |
|---|---|
| What if the producer never closes `ch`? | The consumer blocks after the last value. |
| What if the consumer closes while the producer is sending? | The sender can panic on a subsequent send; don't assign close ownership to the consumer. |
| What if there are multiple producers? | Coordinate all producers' completion before one designated owner closes the channel. |
| What if the consumer stops after three values? | The producer can become stuck sending; add cancellation or another drain policy. |
| Does a channel close imply all processing is finished? | No. It means no more values will be sent, not that consumers have finished handling earlier values. |

**Recall:** Rebuild this with a separate consumer goroutine and a WaitGroup, and name which goroutine should close which channel.

References: [Go blog — pipelines and cancellation](https://go.dev/blog/pipelines).
