package main

import "fmt"

func main() {
	levels := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(levels); i++ {
		newLevels := make([]int, len(levels)-1)
		copy(newLevels, levels[:i])
		copy(newLevels[i:], levels[i+1:])
		fmt.Println(i, newLevels)
	}
}
