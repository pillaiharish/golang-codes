# Q1. Start a goroutine and make `main` wait

[Concurrency index](../README.md) · [Q2: five goroutines](../2-start-five-goroutines-and-wait/README.md) · [Runnable solution](main.go) · [Revision notes](../NOTES.md)

**Level:** 1 — fundamentals | **Concepts:** `go`, goroutine lifetime, `sync.WaitGroup`, `Add`, `Done`, `Wait`, `defer`.

## Interview question

Write a standalone Go program that:

1. Starts **one goroutine** that prints `Hello from goroutine`.
2. Makes `main` **wait for that goroutine to finish**.
3. Only then prints `Main finished`.

Do not use `time.Sleep` to coordinate goroutine completion. Use `sync.WaitGroup`.

### Expected output

```text
Hello from goroutine
Main finished
```

Both lines must appear, in that order, on every successful run.

## Try it first

Before opening [the solution](main.go), write the program from memory. Identify:

- **Who starts the work?** The `main` goroutine.
- **Who owns the join?** `main` calls `Wait`.
- **What signals completion?** The worker calls `Done`.
- **What prevents an early return?** A positive WaitGroup count until `Done` executes.

## Reference solution

The [standalone program](main.go) uses this essential sequence:

```go
var wg sync.WaitGroup // count begins at 0
wg.Add(1)             // register work before starting it

go func() {
    defer wg.Done()   // signal completion on return
    fmt.Println("Hello from goroutine")
}()

wg.Wait()             // do not return from main early
fmt.Println("Main finished")
```

Run from this folder:

```sh
go run main.go
go run -race main.go
```

The counter transitions `0 -> 1 -> 0`; `Wait` returns after the count reaches zero. `WaitGroup` is a *completion barrier*, not a scheduler, error collector, or cancellation mechanism.

## Scenario variations interviewers may ask

| Scenario | Question to reason through | Correct conclusion |
|---|---|---|
| Remove `wg.Wait()` | Will the worker always print? | No. Returning from `main` terminates the program; the worker may not get CPU time. |
| Remove `wg.Done()` | What happens at `Wait`? | It cannot return while the count stays 1. If nothing can make progress, the runtime detects a deadlock. |
| Move `wg.Add(1)` inside the goroutine | Is `Wait` still safe? | No. `main` could observe count 0 and return before the increment. Register work before launching it. |
| Replace `Wait` with `time.Sleep` | Is a long delay enough? | No. Timing is not a completion guarantee; it can be flaky or unnecessarily slow. |
| Use a named worker function | Must the worker be anonymous? | No. `go worker(&wg)` works too, provided completion is signaled once. |
| Worker panics | Does `defer wg.Done()` suppress a panic? | No. Deferred calls run during unwinding, but an unrecovered panic still terminates the program. |
| Use a completion channel | Can it solve this one-worker problem? | Yes: close a `done` channel in the worker and receive from it in `main`; this question specifically practices `WaitGroup`. |

## Common mistakes and why they fail

1. **`go func() { ... }` without trailing `()`** — declares a function value rather than invoking it as a goroutine.
2. **Calling `Done` too early** — `Wait` may return before the actual work finishes.
3. **Forgetting `Done`** — can block the waiter forever.
4. **Using `Add` after starting the worker** — creates a scheduling-dependent bug.
5. **Assuming `Wait` stops a goroutine** — it only waits; cancellation requires a different protocol.

## Explain it in 30 seconds

> I register one outstanding task with `wg.Add(1)` before launching the goroutine. The worker defers `wg.Done()` so its completion decrements the counter. `main` blocks in `wg.Wait()` until that count becomes zero, then prints its final line. I don't use sleep because scheduling and completion time are nondeterministic.

## Cold-recall checks

- Why is the WaitGroup's zero value ready to use?
- Why is `Add(1)` deliberately outside the goroutine?
- Does `WaitGroup` guarantee the goroutine starts immediately?
- What happens if you call `Done` twice after only one `Add(1)`?
- How would you implement the same contract with a `done` channel?

**Mastery evidence:** Type the program without copying; explain one failing execution schedule; run it locally and record the result in [NOTES.md](../NOTES.md). Having this reference solution in the repository is not a claim of completion.

## References

- [Go specification: go statements](https://go.dev/ref/spec#Go_statements)
- [`sync.WaitGroup` reference](https://pkg.go.dev/sync#WaitGroup)

