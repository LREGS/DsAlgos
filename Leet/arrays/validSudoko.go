package main

import (
	"fmt"
	"strconv"
)

//could use bucket sort to sort the numbers in each square and then sort through each column at the same time
//if a duplicate is found return fasle

func isValidSudoko(board [][]byte) bool {

	// //new square every 3 index
	// squares := make(map[int]byte)

	// //column = i
	// column := make(map[int]byte)

	currentSquare := ""

	for i, line := range board {
		statement := fmt.Sprintf("new line " + strconv.Itoa(i) + "\n")
		fmt.Printf(statement)
		for j, _ := range line {
			switch {
			case i <= 2 && j <= 2:
				currentSquare = "one"
			case i <= 2 && 2 < j && j < 6:
				currentSquare = "two"
			case i <= 2 && 6 <= j && j <= 8:
				currentSquare = "three"
			case 2 < i && i < 6 && j <= 2:
				currentSquare = "four"
			case 2 < i && i < 6 && 2 < j && j < 6:
				currentSquare = "five"
			case 2 < i && i < 6 && 6 <= j && j <= 8:
				currentSquare = "six"
			case 6 <= i && i <= 9 && j <= 2:
				currentSquare = "seven"
			case 6 <= i && i <= 8 && 2 < j && j < 6:
				currentSquare = "eight"
			case 6 <= i && i <= 8 && 6 <= j && j <= 8:
				currentSquare = "nine"
			}

		}

	}
	return true

}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	isValidSudoko(board)
}
