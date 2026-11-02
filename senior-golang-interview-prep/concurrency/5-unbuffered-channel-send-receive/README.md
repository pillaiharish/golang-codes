# Q5. Send and receive on an unbuffered channel

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../4-goroutine-panic-and-recover/README.md) · [Next question](../6-buffered-channel-capacity/README.md)

**Level 2 — channel fundamentals.**

## Your task

Create `ch := make(chan int)`. Start a goroutine that sends `42`; have `main` receive and print it. Explain which operation blocks and why.

## Scenario questions — answer before opening the solution

1. What if `ch <- 42` happens in `main` immediately before `<-ch`?
2. What if the goroutine sends twice but `main` receives once?
3. Can an unbuffered channel be used purely for a signal?
4. Does a send guarantee the receiver has finished processing?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q5 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
