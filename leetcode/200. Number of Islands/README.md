# 200. Number of Islands


The goal of the problem is to count the number of islands in a given 2D grid, where '1' represents land and '0' represents water. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.


## Explanation:
BFS: The bfs function explores all connected '1's (land) starting from a given point. It uses a queue to perform a breadth-first traversal, marking each visited cell by setting it to '0' (water).


## Main Algorithm:
Loop through the entire grid.
For every unvisited '1', it starts a BFS and increments the island count.
The BFS explores the entire connected component (island) and marks all its cells as visited.


## Time Complexity:
O(M × N), where M is the number of rows and N is the number of columns. Each cell is visited once.


## Space Complexity:
O(min(M, N)) for the queue in the BFS, because in the worst case, the width of the grid could be fully populated with land.

## Example:
### Given grid:

[  ['1', '1', '1', '1', '0'],
  ['1', '1', '0', '1', '0'],
  ['1', '1', '0', '0', '0'],
  ['0', '0', '0', '0', '0']
]


### Output
Number of islands: 1

