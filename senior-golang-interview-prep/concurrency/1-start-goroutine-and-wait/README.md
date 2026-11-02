# Q1. Start a goroutine and make `main` wait

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Next question](../2-start-five-goroutines-and-wait/README.md)

**Level:** 1 — fundamentals | **Concepts:** `go`, goroutine lifetime, `sync.WaitGroup`, `Add`, `Done`, `Wait`, `defer`.

## Your task

Write a standalone Go program that:

1. Starts **one goroutine** that prints `Hello from goroutine`.
2. Makes `main` **wait for that goroutine to finish**.
3. Only then prints `Main finished`.

First experiment with `time.Sleep(1 * time.Second)` as a deliberately fragile way to keep `main` alive; change the worker duration and notice why it cannot guarantee completion. Then replace the sleep with `sync.WaitGroup` for the final correct solution.

### Expected output

```text
Hello from goroutine
Main finished
```

Both lines must appear, in that order, on every successful run.

## Scenario questions — answer before opening the solution

1. Remove `wg.Wait()` — Will the worker always print?
2. Remove `wg.Done()` — What happens at `Wait`?
3. Move `wg.Add(1)` inside the goroutine — Is `Wait` still safe?
4. Replace `Wait` with `time.Sleep` — Is a long delay enough?
5. Use a named worker function — Must the worker be anonymous?
6. Worker panics — Does `defer wg.Done()` suppress a panic?
7. Use a completion channel — Can it solve this one-worker problem?

## Cold-recall prompts

- Why is the WaitGroup's zero value ready to use?
- Why is `Add(1)` deliberately outside the goroutine?
- Does `WaitGroup` guarantee the goroutine starts immediately?
- What happens if you call `Done` twice after only one `Add(1)`?
- How would you implement the same contract with a `done` channel?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q1 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
