# Q26. Implement a `bridge` for a channel of channels

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../25-tee-broadcast-two-outputs/README.md) · [Next question](../27-rate-limiter/README.md)

**Level 5 — core concurrency patterns.**

## Your task

Flatten `streams <-chan <-chan int` into one `<-chan int`. If the outer channel yields streams `[1,2,3]`, `[4,5]`, `[6]`, output should be `1,2,3,4,5,6` for a **sequential bridge**.

## Scenario questions — answer before opening the solution

1. What if an inner channel never closes?
2. What if the outer channel produces a nil inner channel?
3. What if the consumer stops reading output?
4. Who closes the flattened output?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q26 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
