# Q18. Initialize a singleton using `sync.Once`

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../17-protect-shared-map/README.md) · [Next question](../19-errgroup-cancel-on-error/README.md)

**Level 4 — shared state and synchronization.**

## Your task

Let many goroutines call `getInstance()` concurrently, but create the underlying configuration object exactly once.

## Scenario questions — answer before opening the solution

1. What if the initializer panics?
2. Can `f` recursively call `once.Do` on the same `Once`?
3. Is the initialized object's mutable state automatically safe?
4. What if initialization can return an error?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q18 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
