package main

import "testing"

func emptyBoard() [][]byte {
	board := make([][]byte, 9)
	for r := range board {
		board[r] = []byte(".........")
	}
	return board
}

func TestIsValidSudoku(t *testing.T) {
	validExample := [][]byte{
		[]byte("53..7...."), []byte("6..195..."), []byte(".98....6."),
		[]byte("8...6...3"), []byte("4..8.3..1"), []byte("7...2...6"),
		[]byte(".6....28."), []byte("...419..5"), []byte("....8..79"),
	}
	invalidExample := [][]byte{
		[]byte("83..7...."), []byte("6..195..."), []byte(".98....6."),
		[]byte("8...6...3"), []byte("4..8.3..1"), []byte("7...2...6"),
		[]byte(".6....28."), []byte("...419..5"), []byte("....8..79"),
	}
	rowDup := emptyBoard()
	rowDup[0][0], rowDup[0][8] = '5', '5' // Different boxes, same row.
	colDup := emptyBoard()
	colDup[0][0], colDup[8][0] = '5', '5' // Different boxes, same column.
	boxDup := emptyBoard()
	boxDup[0][0], boxDup[2][2] = '5', '5' // Different rows and columns, same box.
	validAcrossBoxes := emptyBoard()
	validAcrossBoxes[0][0], validAcrossBoxes[3][3] = '5', '5'
	lastBoxDup := emptyBoard()
	lastBoxDup[6][6], lastBoxDup[8][8] = '9', '9'

	tests := []struct {
		name  string
		board [][]byte
		want  bool
	}{
		{"official valid", validExample, true},
		{"official invalid", invalidExample, false},
		{"empty board", emptyBoard(), true},
		{"row duplicate", rowDup, false},
		{"column duplicate", colDup, false},
		{"box duplicate", boxDup, false},
		{"same digit in unrelated groups", validAcrossBoxes, true},
		{"bottom right box index", lastBoxDup, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isValidSudoku(tc.board); got != tc.want {
				t.Fatalf("isValidSudoku returned %t, want %t", got, tc.want)
			}
		})
	}
}
