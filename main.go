package main

import (
	"fmt"
	"os"
	"strconv"
)

const fatalErr string = "Error"

func main() {
	if len(os.Args) < 10 {
		fmt.Println(fatalErr)
		return
	}

	field := parseField(os.Args[1:])

	// Solve the sudoku
	if solveSudoku(&field) {
		// Print the solved field
		for _, line := range field {
			for i, value := range line {
				fmt.Print(value)
				if i != len(line)-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println(fatalErr)
	}
}

func parseField(args []string) [9][9]int {
	if len(args) != 9 {
		fmt.Println(fatalErr)
	}

	var field [9][9]int

	for i := 0; i < 9; i++ {
		row := args[i]

		for j, char := range row {
			if char == '.' {
				field[i][j] = 0 // If the character is '.', set the corresponding value to 0
			} else {
				num, err := strconv.Atoi(string(char))
				if err != nil {
					fmt.Println(fatalErr)
				}
				field[i][j] = num // Convert the character to an integer and set the corresponding value
			}
		}
	}

	return field
}

func solveSudoku(field *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if field[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isSafe(field, row, col, num) {
						field[row][col] = num
						if solveSudoku(field) {
							return true
						}
						field[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func isSafe(field *[9][9]int, row, col, num int) bool {
	return !usedInRow(field, row, num) &&
		!usedInColumn(field, col, num) &&
		!usedInBox(field, row-row%3, col-col%3, num)
}

func usedInRow(field *[9][9]int, row, num int) bool {
	for col := 0; col < 9; col++ {
		if field[row][col] == num {
			return true
		}
	}
	return false
}

func usedInColumn(field *[9][9]int, col, num int) bool {
	for row := 0; row < 9; row++ {
		if field[row][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(field *[9][9]int, boxStartRow, boxStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if field[row+boxStartRow][col+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}
