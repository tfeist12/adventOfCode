package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	regexPattDo   = "do\\(\\)"
	regexPattDont = "don't\\(\\)"
)

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

func Part2(filename string) {
	data, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	sum, err := sumDoMultMatches(data)
	if err != nil {
		fmt.Printf("Error summing matches: %v\n", err)
		return
	}
	fmt.Printf("Sum of mul operations with conditional checking: %d\n", sum)
}
