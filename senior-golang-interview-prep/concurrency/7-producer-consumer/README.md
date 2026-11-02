# Q7. Build a simple producer-consumer

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../6-buffered-channel-capacity/README.md) · [Next question](../8-closed-channel-behavior/README.md)

**Level 2 — channel fundamentals.**

## Your task

One producer sends `1..10` into a channel, closes it after the final send, and one consumer prints all received values. Explain who owns the close operation and when the consumer stops.

## Scenario questions — answer before opening the solution

1. What if the producer never closes `ch`?
2. What if the consumer closes while the producer is sending?
3. What if there are multiple producers?
4. What if the consumer stops after three values?
5. Does a channel close imply all processing is finished?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q7 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
