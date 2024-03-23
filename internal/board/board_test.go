package board

import (
	"strings"
	"testing"
)

func boardsEqual(board1 [][]string, board2 [][]string) bool {
	if len(board1) != len(board2) {
		return false
	}
	for i := range board1 {
		if len(board1[i]) != len(board2[i]) {
			return false
		}
		for j := range board1[i] {
			if board1[i][j] != board2[i][j] {
				return false
			}
		}
	}
	return true
}

func TestNewBlank(t *testing.T) {
	rows := 10
	cols := 10
	board := NewBlank(rows, cols)
	if len(board) != rows {
		t.Errorf("Expected %d rows, got %d", rows, len(board))
	}
	if len(board[0]) != cols {
		t.Errorf("Expected %d columns, got %d", cols, len(board[0]))
	}
}

func TestNewDefault(t *testing.T) {
	rows := 10
	cols := 10
	board := NewDefault(rows, cols)
	if len(board) != rows {
		t.Errorf("Expected %d rows, got %d", rows, len(board))
	}
	if len(board[0]) != cols {
		t.Errorf("Expected %d columns, got %d", cols, len(board[0]))
	}

	expected := [][]string{
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", "X", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", "X", " ", " ", " ", " ", " ", " "},
		{" ", "X", "X", "X", " ", " ", " ", " ", " ", " "},
		{" ", " ", "X", "X", " ", " ", " ", " ", " ", " "},
		{" ", " ", "X", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	}

	if !boardsEqual(board, expected) {
		t.Errorf("Expected:\n%s got:\n%s",
			strings.ReplaceAll(AsString(expected), " ", "*"),
			strings.ReplaceAll(AsString(board), " ", "*"))
	}
}

func TestThreeRowInit(t *testing.T) {
	board := NewBlank(10, 10)
	board[5][5] = "X"
	board[5][6] = "X"
	board[5][7] = "X"

	expected := [][]string{
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", "X", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", "X", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", "X", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	}

	Update(board)

	if !boardsEqual(board, expected) {
		t.Errorf("Expected:\n%s got:\n%s",
			strings.ReplaceAll(AsString(expected), " ", "*"),
			strings.ReplaceAll(AsString(board), " ", "*"))
	}

	expected = [][]string{
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", "X", "X", "X", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	}

	Update(board)

	if !boardsEqual(board, expected) {
		t.Errorf("Expected:\n%s got:\n%s",
			strings.ReplaceAll(AsString(expected), " ", "*"),
			strings.ReplaceAll(AsString(board), " ", "*"))
	}
}

func TestRPentominoInit(t *testing.T) {
	board := NewRPentomino(10, 10)

	// **********
	// **********
	// **********
	// ***XXX****
	// ***X******
	// ***XX*****
	// **********
	// **********
	// **********
	// **********

	expected := [][]string{
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", "X", "X", "X", " ", " ", " "},
		{" ", " ", " ", " ", "X", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", "X", "X", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	}

	Update(board)

	if !boardsEqual(board, expected) {
		t.Errorf("Expected:\n%s got:\n%s",
			strings.ReplaceAll(AsString(expected), " ", "*"),
			strings.ReplaceAll(AsString(board), " ", "*"))
	}
}

func TestGetNeighbors(t *testing.T) {
	board := NewBlank(10, 10)
	board[3][1] = "X"
	board[3][2] = "X"
	board[3][3] = "X"
	board[4][2] = "X"
	board[4][3] = "X"
	board[5][2] = "X"

	// **********
	// **********
	// **********
	// *XXX******
	// **XX******
	// **X*******
	// **********
	// **********
	// **********
	// **********

	neighbors := getNeighbors(board, 3, 1)
	if neighbors != 2 {
		t.Errorf("Expected 2, got %d", neighbors)
	}

	neighbors = getNeighbors(board, 3, 2)
	if neighbors != 4 {
		t.Errorf("Expected 4, got %d", neighbors)
	}

	neighbors = getNeighbors(board, 3, 3)
	if neighbors != 3 {
		t.Errorf("Expected 3, got %d", neighbors)
	}

	neighbors = getNeighbors(board, 4, 2)
	if neighbors != 5 {
		t.Errorf("Expected 5, got %d", neighbors)
	}

	neighbors = getNeighbors(board, 4, 3)
	if neighbors != 4 {
		t.Errorf("Expected 4, got %d", neighbors)
	}

	neighbors = getNeighbors(board, 5, 2)
	if neighbors != 2 {
		t.Errorf("Expected 2, got %d", neighbors)
	}
}

func TestAsString(t *testing.T) {
	board := NewDefault(10, 10)
	expected := "          \n  X       \n   X      \n XXX      \n  XX      \n  X       \n          \n          \n          \n          \n"

	if AsString(board) != expected {
		t.Errorf("Expected:\n%s got:\n%s", strings.ReplaceAll(expected, " ", "*"), strings.ReplaceAll(AsString(board), " ", "*"))
	}
}
