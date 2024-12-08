package main

import (
	"fmt"
	"os"
)

const fileName = "sample-input.txt"

// const fileName = "input.txt"

func main() {
	fileBytes, _ := os.ReadFile(fileName)
	file := string(fileBytes)
	fmt.Println(file)
}

