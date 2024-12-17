package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	regexPattMul = "mul\\(([0-9]+),([0-9]+)\\)"
)

func sumMultMatches(data string) (int, error) {
	// Regex with capture groups
	re := regexp.MustCompile(regexPattMul)
	matches := re.FindAllStringSubmatch(data, -1)

	sum := 0
	for _, match := range matches {
		num1, err1 := strconv.Atoi(match[1])
		if err1 != nil {
			fmt.Printf("Error converting '%s' to an int\n", match[1])
			return 0, err1
		}
		num2, err2 := strconv.Atoi(match[2])
		if err2 != nil {
			fmt.Printf("Error converting '%s' to an int\n", match[2])
			return 0, err2
		}
		sum += num1 * num2
	}
	return sum, nil
}

func Part1(filename string) {
	data, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	sum, err := sumMultMatches(data)
	if err != nil {
		fmt.Printf("Error summing matches: %v\n", err)
		return
	}
	fmt.Printf("Sum of mul operations: %d\n", sum)
}
