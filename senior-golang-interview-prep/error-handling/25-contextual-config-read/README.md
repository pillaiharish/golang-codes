# Q25. Preserve cause and category while reading config
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../24-extract-db-error/README.md) · [Next](../26-layered-error-messages/README.md)

**Level:** Intermediate | **Concepts:** fs.ErrNotExist, invalid content, layers.

## Your task — code first

1. Implement readConfig(path string) error.
2. Classify missing file and invalid contents separately.
3. Include path context while preserving the original filesystem error.

## Failure cases and acceptance tests

- [ ] Missing file matches fs.ErrNotExist.
- [ ] Invalid file matches ErrInvalidConfig.

## Explain before revealing

- What is the behavior on both success and failure?
- Which error identity or typed details can the caller inspect?
- Where does wrapping, logging, cancellation or retry policy belong?
- Which apparently reasonable implementation could conceal a failure?

## Practice workflow

1. Implement independently without opening the answer.
2. Write tests for expected failures and successes.
3. Compare with the [worked solution](solution/README.md), then refactor.
4. Run tests and static checks; record evidence in [revision notes](../NOTES.md).

**Mastery is earned by reproducing the solution and explaining its trade-offs**, not by reading the reference.

