package main

import (
	"container/list"
	"fmt"
)

type SnakeGame struct {
	width, height int
	food          [][]int
	score         int
	snake         *list.List
	snakeSet      map[[2]int]bool
	directions    map[string][2]int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	game := SnakeGame{width: width, height: height, food: food, score: 0,
		snake:      list.New(),
		snakeSet:   make(map[[2]int]bool),
		directions: map[string][2]int{"L": {0, -1}, "R": {0, 1}, "U": {-1, 0}, "D": {1, 0}},
	}
	game.snake.PushBack([2]int{0, 0})
	game.snakeSet[[2]int{0, 0}] = true
	return game
}

func (this *SnakeGame) Move(direction string) int {

	if _, ok := this.directions[direction]; !ok {
		return -1
	}
	head := this.snake.Back().Value.([2]int)
	move := this.directions[direction]
	newHead := [2]int{head[0] + move[0], head[1] + move[1]}

	// touches outside grid
	if newHead[0] < 0 || newHead[0] >= this.height || newHead[1] < 0 || newHead[1] >= this.width {
		return -1
	}
	// Check if the new head position hits the snake body (excluding the tail which will move)
	if this.snakeSet[newHead] && !(newHead == this.snake.Front().Value.([2]int)) {
		return -1
	}
	//touches food
	if this.score < len(this.food) && newHead[0] == this.food[this.score][0] && newHead[1] == this.food[this.score][1] {
		this.score++
		// Don't remove the tail from snakeSet since we grow the snake
	} else {
		// Normal move, remove tail
		tail := this.snake.Remove(this.snake.Front()).([2]int) //Type Assertion .( [2]int):
		delete(this.snakeSet, tail)
	}
	this.snake.PushBack(newHead)
	this.snakeSet[newHead] = true
	return this.score
}

func main() {
	// Test case
	game := Constructor(3, 2, [][]int{{1, 2}, {0, 1}})
	fmt.Println(game.Move("R")) // Output: 0
	fmt.Println(game.Move("D")) // Output: 0
	fmt.Println(game.Move("R")) // Output: 1
	fmt.Println(game.Move("U")) // Output: 1
	fmt.Println(game.Move("L")) // Output: 2
	fmt.Println(game.Move("U")) // Output: -1
}
// Output
// harish $ go run main.go 
// 0
// 0
// 1
// 1
// 2
// -1