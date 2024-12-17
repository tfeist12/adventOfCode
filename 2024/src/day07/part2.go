package day07

import (
	"fmt"
	"strconv"
)

// Helper function to recursively evaluate all combinations
func evaluateHelperAddMultConcat(target int, constants []int, current int, index int) bool {
	if index == len(constants) {
		return current == target
	}

	// Try addition
	if evaluateHelperAddMultConcat(target, constants, current+constants[index], index+1) {
		return true
	}

	// Try multiplication
	if evaluateHelperAddMultConcat(target, constants, current*constants[index], index+1) {
		return true
	}

	// Try concatenation
	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, constants[index]))
	if evaluateHelperAddMultConcat(target, constants, concatenated, index+1) {
		return true
	}

	return false
}

func Part2(filename string) {
	eqMap, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	ansSum := 0
	for ans, consts := range eqMap {
		if evaluateEquation(ans, consts, evaluateHelperAddMultConcat) {
			ansSum += ans
		}
	}
	fmt.Printf("Sum of constants that can make target: %d\n", ansSum)
}
