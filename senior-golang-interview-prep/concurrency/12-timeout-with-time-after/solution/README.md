# Q12. Implement a timeout with `time.After` — worked solution

[Back to Q12 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q13: context cancellation](../../13-context-cancellation/README.md) · [Revision notes](../../NOTES.md)

**Level 3 — select, timeouts, and cancellation.**

## Task

Start a simulated slow operation. Print `operation timed out` if no result arrives within 500 ms.

## Reference answer

```go
resultCh := make(chan string, 1)
go func() {
    time.Sleep(time.Second)
    resultCh <- "finished"
}()

select {
case result := <-resultCh:
    fmt.Println(result)
case <-time.After(500 * time.Millisecond):
    fmt.Println("operation timed out")
}
```

The buffered result channel lets the slow worker finish its send even after the receiver chose timeout. **A timeout in the caller does not stop the underlying work**; here the worker still sleeps for a second. For cooperative real work, pass a `context.Context` so the callee can abandon work early.

## Scenario Q&A

| Question | Answer |
|---|---|
| What if the result and timer are both ready? | `select` may choose either ready case; define boundary semantics if strict timeout behavior matters. |
| Does `time.After` kill the worker? | No. It only provides a timer channel to the waiter. |
| Why use `context.WithTimeout` in a server? | Deadlines/cancellation can propagate into HTTP, DB and child operations that honor context. |
| What if the worker sends on an unbuffered result channel after the timeout? | It can block forever without a receiver. Add cancellation, a bounded buffer, or a receiver. |

**Recall:** Convert this into `context.WithTimeout` and ensure you call the returned cancel function. Define what an outstanding operation means after timeout.

References: [`time.After`](https://pkg.go.dev/time#After), [`context.WithTimeout`](https://pkg.go.dev/context#WithTimeout).
