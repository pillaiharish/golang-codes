# Q22. Implement fan-out and fan-in

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../21-worker-pool/README.md) · [Next question](../23-cancellable-pipeline/README.md)

**Level 5 — core concurrency patterns.**

## Your task

Fan-out: multiple workers read the **same** input channel. Fan-in: merge values from several input channels into one output channel. Write `merge(channels ...<-chan int) <-chan int`.

## Scenario questions — answer before opening the solution

1. What if no inputs are given?
2. What if one input never closes?
3. What if the downstream consumer stops early?
4. Does fan-out make each input item reach every worker?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q22 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
