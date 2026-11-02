# Q16. Create and fix a data race

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../15-select-default-nonblocking/README.md) · [Next question](../17-protect-shared-map/README.md)

**Level 4 — shared state and synchronization.**

## Your task

Start 1,000 goroutines that each increment one shared counter. Observe the incorrect implementation using `go run -race main.go` in your own experiment. Fix it first with a mutex and then using atomics.

## Scenario questions — answer before opening the solution

1. Why doesn't `wg.Wait()` alone protect `counter++`?
2. When use atomics versus mutex?
3. Can you read `counter` while workers update?
4. Does `counter++` always produce a wrong number without a mutex?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q16 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
