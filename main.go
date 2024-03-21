package main

import (
	"github.com/rivo/tview"
	"time"
)

func main() {
	var board [][]string = make([][]string, 50)
	for i := range board {
		board[i] = make([]string, 50)
	}
	initializeBoard(board)
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetSize(50, 50)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				updateBoard(board)
				app.QueueUpdateDraw(func() {
					textView.SetText(boardAsString(board))
				})
			}
		}
	}()

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}

func initializeBoard(board [][]string) {
	for i := range board {
		for j := range board[i] {
			board[i][j] = " "
		}
	}
	board[1][2] = "X"
	board[2][3] = "X"
	board[3][1] = "X"
	board[3][2] = "X"
	board[3][3] = "X"
	board[4][2] = "X"
	board[4][3] = "X"
	board[5][2] = "X"
}

func getNeighbors(board [][]string, x int, y int) int {
	var neighbors int
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i >= len(board) || y+j < 0 || y+j >= len(board[0]) {
				continue
			}
			if board[x+i][y+j] == "X" {
				neighbors++
			}
		}
	}
	return neighbors
}

// Any live cell with fewer than two live neighbors dies, as if by underpopulation.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by overpopulation.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

func updateBoard(board [][]string) {
	for i := range board {
		for j := range board[i] {
			neighbors := getNeighbors(board, i, j)
			switch {
			case neighbors < 2:
				board[i][j] = " "
			case neighbors == 2:
				continue
			case neighbors < 4:
				board[i][j] = "X"
			default:
				board[i][j] = " "
			}
		}
	}
}

func boardAsString(board [][]string) string {
	var boardString string
	for i := range board {
		for j := range board[i] {
			boardString += board[i][j]
		}
		boardString += "\n"
	}
	return boardString
}
