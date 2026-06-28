package main

import (
	"fmt"
	"strconv"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func main() {
	var cells [9]string
	var input int
	var turn int = 1
	var mark string

	// Initialize board
	for i := range cells {
		cells[i] = strconv.Itoa(i + 1)
	}

	printBoard(cells)

	for {
		// Determine current player
		if turn == 1 {
			fmt.Print("Player 1 (X), choose a cell (1-9): ")
			mark = Green + "X" + Reset
		} else {
			fmt.Print("Player 2 (O), choose a cell (1-9): ")
			mark = Red + "O" + Reset
		}

		// Read input
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Please enter a number.")
			continue
		}

		input--

		// Check range
		if input < 0 || input > 8 {
			fmt.Println("Please choose a number between 1 and 9.")
			continue
		}

		// Check if occupied
		if cells[input] == "X" || cells[input] == "O" {
			fmt.Println("That cell is already taken.")
			continue
		}

		// Place mark
		cells[input] = mark

		printBoard(cells)

		// Check for win
		if checkWin(cells, mark) {
			if mark == "X" {
				fmt.Println("Player 1 wins!")
			} else {
				fmt.Println("Player 2 wins!")
			}
			break
		}

		// Check for draw
		if boardFull(cells) {
			fmt.Println("It's a draw!")
			break
		}

		// Switch turns
		if turn == 1 {
			turn = 2
		} else {
			turn = 1
		}
	}
}

// Prints the board
func printBoard(cells [9]string) {
	fmt.Print("\033[2J\033[H")

	fmt.Printf(
		" %s | %s | %s \n"+
			"---+---+---\n"+
			" %s | %s | %s \n"+
			"---+---+---\n"+
			" %s | %s | %s \n",
		cells[0], cells[1], cells[2],
		cells[3], cells[4], cells[5],
		cells[6], cells[7], cells[8],
	)
}

// Returns true if the given player has won
func checkWin(cells [9]string, mark string) bool {
	wins := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, w := range wins {
		if cells[w[0]] == mark &&
			cells[w[1]] == mark &&
			cells[w[2]] == mark {
			return true
		}
	}

	return false
}

// Returns true if there are no empty cells
func boardFull(cells [9]string) bool {
	for _, cell := range cells {
		if cell != "X" && cell != "O" {
			return false
		}
	}
	return true
}
