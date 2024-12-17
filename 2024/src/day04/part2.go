package day04

import (
	"fmt"
)

var patterns = [][][]string{
	{
		{"M", ".", "S"},
		{".", "A", "."},
		{"M", ".", "S"},
	},
	{
		{"S", ".", "M"},
		{".", "A", "."},
		{"S", ".", "M"},
	},
	{
		{"M", ".", "M"},
		{".", "A", "."},
		{"S", ".", "S"},
	},
	{
		{"S", ".", "S"},
		{".", "A", "."},
		{"M", ".", "M"},
	},
}

func checkXmasPattern(grid [][]string, y, x int, pattern [][]string) bool {
	rows := len(grid)
	cols := len(grid[0])
	patternRows := len(pattern)
	patternCols := len(pattern[0])

	// Iterate through the pattern
	for i := 0; i < patternRows; i++ {
		for j := 0; j < patternCols; j++ {
			// Calculate new indices in the grid
			ny := y + i
			nx := x + j
			// Return false if new position is out of bounds
			if ny < 0 || ny >= rows || nx < 0 || nx >= cols {
				return false
			}
			if pattern[i][j] != "." && pattern[i][j] != grid[ny][nx] {
				return false
			}
		}
	}
	return true
}

func countXmasPattern(grid [][]string) int {
	count := 0
	// Iterate through rows
	for i := 0; i < len(grid); i++ {
		// Iterate through columns
		for j := 0; j < len(grid[i]); j++ {
			// Iterate through valid patterns
			for _, pattern := range patterns {
				if checkXmasPattern(grid, i, j, pattern) {
					count += 1
				}
			}
		}
	}
	return count
}

func Part2(filename string) {
	puzzleGrid, err := readFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	patterns := countXmasPattern(puzzleGrid)
	fmt.Printf("Number of 'X-MAS' patterns in the word search: %d\n", patterns)
}
