# Q9. Use directional channels

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../8-closed-channel-behavior/README.md) · [Next question](../10-channel-deadlocks-and-lifecycle/README.md)

**Level 2 — channel fundamentals.**

## Your task

Write one function that can only send integers and another that can only receive them. Make the compiler reject an accidental send from the consumer.

## Scenario questions — answer before opening the solution

1. What if `consume` attempts `in <- 1`?
2. What if `produce` attempts `<-out`?
3. Do directional types prevent all send/close races?
4. Who should close an output channel?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q9 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
