package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filename = "input.txt"

// Read file and return two lists of integers
func readFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting string to int")
			return nil, nil, err
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return list1, list2, nil
}

func calculateDifference(list1 []int, list2 []int) int {
	// Sort lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Calculate sum of differences between lists
	sum := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] >= list2[i] {
			diff := list1[i] - list2[i]
			sum += diff
		} else {
			diff := list2[i] - list1[i]
			sum += diff
		}
	}

	return sum
}

func calculateSimilarity(list1 []int, list2 []int) int {
	similarity := 0
	for _, num1 := range list1 {
		occurences := 0
		for _, num2 := range list2 {
			if num1 == num2 {
				occurences++
			}
		}
		similarity += num1 * occurences
	}
	return similarity
}

func main() {
	list1, list2, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Ensure lists are the same length
	if len(list1) != len(list2) {
		fmt.Println("Lists are not the same length")
		return
	}

	diff := calculateDifference(list1, list2)
	fmt.Printf("Total Difference: %d\n", diff)

	sim := calculateSimilarity(list1, list2)
	fmt.Printf("Similarity: %d\n", sim)
}
