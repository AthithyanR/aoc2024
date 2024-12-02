package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// const fileName = "sample-input.txt"

const fileName = "input.txt"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to open the input file %+v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		text := scanner.Text()

		// load the lists with locationId
		locationIds := strings.Fields(text)

		if len(locationIds) != 2 {
			log.Fatalf("did not get the two locations")
		}

		locationId1, err := strconv.Atoi(locationIds[0])
		if err != nil {
			log.Fatalf("unable to convert locationId1")
		}
		list1 = append(list1, locationId1)

		locationId2, err := strconv.Atoi(locationIds[1])
		if err != nil {
			log.Fatalf("unable to convert locationId2")
		}
		list2 = append(list2, locationId2)
	}

	if len(list1) != len(list2) {
		log.Fatalf("Lists are not equal")
	}

	// part 1

	// sort the list 1 and 2
	// sort.Ints(list1)
	// sort.Ints(list2)

	// distance := 0

	// for i := 0; i < len(list1); i++ {
	// 	diff := list1[i] - list2[i]
	// 	if diff < 0 {
	// 		diff = list2[i] - list1[i]
	// 	}

	// 	distance += diff
	// }

	// fmt.Println(distance)

	// end part 1

	// part 2

	list2FreqMap := make(map[int]int)

	for _, num := range list2 {
		list2FreqMap[num]++
	}

	// log.Println(list2FreqMap)

	simScore := 0

	for _, num := range list1 {
		simScore += num * list2FreqMap[num]
	}

	fmt.Println("simScore is", simScore)

	// end part 2

}
