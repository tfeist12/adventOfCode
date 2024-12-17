package day02

import (
	"fmt"
	"slices"
)

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

func Part2(filename string) {
	reports, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	safeReportsPD := countSafeReportsProbDamp(reports)
	fmt.Printf("Safe reports with problem dampener: %d\n", safeReportsPD)
}
