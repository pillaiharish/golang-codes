# Tree structure, recursion frames, and traversal state

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A binary tree node has zero, one, or two children. Depth measures a root-to-node path; subtree height measures descent to a deepest leaf. Some contracts count edges while others count nodes. A binary search tree imposes constraints on whole subtrees, not merely immediate children.

A recursive function needs a precise meaning for its return value and a base case for an empty subtree. Each call owns local state, but captured/global variables may be shared. A chain of nodes can make recursion depth linear even though a balanced tree has logarithmic height. Preorder, inorder, and postorder specify when a node is processed relative to its children; learn these terms without assuming a particular traversal solves a given task.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
type Branch struct { Value int; Left, Right *Branch }
root := &Branch{Value: 8, Left: &Branch{Value: 3}}
fmt.Println(root.Value, root.Left.Value, root.Right == nil)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **8 3 true**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Draw one tree and list preorder/inorder/postorder by hand. State the depth of each node in both edge and node conventions.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Node-versus-edge off-by-one errors; applying a local ordering check to a global BST contract; treating a tree as balanced without a guarantee.

## Questions that use this building block

- [#226 Invert Binary Tree](../../226.%20Invert%20Binary%20Tree/README.md)
- [#104 Maximum Depth of Binary Tree](../../104.%20Maximum%20Depth%20of%20Binary%20Tree/README.md)
- [#543 Diameter of Binary Tree](../../543.%20Diameter%20of%20Binary%20Tree/README.md)
- [#110 Balanced Binary Tree](../../110.%20Balanced%20Binary%20Tree/README.md)
- [#100 Same Tree](../../100.%20Same%20Tree/README.md)
- [#572 Subtree of Another Tree](../../572.%20Subtree%20of%20Another%20Tree/README.md)
- [#235 Lowest Common Ancestor of a Binary Search Tree](../../235.%20Lowest%20Common%20Ancestor%20of%20a%20Binary%20Search%20Tree/README.md)
- [#102 Binary Tree Level Order Traversal](../../102.%20Binary%20Tree%20Level%20Order%20Traversal/README.md)
- [#199 Binary Tree Right Side View](../../199.%20Binary%20Tree%20Right%20Side%20View/README.md)
- [#1448 Count Good Nodes in Binary Tree](../../1448.%20Count%20Good%20Nodes%20in%20Binary%20Tree/README.md)
- [#98 Validate Binary Search Tree](../../98.%20Validate%20Binary%20Search%20Tree/README.md)
- [#230 Kth Smallest Element in a BST](../../230.%20Kth%20Smallest%20Element%20in%20a%20BST/README.md)
- [#105 Construct Binary Tree from Preorder and Inorder Traversal](../../105.%20Construct%20Binary%20Tree%20from%20Preorder%20and%20Inorder%20Traversal/README.md)
- [#124 Binary Tree Maximum Path Sum](../../124.%20Binary%20Tree%20Maximum%20Path%20Sum/README.md)
- [#297 Serialize and Deserialize Binary Tree](../../297.%20Serialize%20and%20Deserialize%20Binary%20Tree/README.md)
- [#78 Subsets](../../78.%20Subsets/README.md)
- [#39 Combination Sum](../../39.%20Combination%20Sum/README.md)
- [#46 Permutations](../../46.%20Permutations/README.md)
- [#90 Subsets II](../../90.%20Subsets%20II/README.md)
- [#40 Combination Sum II](../../40.%20Combination%20Sum%20II/README.md)
- [#79 Word Search](../../79.%20Word%20Search/README.md)
- [#131 Palindrome Partitioning](../../131.%20Palindrome%20Partitioning/README.md)
- [#17 Letter Combinations of a Phone Number](../../17.%20Letter%20Combinations%20of%20a%20Phone%20Number/README.md)
- [#51 N-Queens](../../51.%20N-Queens/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../leetcode/0-problem-solving-building-blocks/10-stacks-recursion-and-tree-traversal/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
