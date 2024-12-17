package day07

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
func evaluateEquation(target int, constants []int, helper func(int, []int, int, int) bool) bool {
	if len(constants) == 0 {
		return false
	}

	return helper(target, constants, constants[0], 1)
}
