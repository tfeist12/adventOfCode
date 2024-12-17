package day01

import (
	"fmt"
	"sort"
)

func calculateDifference(list1 []int, list2 []int) int {
	// Sort lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Calculate sum of differences between lists
	sum := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] >= list2[i] {
			diff := list1[i] - list2[i]
			sum += diff
		} else {
			diff := list2[i] - list1[i]
			sum += diff
		}
	}

	return sum
}

func Part1(filename string) {
	list1, list2, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Ensure lists are the same length
	if len(list1) != len(list2) {
		fmt.Println("Lists are not the same length")
		return
	}

	diff := calculateDifference(list1, list2)
	fmt.Printf("Total Difference: %d\n", diff)
}
