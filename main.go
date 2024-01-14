package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	boardSize = 8
	emptyCell = " "
	playerX   = "X"
	playerO   = "O"
)

type Board [boardSize][boardSize]string

func initializeBoard() Board {
	board := Board{}
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if (row+col)%2 == 0 {
				board[row][col] = emptyCell
			} else if row < 3 {
				board[row][col] = playerX
			} else if row > 4 {
				board[row][col] = playerO
			} else {
				board[row][col] = emptyCell
			}
		}
	}
	return board
}

func printBoard(board Board) {
	fmt.Println("  0 1 2 3 4 5 6 7")
	for row := 0; row < boardSize; row++ {
		fmt.Printf("%d ", row)
		for col := 0; col < boardSize; col++ {
			fmt.Printf("%s ", board[row][col])
		}
		fmt.Println()
	}
}

func main() {
	board := initializeBoard()
	currentPlayer := playerX
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Current player: %s\n", currentPlayer)
		printBoard(board)
		fmt.Print("Enter move (e.g., '2 3 to 3 4'): ")
		scanner.Scan()
		move := scanner.Text()

		if isValidMove(move, board, currentPlayer) {
			board = makeMove(move, board)
			currentPlayer = togglePlayer(currentPlayer)
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}

func isValidMove(move string, board Board, currentPlayer string) bool {
	// Parse the move string into fromRow, fromCol, toRow, toCol
	var fromRow, fromCol, toRow, toCol int
	_, err := fmt.Sscanf(move, "%d %d to %d %d", &fromRow, &fromCol, &toRow, &toCol)
	if err != nil {
		return false
	}

	// Check if the source cell is valid and contains the current player's piece
	if !isValidCell(fromRow, fromCol) || board[fromRow][fromCol] != currentPlayer {
		return false
	}

	// Check if the destination cell is valid and empty
	if !isValidCell(toRow, toCol) || board[toRow][toCol] != emptyCell {
		return false
	}

	// Check if the move is diagonal
	if abs(toRow-fromRow) != 1 || abs(toCol-fromCol) != 1 {
		return false
	}

	// Check if the move direction is valid based on the player (forward for playerX, backward for playerO)
	if (currentPlayer == playerX && toRow < fromRow) || (currentPlayer == playerO && toRow > fromRow) {
		return false
	}

	return true
}

func isValidCell(row, col int) bool {
	return row >= 0 && row < boardSize && col >= 0 && col < boardSize
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func makeMove(move string, board Board) Board {
	// Parse the move string into fromRow, fromCol, toRow, toCol
	var fromRow, fromCol, toRow, toCol int
	_, err := fmt.Sscanf(move, "%d %d to %d %d", &fromRow, &fromCol, &toRow, &toCol)
	if err != nil {
		return board // Return the unchanged board if the move is invalid
	}

	// Copy the current board to avoid modifying the original
	newBoard := board

	// Update the destination cell with the piece from the source cell
	newBoard[toRow][toCol] = newBoard[fromRow][fromCol]

	// Empty the source cell
	newBoard[fromRow][fromCol] = emptyCell

	return newBoard
}

func togglePlayer(currentPlayer string) string {
	if currentPlayer == playerX {
		return playerO
	}
	return playerX
}
