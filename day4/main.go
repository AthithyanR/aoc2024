package main

import (
	"fmt"
	"os"
	"strings"
)

// const fileName = "sample-input.txt"

const fileName = "input.txt"

const XMAS string = "XMAS"

func main() {
	fileBytes, _ := os.ReadFile(fileName)
	file := string(fileBytes)

	input := [][]string{}
	for _, line := range strings.Split(file, "\n") {
		chars := strings.Split(line, "")
		input = append(input, chars)
	}

	count := 0

	// define height and width
	H := len(input)
	W := len(input[0])

	// guard function
	// inBounds := func(row int, col int) bool {
	// 	return row >= 0 && row < H && col >= 0 && col < W
	// }

	// part 1
	// // for each row & col
	// for row := 0; row < H; row++ {
	// 	for col := 0; col < W; col++ {
	// 		// if X
	// 		if input[row][col] == "X" {
	// 			// define directions & use them as delta
	// 			for dRow := -1; dRow <= 1; dRow++ {
	// 				for dCol := -1; dCol <= 1; dCol++ {
	// 					// skip self (not mandatory)
	// 					if dRow == 0 && dCol == 0 {
	// 						continue
	// 					}

	// 					allGood := true
	// 					for i := 0; i < len(XMAS); i++ {
	// 						// cur + (delta * growth)
	// 						nRow := row + (dRow * i)
	// 						nCol := col + (dCol * i)

	// 						if inBounds(nRow, nCol) && input[nRow][nCol] == string(XMAS[i]) {
	// 						} else {
	// 							allGood = false
	// 							break
	// 						}
	// 					}

	// 					if allGood {
	// 						count += 1
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	pairs := [4][2]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	// for each row & col (skip ones that could not have neighbors)
	for row := 1; row < H-1; row++ {
		for col := 1; col < W-1; col++ {
			// if A
			if input[row][col] == "A" {
				str := ""
				for _, pair := range pairs {
					str += input[row+pair[0]][col+pair[1]]
				}
				if str == "MMSS" || str == "SSMM" || str == "SMMS" || str == "MSSM" {
					count += 1
				}
			}
		}
	}

	fmt.Println(count)
}

/*
Note to self:

my initial approach was to
	* go through the matrix horizontally, vertically, diagonally right and left + backwards for all and find the count, this is unncessary and complex

better solution is to
	* go through the matrix once
	* find if X
		* define directions & use them as delta
		* check until growth complete add count
*/
