package day01

import (
	"fmt"
)

func calculateSimilarity(list1 []int, list2 []int) int {
	similarity := 0
	for _, num1 := range list1 {
		occurences := 0
		for _, num2 := range list2 {
			if num1 == num2 {
				occurences++
			}
		}
		similarity += num1 * occurences
	}
	return similarity
}

func Part2(filename string) {
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

	sim := calculateSimilarity(list1, list2)
	fmt.Printf("Similarity: %d\n", sim)
}
