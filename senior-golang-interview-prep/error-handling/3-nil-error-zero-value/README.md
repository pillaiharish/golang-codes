# Q3. Understand the zero value of error

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../2-divide-return-error/README.md) · [Next question](../4-new-vs-errorf/README.md)

**Level:** Starter | **Concepts:** nil interfaces, successful returns.

## Your task — code first

1. Declare var err error and check it against nil.
2. Return nil from a successful function.
3. Explain why an interface is nil only when both its dynamic type and value are absent.

## Failure cases and acceptance tests

- [ ] Zero-value error compares equal to nil.
- [ ] A successful function returns nil.

## Explain before revealing

- What is the expected behavior on success and failure?
- Which concrete error value or type should a caller be able to inspect?
- Where should this error be handled, wrapped, logged, or retried?
- What is one implementation that looks correct but loses information?

## Practice workflow

1. Implement a Go solution without opening the answer.
2. Write a failure-case test first and then the success case.
3. Compare with [the worked solution](solution/README.md); refactor your attempt.
4. Run tests and static checks, then record evidence in [revision notes](../NOTES.md).

**Mastery is not implied by the reference code.** Mark this exercise complete only after recreating and explaining it.

