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
		t.Errorf("Expected %v, got %v", expected, board)
	}
}

func TestUpdate(t *testing.T) {
	board := NewDefault(10, 10)
	Update(board)

	expected := [][]string{
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", "X", " ", " ", " ", " ", " ", " "},
		{" ", "X", " ", "X", "X", " ", " ", " ", " ", " "},
		{" ", "X", " ", "X", "X", " ", " ", " ", " ", " "},
		{" ", " ", "X", "X", "X", " ", " ", " ", " ", " "},
		{" ", " ", " ", "X", "X", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	}

	if !boardsEqual(board, expected) {
		t.Errorf("Expected: %v got: %v", AsString(expected), AsString(board))
	}

	Update(board)

	expected = [][]string{
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", "X", "X", "X", " ", " ", " ", " ", " "},
		{" ", "X", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", "X", " ", " ", "X", " ", " ", " ", " ", " "},
		{" ", " ", "X", " ", "X", "X", " ", " ", " ", " "},
		{" ", " ", " ", "X", "X", "X", " ", " ", " ", " "},
		{" ", " ", " ", " ", "X", "X", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	}

	if !boardsEqual(board, expected) {
		t.Errorf("Expected: %v got: %v", AsString(expected), AsString(board))
	}
}

func TestAsString(t *testing.T) {
	board := NewDefault(10, 10)
	expected := "          \n  X       \n   X      \n XXX      \n  XX      \n  X       \n          \n          \n          \n          \n"

	if AsString(board) != expected {
		t.Errorf("Expected:\n%s got:\n%s", strings.ReplaceAll(expected, " ", "*"), strings.ReplaceAll(AsString(board), " ", "*"))
	}
}
