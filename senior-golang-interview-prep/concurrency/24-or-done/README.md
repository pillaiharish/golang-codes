# Q24. Implement `orDone`

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../23-cancellable-pipeline/README.md) · [Next question](../25-tee-broadcast-two-outputs/README.md)

**Level 5 — core concurrency patterns.**

## Your task

Implement `func orDone(ctx context.Context, ch <-chan int) <-chan int`: relay values, exit if the source closes or context cancels, and close the returned channel exactly once.

## Scenario questions — answer before opening the solution

1. What if `ctx` is canceled before the first value?
2. What if `ch` closes normally?
3. What if both a value and cancellation are ready?
4. Does this function shut down the original source goroutine?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q24 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
