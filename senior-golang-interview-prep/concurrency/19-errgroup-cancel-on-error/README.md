# Q19. Run concurrent tasks with `errgroup`

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../18-singleton-sync-once/README.md) · [Next question](../20-semaphore-limit-concurrency/README.md)

**Level 4 — shared state and synchronization.**

## Your task

Run independent tasks concurrently. If a task returns an error, cancel sibling work and return the first non-nil error after all functions exit. Use `golang.org/x/sync/errgroup`.

## Scenario questions — answer before opening the solution

1. How is this different from `sync.WaitGroup`?
2. Can `g.Wait()` hang after one task fails?
3. What if you want at most N tasks executing?
4. Does the first error guarantee all tasks had no side effects?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q19 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
