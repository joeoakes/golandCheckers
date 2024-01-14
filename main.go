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
	// Implement your move validation logic here
	// Check if the move is valid according to the rules
	return true
}

func makeMove(move string, board Board) Board {
	// Implement your move logic here
	// Update the board based on the move
	return board
}

func togglePlayer(currentPlayer string) string {
	if currentPlayer == playerX {
		return playerO
	}
	return playerX
}
