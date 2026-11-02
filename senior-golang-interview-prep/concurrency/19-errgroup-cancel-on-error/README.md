# Q19. Run concurrent tasks with `errgroup`

[Concurrency index](../README.md) · [Q13: context](../13-context-cancellation/README.md) · [Revision notes](../NOTES.md)

**Level 4 — shared state and synchronization.**

## Task

Run independent tasks concurrently. If a task returns an error, cancel sibling work and return the first non-nil error after all functions exit. Use `golang.org/x/sync/errgroup`.

## Reference answer

```go
g, ctx := errgroup.WithContext(parentCtx)
g.Go(func() error { return task1(ctx) })
g.Go(func() error { return task2(ctx) })
if err := g.Wait(); err != nil { return err }
return nil
```

This snippet belongs inside a function returning `error` and requires the external `golang.org/x/sync/errgroup` module. `WithContext` cancels its derived context on the first non-nil task error (or after `Wait` returns). `Wait` still waits for **all** started functions; it does not kill noncooperative work.

## Scenario Q&A

| Question | Answer |
|---|---|
| How is this different from `sync.WaitGroup`? | It adds error collection and optional sibling cancellation through context. |
| Can `g.Wait()` hang after one task fails? | Yes, if another task never returns or ignores cancellation. |
| What if you want at most N tasks executing? | Use `g.SetLimit(N)` or bound admissions before task creation. |
| Does the first error guarantee all tasks had no side effects? | No. Other tasks may have already completed or partially performed work. |

**Recall:** Explain fail-fast signaling versus wait-for-all joining. Design partial-result/error reporting if the caller needs more than the first error.

Reference: [`golang.org/x/sync/errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup).
