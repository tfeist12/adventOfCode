package day07

import (
	"fmt"
)

// Helper function to recursively evaluate all combinations
func evaluateHelperAddMult(target int, constants []int, current int, index int) bool {
	if index == len(constants) {
		return current == target
	}

	// Try addition
	if evaluateHelperAddMult(target, constants, current+constants[index], index+1) {
		return true
	}

	// Try multiplication
	if evaluateHelperAddMult(target, constants, current*constants[index], index+1) {
		return true
	}

	return false
}

func Part1(filename string) {
	eqMap, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	ansSum := 0
	for ans, consts := range eqMap {
		if evaluateEquation(ans, consts, evaluateHelperAddMult) {
			ansSum += ans
		}
	}
	fmt.Printf("Sum of constants that can make target: %d\n", ansSum)
}
