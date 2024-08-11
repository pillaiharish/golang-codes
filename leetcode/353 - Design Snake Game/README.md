# Snake Game Implementation

## Key Components

### Snake Representation
- The snake is represented using a deque (double-ended queue), where the front of the deque represents the tail and the back represents the head.
- A set (`snakeSet`) is used to track the positions occupied by the snake, facilitating quick checks for collisions with itself.

### Direction Map
- The `directions` map defines the movement delta for each direction:
  - 'U' (Up): (-1, 0)
  - 'D' (Down): (1, 0)
  - 'L' (Left): (0, -1)
  - 'R' (Right): (0, 1)

### Move Method
1. **Boundary and Self Collision Check**: Before moving, the game checks if the new head position collides with the wall or the snake itself.
2. **Food Consumption**: If the snake eats food, the score increases, and the snake grows (the tail isn't removed).
3. **Regular Movement**: The snake moves by adding a new head position and removing the tail, unless the snake is growing (after eating food).

### Game Over
- The game returns -1 if the snake runs into a wall or itself, indicating the game is over.


## Full Explanation of the Line
``` go
tail := this.snake.Remove(this.snake.Front()).([2]int)
```

## Step 1: this.snake.Front():

The method Front() retrieves the first element of the deque, which corresponds to the snake's current tail position.

## Step 2: this.snake.Remove(this.snake.Front()):

The Remove() method is then called with the element obtained from Front(), which removes this element from the deque. This reflects the snake "moving forward" by effectively removing its tail from the previous position.

## Step 3: Type Assertion .( [2]int):

The Remove() method returns the value of the removed element. Since this value is expected to be of type [2]int (a 2-element integer array representing the x and y coordinates), we assert its type to [2]int.

## Step 4: Assignment to tail:

Finally, the result of the type assertion is assigned to the variable tail. Now, tail holds the coordinates of the removed tail position.

## Why is This Done?
In the context of the snake game:

Managing Snake Movement: When the snake moves forward without eating food, the tail’s previous position is no longer part of the snake’s body, so it is removed from the deque.
Tracking the Tail: The tail variable allows you to keep track of the position that was just removed, which might be useful if you need to update the game grid or check for collisions.
This line of code is a concise and efficient way to manage the snake's movement by removing its tail and tracking its last position.


In the context of implementing the Snake game, the decision to represent the head of the snake at the back of the deque and the tail at the front is primarily based on how we typically add and remove elements during the gameplay:

## Key Concepts:
### Head Movement:

The head of the snake is the part that moves forward with each step in the game. Every time the snake moves, a new position is added to the head, which is why the head is at the back of the deque.

### Tail Movement:

If the snake hasn't just eaten food, the tail moves forward to simulate the snake moving without growing. This involves removing the tail's current position, which is naturally done by removing the front element from the deque.

## Detailed Explanation:

### Deque Operations:

Back of Deque (Head): When the snake moves, we append the new head position to the back of the deque. This reflects the snake's growth or movement toward a new position.
Front of Deque (Tail): If the snake doesn't eat food, the tail (which is at the front) moves forward, and we remove the front element of the deque.

### Growth and Shrinking:

When the snake eats food, the new position of the head is added, but the tail remains in place, effectively growing the snake. The head continues to move forward, and no element is removed from the deque.
When the snake moves without eating food, the tail position (front of the deque) is removed to reflect the snake's movement while maintaining its length.

## Why This Representation?

### Efficient Insertions and Deletions:

A deque allows for O(1) time complexity for both adding elements to the back and removing elements from the front, which matches the natural operations needed for simulating the snake's movement.

### Clarity in Operations:

Keeping the head at the back and the tail at the front provides a clear and direct way to manage the snake's behavior. It simplifies operations like adding a new head position and removing the tail position, which corresponds directly to the game's logic.

## Summary:
The decision to represent the head of the snake at the back of the deque and the tail at the front is driven by the need for efficient operations that correspond to how the snake moves and grows in the game. It ensures that each move, whether it involves growth or simple movement, is handled in constant time and aligns with the game's mechanics.