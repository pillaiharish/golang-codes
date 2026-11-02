# Q9. Use directional channels

[Concurrency index](../README.md) · [Q7: producer-consumer](../7-producer-consumer/README.md) · [Revision notes](../NOTES.md)

**Level 2 — channel fundamentals.**

## Task

Write one function that can only send integers and another that can only receive them. Make the compiler reject an accidental send from the consumer.

## Reference answer

```go
func produce(out chan<- int) {
    defer close(out)
    for i := 1; i <= 3; i++ { out <- i }
}

func consume(in <-chan int) {
    for v := range in { fmt.Println(v) }
}

// In main:
ch := make(chan int)
go produce(ch)
consume(ch)
```

`chan<- int` is send-only; `<-chan int` is receive-only. A bidirectional channel can be passed to a function accepting either direction, but a directional channel cannot be used for the opposite operation. The producer's send-only parameter can still be closed by the producer.

## What-if scenarios

| Question | Answer |
|---|---|
| What if `consume` attempts `in <- 1`? | Compile-time error: cannot send on a receive-only channel. |
| What if `produce` attempts `<-out`? | Compile-time error: cannot receive from a send-only channel. |
| Do directional types prevent all send/close races? | No. They restrict API operations, not the runtime lifetime contract. |
| Who should close an output channel? | The component that can establish no future sends, often its producer/coordinator. |

**Recall:** Explain why interface/API design can make concurrency contracts clearer even when the runtime behavior of the underlying channel is unchanged.

References: [Go specification — channel types](https://go.dev/ref/spec#Channel_types).
