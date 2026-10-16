# 94 — Binary Tree Inorder Traversal

## Before you solve: prerequisite basics

Study these in sequence and type the tiny Go example in each before reopening this solution:

1. [01 arrays and slices](../0-problem-solving-building-blocks/01-arrays-and-slices/README.md)
2. [10 stacks recursion and tree traversal](../0-problem-solving-building-blocks/10-stacks-recursion-and-tree-traversal/README.md)
3. [11 structs pointers methods and linked lists](../0-problem-solving-building-blocks/11-structs-pointers-methods-and-linked-lists/README.md)
4. [17 complexity amortization and input contracts](../0-problem-solving-building-blocks/17-complexity-amortization-and-input-contracts/README.md)

[Full mapping for all 23 LeetCode problems](../0-problem-solving-building-blocks/README.md#b-problem--exact-prerequisites-all-23-problem-folders) · [Practice log](../REVISION-LOG.md)

[← Index](../README.md) · [Building block: recursion and stack](../0-problem-solving-building-blocks/README.md#7-trees-recursion-gives-the-traversal-order) · [Go solutions](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md) · [LeetCode](https://leetcode.com/problems/binary-tree-inorder-traversal/)

## Recognition

**Clue:** "return the inorder traversal of a binary tree." Remember a three-part order, not a specific implementation: **Left → Node → Right (LNR).**

**Mental trigger:** **Go left until nil; visit node; go right.**

## Understanding before coding

A tree node contains a value and pointers `Left` and `Right`. The inorder of a subtree is the inorder of its left child, followed by its own value, followed by the inorder of its right child. Recursion is an exact translation of that definition.

The key Go concept is **how all recursive calls append to the same result slice**. The submitted version passes a pointer to the slice header: `output *[]int`; it appends using `*output = append(*output, root.Val)`. The pointer matters because append can replace the slice header with one pointing to a new backing array.

## Invariant

When `inorder(node, &result)` finishes, **all nodes of that subtree** have been appended in Left-Node-Right order after any previously accumulated result.

## Recipe 1 — your pointer-to-slice method

1. Create `result := []int{}`.
2. Define `inorder(node, &result)`.
3. If node is nil, return.
4. Recurse into left.
5. Append current node through the pointer: `*out = append(*out, node.Val)`.
6. Recurse into right.

### Trace: root `1 → right 2 → left 3`

```text
    1
     \
      2
     /
    3
```

| Action | Result |
|---|---|
| inorder(left of 1) = nil | [] |
| visit 1 | [1] |
| inorder(left of 2) → visit 3 | [1,3] |
| visit 2 | [1,3,2] |

## Alternative A — closure (same recursion, simpler interface)

A nested `dfs := func(...)` captures `result`. No pointer-to-slice parameter is necessary. Recursive closures need a `var walk func(*TreeNode)` declaration before assigning the function literal. See `inorderClosure`.

## Alternative B — iterative stack (when recursion worries you)

1. While current exists, push it; move left.
2. If left is exhausted, pop the stack's last node.
3. Append its value.
4. Move to its right child.
5. Stop when `current == nil` and stack is empty.

The stack holds ancestors waiting for their node visit. See `inorderIterative`.

## Complexity

All methods visit n nodes once: **O(n) time**. Output requires **O(n)** space. Additional recursion depth or explicit stack: **O(h)**, where h is tree height (balanced O(log n), worst-case skew O(n)).

## Failure traps

- Reversing the order results in preorder or postorder.
- Forgetting nil base case.
- Appending to a *copy* of a slice header that never reaches the caller.
- Returning a nil slice vs empty slice: both represent no elements, but tests must make the intended API contract explicit.

## 30-second spoken answer

> Inorder is left, node, right. I recursively visit the left child, append the current value to one shared result, then visit the right child. The Go result can be captured by a closure or passed by slice-header pointer. Each node is processed once: O(n) time, with O(h) stack space besides the O(n) output. An iterative version simulates recursion with an explicit stack.

## Cold recall

1. What order gives sorted output for a binary search tree?
2. Why does `*out = append(*out, v)` update the caller?
3. What does the explicit stack contain during iterative traversal?
4. When can recursion depth reach O(n)?
5. Recode the **pointer**, **closure**, and **stack** versions without looking.

**Personal review:** Confidence `— / 5` · Last reviewed `—` · Next review `—`.
