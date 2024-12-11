package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	exampleFilename = "example.txt"
	filename        = "input.txt"
)

// Reads a file and returns a map of equations
// Equation answers are the keys, constant lists are the values
func readFile(filename string) (map[int][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Map of equation answers to constants
	eqMap := make(map[int][]int)

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Extract equation answer
		ansSplit := strings.Split(line, ":")
		eqAns, err := strconv.Atoi(ansSplit[0])
		if err != nil {
			return nil, err
		}

		// Extract equation constants
		var eqConsts []int
		for _, s := range strings.Split(ansSplit[1][1:], " ") {
			eqConst, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			eqConsts = append(eqConsts, eqConst)
		}

		// Build return map
		eqMap[eqAns] = eqConsts
	}

	return eqMap, nil
}

// Evaluates if the equation can be made true with the given constants
func evaluateEquation(target int, constants []int) bool {
	if len(constants) == 0 {
		return false
	}

	return evaluateHelper(target, constants, constants[0], 1)
}

// Helper function to recursively evaluate all combinations
func evaluateHelper(target int, constants []int, current int, index int) bool {
	if index == len(constants) {
		return current == target
	}

	// Try addition
	if evaluateHelper(target, constants, current+constants[index], index+1) {
		return true
	}

	// Try multiplication
	if evaluateHelper(target, constants, current*constants[index], index+1) {
		return true
	}

	// Try concatenation
	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, constants[index]))
	if evaluateHelper(target, constants, concatenated, index+1) {
		return true
	}

	return false
}

func main() {
	eqMap, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	ansSum := 0
	for ans, consts := range eqMap {
		if evaluateEquation(ans, consts) {
			ansSum += ans
		}
	}
	fmt.Printf("Sum of constants that can make target: %d\n", ansSum)
}
