package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const fileName = "sample-input.txt"

const fileName = "input.txt"

func main() {
	fileBytes, _ := os.ReadFile(fileName)
	file := string(fileBytes)
	segments := strings.Split(file, "\n\n")
	rulesRaw, updatesRaw := segments[0], segments[1]

	rules := [][]int{}
	for _, rule := range strings.Split(rulesRaw, "\n") {
		res := strings.Split(rule, "|")
		first, _ := strconv.Atoi(res[0])
		second, _ := strconv.Atoi(res[1])
		entry := []int{first, second}
		rules = append(rules, []int(entry))
	}

	updates := [][]int{}
	for _, updateStr := range strings.Split(updatesRaw, "\n") {
		entry := []int{}
		for _, pageNumStr := range strings.Split(updateStr, ",") {
			pageNum, _ := strconv.Atoi(pageNumStr)
			entry = append(entry, pageNum)
		}
		updates = append(updates, entry)
	}

	p1 := 0

	for _, update := range updates {
		isGood := true

		for _, rule := range rules {
			firstIndex := -1
			secondIndex := -1

			for index, pageNum := range update {
				if firstIndex == -1 && pageNum == rule[0] {
					firstIndex = index
				}

				if secondIndex == -1 && pageNum == rule[1] {
					secondIndex = index
				}
			}

			if firstIndex > -1 && secondIndex > -1 {
				if firstIndex >= secondIndex {
					isGood = false
					break
				}
			}
		}

		if isGood {
			middleIndex := len(update) / 2
			p1 += update[middleIndex]
		}
	}

	fmt.Println(p1) // 5452

}
