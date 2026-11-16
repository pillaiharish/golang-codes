# Cold recall drills — queue, grid, BFS

[Queue mastery](QUEUE-MASTERY.md) · [Grid mastery](GRID-MASTERY.md) · [BFS mastery](BFS-MASTERY.md) · [Problem #1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/REVISION-CARD.md) · [Revision log](../../REVISION-LOG.md)

**Do not look at implementations while attempting a drill.** Read the short task, predict outputs/trace, write code, then compare and run tests. A passing test today is not proof of recall next month.

| Checkpoint | Blank-file exercise | Proof of understanding |
|---|---|---|
| D0 | Write a FIFO queue from scratch | Empty, FIFO, drain/reuse cases pass |
| D1 | Write row/col 4-direction bounds from memory | Edge and ragged rows handled |
| D3 | Implement BFS over an adjacency-list graph | Cycles safe; explain mark-on-enqueue |
| D7 | Implement shortest path in a **plain** zero/one grid | Correct level boundaries and unreachable case |
| D14 | Solve the existing #1293 without hints | State includes remaining budget; shortest steps justified |
| D30 | Change problem form: tree levels, multiple sources, or obstacle budget | Identify needed state and queue correctly |

When a recall fails, revisit only that mastery section, then retry from a blank editor. If the same code is repeatedly copied, hide the card earlier and change input names/shapes to test transfer.

## Targeted failure probes

1. **Zero value:** What does `var q Queue[int]` contain? Why does `append` work?
2. **FIFO:** How do `head`, `len(queue)`, and logical length differ?
3. **Copy/retention:** Does `queue=queue[1:]` shift all elements? Can it retain memory?
4. **Grid:** What breaks with `grid[nr][nc]` if row isn't checked first?
5. **BFS:** Why must `levelEnd` be captured before appending?
6. **State:** At the same cell, are `remaining=0` and `remaining=2` equivalent for #1293?
7. **Representation:** Why does a grid not require a separately allocated adjacency list?
8. **Complexity:** Where does `(k+1)` appear in #1293's upper bound?

**Leave actual progress blank** until each attempt. Use the repository's revision log rather than committing imagined confidence or dates.
