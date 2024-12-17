package day05

import (
	"fmt"
)

func Part1(filename string) {
	rules, updates, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	validMiddleSum := 0
	for _, update := range updates {
		if isValidUpdate(rules, update) {
			validMiddleSum += getMiddlePage(update)
		}
	}
	fmt.Printf("Sum of valid update middle pages: %d\n", validMiddleSum)
}
