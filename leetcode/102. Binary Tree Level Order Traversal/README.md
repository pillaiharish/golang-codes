# 102. Binary Tree Level Order Traversal

https://leetcode.com/problems/binary-tree-level-order-traversal/description/

### 1) Queue: A queue is used to track nodes at the current level. Each time, all nodes at the current level are processed, and their children are added to the queue for the next level.
### 2) Result: A 2D slice is used to store the values of nodes at each level.


## Time Complexity:
O(n), where n is the number of nodes in the binary tree. Each node is visited exactly once, so the time complexity depends linearly on the number of nodes.


## Space Complexity:
O(n), where n is the number of nodes. In the worst case, all the leaf nodes are stored in the queue at the last level, which could be up to n/2 nodes. Thus, the space complexity is proportional to n.


The algorithm performs a level-order traversal using a queue, which processes each level of the tree before moving to the next, ensuring all nodes are visited.