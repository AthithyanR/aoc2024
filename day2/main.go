package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fileName = "sample-input.txt"

// const fileName = "input.txt"

// func handleReport(report string) int {
// 	levelstrs := strings.Fields(report)
// 	levels := []int{}

// 	for _, levelStr := range levelstrs {
// 		level, _ := strconv.Atoi(levelStr)
// 		levels = append(levels, level)
// 	}

// 	isAsc := levels[1] > levels[0]

// 	for i := 1; i < len(levels); i++ {
// 		if isAsc {
// 			if levels[i] > levels[i-1] {
// 				diff := levels[i] - levels[i-1]
// 				if diff < 1 || diff > 3 {
// 					return 0
// 				}
// 			} else {
// 				return 0
// 			}
// 		} else {
// 			if levels[i] < levels[i-1] {
// 				diff := levels[i-1] - levels[i]
// 				if diff < 1 || diff > 3 {
// 					return 0
// 				}
// 			} else {
// 				return 0
// 			}
// 		}
// 	}

// 	return 1
// }

func handleReport(report string) int {
	fmt.Println(report)

	levelstrs := strings.Fields(report)
	levels := []int{}

	for _, levelStr := range levelstrs {
		level, _ := strconv.Atoi(levelStr)
		levels = append(levels, level)
	}

	fmt.Println(levels)

	for i := 0; i < len(levels); i++ {
		newLevels := append(levels[:i], levels[i+1:]...)
		fmt.Println(i, newLevels)

		isAsc := newLevels[1] > newLevels[0]

		isOkay := true

		for i := 1; i < len(newLevels) && isOkay; i++ {
			if isAsc {
				if newLevels[i] > newLevels[i-1] {
					diff := newLevels[i] - newLevels[i-1]
					if diff < 1 || diff > 3 {
						isOkay = false
					}
				} else {
					isOkay = false
				}
			} else {
				if newLevels[i] < newLevels[i-1] {
					diff := newLevels[i-1] - newLevels[i]
					if diff < 1 || diff > 3 {
						isOkay = false
					}
				} else {
					isOkay = false
				}
			}
		}

		if isOkay {
			return 1
		}

	}

	return 0

}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to open the input file %+v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeCount := 0

	for scanner.Scan() {
		report := scanner.Text()
		safeCount += handleReport(report)
		// fmt.Println(safeCount)
	}

	fmt.Println(safeCount)

	// part 1

	// end part 1

	// part 2

	// end part 2

}
