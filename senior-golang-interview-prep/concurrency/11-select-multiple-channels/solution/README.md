# Q11. Use `select` with multiple channels — worked solution

[Back to Q11 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q12: timeout](../../12-timeout-with-time-after/README.md) · [Revision notes](../../NOTES.md)

**Level 3 — select, timeouts, and cancellation.**

## Task

Start two senders with different delays and receive whichever channel becomes ready first. Then reason about simultaneous readiness and the sender you did not select.

## Reference answer

```go
ch1 := make(chan string, 1)
ch2 := make(chan string, 1)
go func() { time.Sleep(100*time.Millisecond); ch1 <- "one" }()
go func() { time.Sleep(200*time.Millisecond); ch2 <- "two" }()
select {
case v := <-ch1:
    fmt.Println(v)
case v := <-ch2:
    fmt.Println(v)
}
```

`select` chooses one ready communication case. If multiple cases can proceed, Go makes a pseudo-random choice rather than always choosing the first source-code case. Delays make an outcome likely, not contractually guaranteed. Buffering both one-time result channels avoids leaving the unchosen sender blocked when `main` moves on.

## Scenario Q&A

| Question | Answer |
|---|---|
| If both cases are ready, is the first case preferred? | No. One ready case is selected pseudo-randomly. |
| If neither case is ready? | `select` blocks unless a `default` case exists. |
| If one channel is nil? | That communication case is disabled, not ready. |
| If a selected channel is closed? | Receive proceeds immediately and may return the zero value; use the two-value form to detect close. |
| If the other sender blocks forever? | The application may leak that goroutine; design cancellation, buffering, or draining. |

**Recall:** Explain why a single `select` executes only one case; modify the program to receive and print **both** results.

Reference: [Go specification — select statements](https://go.dev/ref/spec#Select_statements).
