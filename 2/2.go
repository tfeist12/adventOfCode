package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	filename = "input.txt"
)

// Read file and return a list of a list of integers
func readFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strNums := strings.Split(line, " ")
		var nums []int
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Println("Error converting string to int")
				return nil, err
			}
			nums = append(nums, num)
		}
		reports = append(reports, nums)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

// Check if a single report is safe
func isSafe(report []int) bool {
	initialIncrease := report[1] > report[0]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		// Ensure levels differ by at least 1 and at most 3
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
		// Ensure levels are either only increasing or only decreasing
		notIncreasing := initialIncrease && report[i] < report[i-1]
		notDecreasing := !initialIncrease && report[i] > report[i-1]
		if notIncreasing || notDecreasing {
			return false
		}
	}
	return true
}

// Count number of safe reports
func countSafeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}

// Check if removing a single element from the report makes it safe
func removalMakesSafe(report []int) bool {
	for i := 0; i < len(report); i++ {
		modifiedSlice := slices.Delete(slices.Clone(report), i, i+1)
		if isSafe(modifiedSlice) {
			return true
		}
	}
	return false
}

// Count number of safe reports with problem dampener
func countSafeReportsProbDamp(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if removalMakesSafe(report) {
			count++
		}
	}
	return count
}

func main() {
	reports, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	safeReports := countSafeReports(reports)
	fmt.Printf("Safe reports without problem detector: %d\n", safeReports)

	safeReportsPD := countSafeReportsProbDamp(reports)
	fmt.Printf("Safe reports with problem dampener: %d\n", safeReportsPD)
}
