package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

const (
	filename      = "input.txt"
	regexPattMul  = "mul\\(([0-9]+),([0-9]+)\\)"
	regexPattDo   = "do\\(\\)"
	regexPattDont = "don't\\(\\)"
)

// Read full file into a string
func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

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

func sumDoMultMatches(data string) (int, error) {
	// Define the regex patterns
	reDo := regexp.MustCompile(regexPattDo)
	reDont := regexp.MustCompile(regexPattDont)
	reAll := regexp.MustCompile(
		fmt.Sprintf("%s|%s|%s", regexPattDo, regexPattDont, regexPattMul),
	)

	// Find all matches for each pattern
	matches := reAll.FindAllString(data, -1)

	// Initialize loop variables
	shouldSum := true
	sum := 0

	for _, match := range matches {
		if reDo.MatchString(match) {
			shouldSum = true
		} else if reDont.MatchString(match) {
			shouldSum = false
		} else {
			// Extract the numbers from the match
			re := regexp.MustCompile(regexPattMul)
			nums := re.FindStringSubmatch(match)
			num1, err1 := strconv.Atoi(nums[1])
			if err1 != nil {
				fmt.Printf("Error converting '%s' to an int\n", nums[1])
				return 0, err1
			}
			num2, err2 := strconv.Atoi(nums[2])
			if err2 != nil {
				fmt.Printf("Error converting '%s' to an int\n", nums[2])
				return 0, err2
			}
			// Sum the numbers if flag is set
			if shouldSum {
				sum += num1 * num2
			}
		}
	}

	return sum, nil
}

func main() {
	data, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Part 1
	sum, err := sumMultMatches(data)
	if err != nil {
		fmt.Printf("Error summing matches: %v\n", err)
		return
	}
	fmt.Printf("Sum of mul operations: %d\n", sum)

	// Part 2
	sum, err = sumDoMultMatches(data)
	if err != nil {
		fmt.Printf("Error summing matches: %v\n", err)
		return
	}
	fmt.Printf("Sum of mul operations with conditional checking: %d\n", sum)
}
