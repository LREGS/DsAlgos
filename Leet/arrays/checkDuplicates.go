// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func isValidSudoku(board [][]byte) bool {
// 	squares := make(map[string][]byte)

// 	columns := make(map[int][]byte)
// 	rows := make(map[int][]byte)

// 	for row, line := range board {
// 		for column, value := range line {
// 			//this is the value of an empty string
// 			if value == '.' {
// 				continue
// 			}

// 			if ContainsDuplicate(columns[column], value) {
// 				fmt.Println(columns[column], value)
// 				return false
// 			}
// 			columns[column] = append(columns[column], value)

// 			if ContainsDuplicate(rows[row], value) {
// 				fmt.Println(rows[row], value)
// 				return false
// 			}
// 			rows[row] = append(rows[row], value)

// 			grid := fmt.Sprintf(strconv.Itoa(row/3) + strconv.Itoa(column/3))
// 			if ContainsDuplicate(squares[grid], value) {
// 				return false
// 			}
// 			squares[grid] = append(squares[grid], value)

// 		}
// 	}

// 	return true
// }

// func ContainsDuplicate(s []byte, v byte) bool {
// 	for _, value := range s {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	board := [][]byte{
// 		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		{'6', '5', '.', '1', '9', '5', '.', '.', '.'},
// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
// 	}
// 	err := isValidSudoku(board)
// 	if err != true {
// 		fmt.Println("invalid")
// 	} else {
// 		fmt.Println("valid")
// 	}
// }
