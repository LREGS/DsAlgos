package main

import (
	"fmt"
)

//could use bucket sort to sort the numbers in each square and then sort through each column at the same time
//if a duplicate is found return fasle

func isValidSudoko(board [][]byte) bool {

	//new square every 3 index
	squares := make(map[int][]byte)

	//column = i
	column := make(map[int][]byte)

	// currentSquare := ""

	for i, line := range board {
		for j, value := range line {
			if value == 46 {
				continue
			}
			err := checkDuplicate(j, value, column)
			if err != true {
				return false
			}
			column[j] = append(column[j], value)

			switch {
			case i <= 2 && j <= 2:
				err := checkDuplicate(i+1, value, column)
				if err != true {
					return false
				}
				squares[1] = append(squares[1], value)
			case i <= 2 && 2 < j && j < 6:
				err := checkDuplicate(i+1, value, column)
				if err != true {
					return false
				}
				squares[2] = append(squares[2], value)
			case i <= 2 && 6 <= j && j <= 8:
				err := checkDuplicate(i+1, value, column)
				if err != true {
					return false
				}
				squares[3] = append(squares[3], value)
			case 2 < i && i < 6 && j <= 2:
				err := checkDuplicate(i+1, value, column)
				if err != true {
					return false
				}
				squares[4] = append(squares[4], value)
			case 2 < i && i < 6 && 2 < j && j < 6:
				err := checkDuplicate(i+1, value, column)
				if err != true {
					return false
				}
				squares[5] = append(squares[5], value)
			case 2 < i && i < 6 && 6 <= j && j <= 8:
				err := checkDuplicate(i+1, value, column)
				if err != true {
					return false
				}
				squares[6] = append(squares[6], value)
			case 6 <= i && i <= 9 && j <= 2:
				err := checkDuplicate(i+1, value, column)
				if err != true {
					return false
				}
				squares[7] = append(squares[7], value)
			case 6 <= i && i <= 8 && 2 < j && j < 6:
				err := checkDuplicate(i+1, value, column)
				if err != true {
					return false
				}
				squares[8] = append(squares[8], value)
			case 6 <= i && i <= 8 && 6 <= j && j <= 8:
				err := checkDuplicate(i+1, value, column)
				if err != true {
					return false
				}
				squares[9] = append(squares[9], value)
			}

		}

	}
	fmt.Println(squares, column)
	return true

}

func checkDuplicate(k int, v byte, m map[int][]byte) bool {
	for _, value := range m[k] {
		if value == v {
			return false
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
	err := isValidSudoko(board)
	if err != false {
		fmt.Println("valid sudoko")
	} else {
		fmt.Println("invalid sudoko")
	}
}
