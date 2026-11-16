# Mastery system — problem-driven Go recall

[LeetCode index](../README.md) · [Existing foundations](README.md) · [Layer 1/2 course](../../god-is-great-prep/README.md) · [Card template](MASTERY-CARD-TEMPLATE.md)

## The purpose

**One canonical home per concept.** A LeetCode problem owns its solution and trace; a foundation lesson owns the reusable Go mental model, progressively expanded with Layer 2 operations, Layer 3 traps, Layer 4 patterns, and later design/optimization. Cross-link, **do not clone problems or rewrite the learner's original code**.

**Second completed pathway:** [Array](01-arrays-and-slices/ARRAY-MASTERY.md) → [Slice](01-arrays-and-slices/SLICE-MASTERY.md) → [memory traps](01-arrays-and-slices/TRAPS-AND-MEMORY.md) → [existing #238](../238.%20Product%20of%20Array%20Except%20Self/README.md) and [#41](../41.%20First%20Missing%20Positive/README.md). No existing LeetCode problem is cloned.

**First completed pathway:** [Queue](09-queues-bfs-and-grids/QUEUE-MASTERY.md) → [Grid](09-queues-bfs-and-grids/GRID-MASTERY.md) → [BFS](09-queues-bfs-and-grids/BFS-MASTERY.md) → [existing #1293 problem](../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md).

## Two axes

- **Topic:** built-in data types (array, slice, map, string), abstract structures (queue, stack, set), concrete structures (list, heap, tree, trie, graph, ring), algorithms (BFS, DFS, binary search), and patterns (sliding window, prefix sum).
- **Depth:** Layer 2 = initialization/zero/defaults/operations/costs; Layer 3 = copy, aliasing, traps, memory correctness; Layer 4 = proven reusable building blocks; later = LeetCode transfer, Go API design, performance profiling.

**Memory semantics early; memory optimization late.** Explain storage sharing and pointer retention before attempting performance tuning.

## Three reading modes

1. **60-second lookup:** recognize use case, core invariant, basic syntax, one bug.
2. **10-minute refresh:** lifecycle table, trace, tests, traps and comparisons.
3. **Rebuild from memory:** close notes; implement, run tests, explain tradeoffs, solve a related problem.

### Problem → prerequisite → problem loop

1. Attempt the existing problem **before** reading its guide.
2. Write down the exact failure: e.g., "cannot write dequeue", "mixed up grid bounds", "visited dimension wrong".
3. Follow links to a foundation's quick recall, then the corresponding tested practice package.
4. Close the reference and implement from a blank file (do not simply retype while looking).
5. Run tests and edge cases; explain the invariant aloud.
6. Return to the **same existing problem folder**, code independently, and record the outcome in [REVISION-LOG.md](../REVISION-LOG.md).
7. Revisit at roughly days 1, 3, 7, 14, 30, sooner on a failed attempt; apply the block in a different problem.

## Current concept map

| Concept | Canonical home | Evidence and problems |
|---|---|---|
| Queue | [Queue mastery](09-queues-bfs-and-grids/QUEUE-MASTERY.md) | Learner's original [dynamic queue](09-queues-bfs-and-grids/practice-01-dynamic-queue/README.md); [head-index progression](09-queues-bfs-and-grids/practice-03-head-index-queue/README.md); #102, #1293 |
| Grid indexing | [Grid mastery](09-queues-bfs-and-grids/GRID-MASTERY.md) | [Four-neighbor practice](09-queues-bfs-and-grids/practice-04-grid-neighbors/README.md); #200, #1293 |
| BFS | [BFS mastery](09-queues-bfs-and-grids/BFS-MASTERY.md) | Learner's [adjacency-list BFS](09-queues-bfs-and-grids/practice-02-graph-bfs/README.md); [plain-grid BFS](09-queues-bfs-and-grids/practice-05-grid-bfs/README.md); #102, #200, #1293 |
| Array / slice | [Array mastery](01-arrays-and-slices/ARRAY-MASTERY.md), [Slice mastery](01-arrays-and-slices/SLICE-MASTERY.md), [traps](01-arrays-and-slices/TRAPS-AND-MEMORY.md) | [Runnables](01-arrays-and-slices/practice-02-slice-operations/README.md), #36, #238, #41, #1293; full concept cards populated |
| Map / set | [Existing lesson 03](03-maps-sets-frequency-and-grouping/README.md) | #1, #36, #128 (full card extension later) |
| Text / runes | [Existing lesson 04](04-strings-bytes-runes-and-parsing/README.md) | #271, #242 (full card extension later) |
| Stack / recursion | [Existing lesson 10](10-stacks-recursion-and-tree-traversal/README.md) | #94 (full card extension later) |
| Ring / deque | [Existing lesson 12](12-circular-buffers-and-deques/README.md) | #641, #239 (full card extension later) |

"Later" means **no mastery claim**. Existing lessons remain the reference, rather than generating empty duplicate folders.

## Evidence of learning (not a checked README)

- **Recognize:** name the prerequisite and say why.
- **Reproduce:** implement with the reference and test behavior.
- **Recall:** implement again without reference, including empty/boundary cases.
- **Transfer:** apply in a different problem or representation.
- **Defend:** explain invariants, complexity, tradeoffs and a subtle failure.

Only mark stages reached in actual independent attempts. Suggested entry format in the existing revision log:

`YYYY-MM-DD | Concept: queue | Cold? yes/no | Tests? | Gap: head reset | Next: YYYY-MM-DD`

## PR acceptance gate

- Existing learner implementations and existing LeetCode solutions are **preserved**.
- All runnable examples compile, tests cover success and failures, and commands are documented.
- A concept page provides quick recall + lifecycle + traps + code links + related problems + memory exercises.
- References are bidirectional. No disconnected Layer 3 trap notebook.
- Layer 1/2 material is linked, not mass-copied. Every intentionally broken snippet is labeled.
- Never claim the user has mastered a topic merely because the code's tests pass.

**Next extensions, when problems reveal gaps:** slice, map, string, stack, linked list, then heap/tree/graph; add design and optimization last.
