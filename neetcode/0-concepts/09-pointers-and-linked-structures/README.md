# Pointers, nodes, identity, and linked structures

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A node pointer represents object identity. Two nodes may have equal values while being different objects. Changing a pointer variable redirects that variable; changing p.Next alters the pointed-to node. A nil check must precede dereferencing. If rewiring can lose access to a remainder, preserve the required reference before mutation.

Shallow copies of structs with pointer fields retain references to the same other nodes. A deep-copy contract demands distinct objects with equivalent relationships. Cycles invalidate traversals that assume eventual nil. Dummy nodes can simplify boundary bookkeeping, but their use and lifetime must be justified in your own implementation. The local starter supplies node types because LeetCode normally provides them.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
type Cell struct { Value int; Next *Cell }
a := &Cell{Value: 4}
b := &Cell{Value: 4}
a.Next = b
fmt.Println(a == b, a.Value == b.Value, a.Next == b)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **false true true**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Draw three nodes and two aliases. Change a field, then redirect one alias; mark which observations change.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Comparing values when identity matters; losing a link during mutation; copying random pointers into the original structure; infinite traversal of cycles.

## Questions that use this building block

- [#206 Reverse Linked List](../../206.%20Reverse%20Linked%20List/README.md)
- [#21 Merge Two Sorted Lists](../../21.%20Merge%20Two%20Sorted%20Lists/README.md)
- [#143 Reorder List](../../143.%20Reorder%20List/README.md)
- [#19 Remove Nth Node From End of List](../../19.%20Remove%20Nth%20Node%20From%20End%20of%20List/README.md)
- [#138 Copy List with Random Pointer](../../138.%20Copy%20List%20with%20Random%20Pointer/README.md)
- [#2 Add Two Numbers](../../2.%20Add%20Two%20Numbers/README.md)
- [#141 Linked List Cycle](../../141.%20Linked%20List%20Cycle/README.md)
- [#287 Find the Duplicate Number](../../287.%20Find%20the%20Duplicate%20Number/README.md)
- [#146 LRU Cache](../../146.%20LRU%20Cache/README.md)
- [#23 Merge k Sorted Lists](../../23.%20Merge%20k%20Sorted%20Lists/README.md)
- [#25 Reverse Nodes in k-Group](../../25.%20Reverse%20Nodes%20in%20k-Group/README.md)
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
- [#208 Implement Trie (Prefix Tree)](../../208.%20Implement%20Trie%20%28Prefix%20Tree%29/README.md)
- [#211 Design Add and Search Words Data Structure](../../211.%20Design%20Add%20and%20Search%20Words%20Data%20Structure/README.md)
- [#212 Word Search II](../../212.%20Word%20Search%20II/README.md)
- [#133 Clone Graph](../../133.%20Clone%20Graph/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../leetcode/0-problem-solving-building-blocks/11-structs-pointers-methods-and-linked-lists/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
