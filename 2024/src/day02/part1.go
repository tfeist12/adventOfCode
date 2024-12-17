package day02

import (
	"fmt"
)

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

func Part1(filename string) {
	reports, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	safeReports := countSafeReports(reports)
	fmt.Printf("Safe reports without problem detector: %d\n", safeReports)
}
