package main

import "fmt"

// isValidSudoku checks only the filled cells. It does not solve the puzzle.
// Time: O(9*9) for the fixed board; space: O(9*9) for three fixed tables.
func isValidSudoku(board [][]byte) bool {
	// Step 1: Track each digit independently for every row, column and 3x3 box.
	// Why: A duplicate is illegal only within one of these three groups.
	var rows, cols, boxes [9][9]bool

	// Step 2: Visit all 81 cells exactly once.
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cell := board[r][c]
			if cell == '.' {
				continue // Why: Empty positions impose no duplicate constraint.
			}

			// Step 3: Convert ASCII '1'..'9' to indexes 0..8.
			// '1'-'1' is 0, '9'-'1' is 8. The problem guarantees valid symbols.
			digit := int(cell - '1')
			// Each 3x3 box is identified by its row band and column band.
			// (r/3)*3 + c/3 maps top-left to 0 and bottom-right to 8.
			boxIndex := (r/3)*3 + c/3

			// Step 4: A repeated digit in ANY group makes the board invalid.
			if rows[r][digit] || cols[c][digit] || boxes[boxIndex][digit] {
				return false
			}

			// Step 5: Mark the cell after checking so it cannot match itself.
			rows[r][digit] = true
			cols[c][digit] = true
			boxes[boxIndex][digit] = true
		}
	}
	return true // No constraint was violated; completeness is not required.
}

func main() {
	board := [][]byte{
		[]byte("53..7...."), []byte("6..195..."), []byte(".98....6."),
		[]byte("8...6...3"), []byte("4..8.3..1"), []byte("7...2...6"),
		[]byte(".6....28."), []byte("...419..5"), []byte("....8..79"),
	}
	fmt.Println(isValidSudoku(board)) // true
}
