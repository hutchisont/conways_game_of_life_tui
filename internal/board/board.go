package board

// Creates a new board with the given number of rows and columns
// Use this if you want to manage the initial state of the board yourself
func NewBlank(row int, col int) [][]string {
	board := make([][]string, row)
	for i := range board {
		board[i] = make([]string, col)
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = " "
		}
	}

	return board
}

// Creates a new board with the given number of rows and columns and initializes
// it with the default starting state
// Use this if you just want to go without thinking about initializing a board
func NewDefault(row int, col int) [][]string {
	board := NewBlank(row, col)
	defaultInitialize(board)
	return board
}

func NewRPentomino(row int, col int) [][]string {
	board := NewBlank(row, col)

	// will be the below pattern in the center of the board
	// **********
	// ****OO****
	// ***OO*****
	// ****O*****
	// **********

	baseRow := row / 2
	baseCol := col / 2

	board[baseRow][baseCol] = "O"
	board[baseRow][baseCol+1] = "O"
	board[baseRow+1][baseCol-1] = "O"
	board[baseRow+1][baseCol] = "O"
	board[baseRow+2][baseCol] = "O"

	return board
}

func defaultInitialize(board [][]string) {
	// **********
	// **O*******
	// ***O******
	// *OOO******
	// **OO******
	// **O*******
	// **********
	// **********
	// **********
	// **********

	board[1][2] = "O"
	board[2][3] = "O"
	board[3][1] = "O"
	board[3][2] = "O"
	board[3][3] = "O"
	board[4][2] = "O"
	board[4][3] = "O"
	board[5][2] = "O"
}

func getNeighbors(board [][]string, r int, c int) int {
	var neighbors int
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if r+i < 0 || r+i >= len(board) || c+j < 0 || c+j >= len(board[0]) {
				continue
			}
			if board[r+i][c+j] == "O" {
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
func Update(board [][]string) {
	newBoard := NewBlank(len(board), len(board[0]))
	for i := range board {
		for j := range board[i] {
			neighbors := getNeighbors(board, i, j)
			switch {
			case neighbors < 2:
				newBoard[i][j] = " "
			case neighbors == 2:
				newBoard[i][j] = board[i][j]
			case neighbors < 4:
				newBoard[i][j] = "O"
			default:
				newBoard[i][j] = " "
			}
		}
	}

	for i := range board {
		copy(board[i], newBoard[i])
	}
}

func AsString(board [][]string) string {
	var boardString string
	for i := range board {
		for j := range board[i] {
			boardString += board[i][j]
		}
		boardString += "\n"
	}
	return boardString
}
