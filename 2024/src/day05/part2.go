package day05

import (
	"fmt"
)

// Fix the update list so that it is valid based on the rules
func sortUpdate(rules map[int][]int, update []int) []int {
	position := make(map[int]int)
	for i, num := range update {
		position[num] = i
	}

	// Function to swap elements in the update list
	swap := func(i, j int) {
		update[i], update[j] = update[j], update[i]
		position[update[i]] = i
		position[update[j]] = j
	}

	// Iterate through the rules and fix the update list
	for !isValidUpdate(rules, update) {
		for key, values := range rules {
			for _, value := range values {
				keyPos, keyExists := position[key]
				valuePos, valueExists := position[value]
				if keyExists && valueExists && valuePos < keyPos {
					swap(keyPos, valuePos)
				}
			}
		}
	}

	return update
}

func Part2(filename string) {
	rules, updates, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sortedMiddleSum := 0
	for _, update := range updates {
		if !isValidUpdate(rules, update) {
			sortedUpdate := sortUpdate(rules, update)
			sortedMiddleSum += getMiddlePage(sortedUpdate)
		}
	}
	fmt.Printf("Sum of sorted update middle pages: %d\n", sortedMiddleSum)
}
