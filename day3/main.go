package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// const fileName = "sample-input.txt"

const fileName = "input.txt"

func main() {
	fileBytes, _ := os.ReadFile(fileName)
	file := string(fileBytes)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)
	matches := re.FindAllStringSubmatch(file, -1)
	fmt.Println(matches)

	// sum := 0

	// for _, match := range matches {
	// 	num1, _ := strconv.Atoi(match[1])
	// 	num2, _ := strconv.Atoi(match[2])
	// 	sum += (num1 * num2)
	// }

	// fmt.Println(sum)

	sum := 0
	do := true

	for _, match := range matches {
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else {
			if do {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				sum += (num1 * num2)
			}
		}
	}

	fmt.Println(sum)
}
