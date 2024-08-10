# Circular Deque Implementation

## Explanation

### Circular Array
The array `data` is used to store the elements, with a size of `k + 1` to differentiate between full and empty states. The `front` and `rear` pointers manage the insertion and deletion operations.

### Insert Operations
- **InsertFront**: Moves the `front` pointer backward (circularly) and inserts the element at the new front position.
- **InsertLast**: Inserts the element at the `rear` position and moves the `rear` pointer forward (circularly).

### Delete Operations
- **DeleteFront**: Moves the `front` pointer forward, effectively removing the element at the front.
- **DeleteLast**: Moves the `rear` pointer backward, removing the element at the rear.

### Boundary Conditions
- **IsFull**: Checks if the deque is full by comparing if the next `rear` position is the `front`.
- **IsEmpty**: Checks if the deque is empty by comparing if `front` and `rear` are equal.

### Edge Cases
The circular nature ensures that when the deque reaches the end of the array, it wraps around to the beginning.

## Time Complexity
All operations (`InsertFront`, `InsertLast`, `DeleteFront`, `DeleteLast`, `GetFront`, `GetRear`, `IsEmpty`, `IsFull`) run in \(O(1)\) time.

This implementation efficiently handles the operations of a circular deque and adheres to the problem constraints.

