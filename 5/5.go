package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	filename = "input.txt"
)

// Read page rules and updates from a file
// Returns a map of rules and a 2D list of page updates
func readFile(filename string) (map[int][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var rules []string
	var updates [][]int
	var isRule bool = true

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRule = false
			continue
		}
		if isRule {
			rules = append(rules, line)
		} else {
			var update []int
			for _, s := range strings.Split(line, ",") {
				num, err := strconv.Atoi(s)
				if err != nil {
					return nil, nil, err
				}
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	rulesMap := make(map[int][]int)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("Invalid rule: %s", rule)
		}
		before, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Unable to convert '%s' to an int", parts[0])
			return nil, nil, err
		}
		after, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Unable to convert '%s' to an int", parts[1])
			return nil, nil, err
		}
		rulesMap[before] = append(rulesMap[before], after)
	}

	return rulesMap, updates, nil
}

// Get the numbete that the center index of an update
func getMiddlePage(arr []int) int {
	index := len(arr) / 2
	return arr[index]
}

// Check if the update is valid based on the rules
func isValidUpdate(rules map[int][]int, update []int) bool {
	// Create a map to store the position of each number in the update list
	position := make(map[int]int)
	for i, num := range update {
		position[num] = i
	}

	// Iterate through the rules
	for key, values := range rules {
		// Get the position of the key in the update list
		keyPos, exists := position[key]
		if !exists {
			continue
		}

		// Check if any of the values appear before the key in the update list
		for _, value := range values {
			valuePos, exists := position[value]
			if exists && valuePos < keyPos {
				return false
			}
		}
	}

	return true
}

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

func main() {
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

	sortedMiddleSum := 0
	for _, update := range updates {
		if !isValidUpdate(rules, update) {
			sortedUpdate := sortUpdate(rules, update)
			sortedMiddleSum += getMiddlePage(sortedUpdate)
		}
	}
	fmt.Printf("Sum of sorted update middle pages: %d\n", sortedMiddleSum)
}
